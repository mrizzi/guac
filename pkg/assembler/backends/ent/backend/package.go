package backend

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"sort"

	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagenamespace"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagetype"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/pkg/errors"
)

func (b *EntBackend) Packages(ctx context.Context, pkgSpec *model.PkgSpec) ([]*model.Package, error) {
	query := b.client.PackageType.Query().Order(ent.Asc(packagetype.FieldType))

	paths := getPreloads(ctx)

	if pkgSpec == nil {
		pkgSpec = &model.PkgSpec{}
	}

	query.Where(optionalPredicate(pkgSpec.Type, packagetype.TypeEQ))

	if PathContains(paths, "namespaces") {
		query.WithNamespaces(func(q *ent.PackageNamespaceQuery) {
			q.Order(ent.Asc(packagenamespace.FieldNamespace))
			q.Where(optionalPredicate(pkgSpec.Namespace, packagenamespace.NamespaceEQ))

			if PathContains(paths, "namespaces.names") {
				q.WithNames(func(q *ent.PackageNameQuery) {
					q.Order(ent.Asc(packagename.FieldName))
					q.Where(optionalPredicate(pkgSpec.Name, packagename.NameEQ))

					if PathContains(paths, "namespaces.names.versions") {
						q.WithVersions(func(q *ent.PackageVersionQuery) {
							q.Order(ent.Asc(packageversion.FieldVersion))
							q.Where(
								optionalPredicate(pkgSpec.Version, packageversion.VersionEQ),
								optionalPredicate(pkgSpec.Subpath, packageversion.SubpathEQ),
								packageversion.QualifiersMatchSpec(pkgSpec.Qualifiers),
							)
						})
					}
				})
			}
		})
	}

	// FIXME: (ivanvanderbyl) This could be much more compact and use a single query as above.
	if pkgSpec != nil {
		query.Where(optionalPredicate(pkgSpec.ID, IDEQ))
	} else {
		query.Limit(100)
	}

	pkgs, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return collect(pkgs, toModelPackage), nil
}

func (b *EntBackend) IngestPackages(ctx context.Context, pkgs []*model.PkgInputSpec) ([]*model.Package, error) {
	// FIXME: (ivanvanderbyl) This will be suboptimal, but we can't batch insert relations with upserts.
	models := make([]*model.Package, len(pkgs))
	for i, pkg := range pkgs {
		p, err := b.IngestPackage(ctx, *pkg)
		if err != nil {
			return nil, err
		}
		models[i] = p
	}
	return models, nil
}

func (b *EntBackend) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	pkgVersion, err := WithinTX(ctx, b.client, func(ctx context.Context) (*ent.PackageVersion, error) {
		client := ent.FromContext(ctx)
		p, err := upsertPackage(ctx, client, pkg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to upsert package")
		}
		return p, nil
	})
	if err != nil {
		return nil, err
	}

	record, err := pkgTreeFromVersion(ctx, pkgVersion.Unwrap())
	if err != nil {
		return nil, err
	}

	return toModelPackage(record), nil
}

// upsertPackage is a helper function to create or update a package node and its associated edges.
// It is used in multiple places, so we extract it to a function.
func upsertPackage(ctx context.Context, client *ent.Client, pkg model.PkgInputSpec) (*ent.PackageVersion, error) {
	pkgID, err := client.PackageType.Create().SetType(pkg.Type).
		OnConflict(sql.ConflictColumns(packagetype.FieldType)).UpdateNewValues().ID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "upsert package node")
	}

	nsID, err := client.PackageNamespace.Create().SetPackageID(pkgID).SetNamespace(valueOrDefault(pkg.Namespace, "")).
		OnConflict(sql.ConflictColumns(packagenamespace.FieldNamespace, packagenamespace.FieldPackageID)).UpdateNewValues().ID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "upsert package namespace")
	}

	nameID, err := client.PackageName.Create().SetNamespaceID(nsID).SetName(pkg.Name).
		OnConflict(sql.ConflictColumns(packagename.FieldName, packagename.FieldNamespaceID)).UpdateNewValues().ID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "upsert package name")
	}

	pvID, err := client.PackageVersion.Create().
		SetNameID(nameID).
		SetVersion(valueOrDefault(pkg.Version, "")).
		SetSubpath(valueOrDefault(pkg.Subpath, "")).
		SetQualifiers(normalizeInputQualifiers(pkg.Qualifiers)).
		SetHash(versionHashFromInputSpec(pkg)).
		OnConflict(
			sql.ConflictColumns(
				packageversion.FieldHash,
				packageversion.FieldNameID,
			),
		).
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "upsert package version")
	}

	pv, err := client.PackageVersion.Get(ctx, pvID)
	if err != nil {
		return nil, errors.Wrap(err, "get package version")
	}

	return pv, nil
}

func versionHashFromInputSpec(pkg model.PkgInputSpec) string {
	return hashPackageVersion(
		valueOrDefault(pkg.Version, ""),
		valueOrDefault(pkg.Subpath, ""),
		normalizeInputQualifiers(pkg.Qualifiers))
}

func hashPackageVersion(version, subpath string, qualifiers []model.PackageQualifier) string {
	hash := sha1.New()
	hash.Write([]byte(version))
	hash.Write([]byte(subpath))
	qualifiersBuffer := bytes.NewBuffer(nil)

	sort.Slice(qualifiers, func(i, j int) bool { return qualifiers[i].Key < qualifiers[j].Key })

	for _, qualifier := range qualifiers {
		qualifiersBuffer.WriteString(qualifier.Key)
		qualifiersBuffer.WriteString(qualifier.Value)
	}

	hash.Write(qualifiersBuffer.Bytes())
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func normalizeInputQualifiers(inputs []*model.PackageQualifierInputSpec) []model.PackageQualifier {
	qualifiers := []model.PackageQualifier{}
	for _, q := range inputs {
		qualifiers = append(qualifiers, model.PackageQualifier{
			Key:   q.Key,
			Value: q.Value,
		})
	}

	return qualifiers
}

func pkgVersionPredicates(spec *model.PkgSpec) []predicate.PackageVersion {
	if spec == nil {
		return nil
	}
	rv := []predicate.PackageVersion{
		optionalPredicate(spec.ID, IDEQ),
		packageversion.VersionEQ(valueOrDefault(spec.Version, "")),
		// optionalPredicate(spec.Version, packageversion.Version),
		optionalPredicate(spec.Subpath, packageversion.Subpath),
		packageversion.QualifiersMatchSpec(spec.Qualifiers),
		packageversion.HasNameWith(
			optionalPredicate(spec.Name, packagename.Name),
			packagename.HasNamespaceWith(
				optionalPredicate(spec.Namespace, packagenamespace.Namespace),
				packagenamespace.HasPackageWith(
					optionalPredicate(spec.Type, packagetype.Type),
				),
			),
		),
	}

	return rv
}

func pkgNamePredicates(spec *model.PkgNameSpec) []predicate.PackageName {
	if spec == nil {
		return nil
	}
	return []predicate.PackageName{
		optionalPredicate(spec.ID, IDEQ),
		optionalPredicate(spec.Name, packagename.Name),
		packagename.HasNamespaceWith(
			optionalPredicate(spec.Namespace, packagenamespace.Namespace),
			packagenamespace.HasPackageWith(
				optionalPredicate(spec.Type, packagetype.Type),
			),
		),
	}
}