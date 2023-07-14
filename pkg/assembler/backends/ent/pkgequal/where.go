// Code generated by ent, DO NOT EDIT.

package pkgequal

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLTE(FieldID, id))
}

// PackageVersionID applies equality check predicate on the "package_version_id" field. It's identical to PackageVersionIDEQ.
func PackageVersionID(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldPackageVersionID, v))
}

// SimilarID applies equality check predicate on the "similar_id" field. It's identical to SimilarIDEQ.
func SimilarID(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldSimilarID, v))
}

// Origin applies equality check predicate on the "origin" field. It's identical to OriginEQ.
func Origin(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldOrigin, v))
}

// Collector applies equality check predicate on the "collector" field. It's identical to CollectorEQ.
func Collector(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldCollector, v))
}

// Justification applies equality check predicate on the "justification" field. It's identical to JustificationEQ.
func Justification(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldJustification, v))
}

// PackageVersionIDEQ applies the EQ predicate on the "package_version_id" field.
func PackageVersionIDEQ(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldPackageVersionID, v))
}

// PackageVersionIDNEQ applies the NEQ predicate on the "package_version_id" field.
func PackageVersionIDNEQ(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldPackageVersionID, v))
}

// PackageVersionIDIn applies the In predicate on the "package_version_id" field.
func PackageVersionIDIn(vs ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldPackageVersionID, vs...))
}

// PackageVersionIDNotIn applies the NotIn predicate on the "package_version_id" field.
func PackageVersionIDNotIn(vs ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldPackageVersionID, vs...))
}

// SimilarIDEQ applies the EQ predicate on the "similar_id" field.
func SimilarIDEQ(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldSimilarID, v))
}

// SimilarIDNEQ applies the NEQ predicate on the "similar_id" field.
func SimilarIDNEQ(v int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldSimilarID, v))
}

// SimilarIDIn applies the In predicate on the "similar_id" field.
func SimilarIDIn(vs ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldSimilarID, vs...))
}

// SimilarIDNotIn applies the NotIn predicate on the "similar_id" field.
func SimilarIDNotIn(vs ...int) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldSimilarID, vs...))
}

// OriginEQ applies the EQ predicate on the "origin" field.
func OriginEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldOrigin, v))
}

// OriginNEQ applies the NEQ predicate on the "origin" field.
func OriginNEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldOrigin, v))
}

// OriginIn applies the In predicate on the "origin" field.
func OriginIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldOrigin, vs...))
}

// OriginNotIn applies the NotIn predicate on the "origin" field.
func OriginNotIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldOrigin, vs...))
}

// OriginGT applies the GT predicate on the "origin" field.
func OriginGT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGT(FieldOrigin, v))
}

// OriginGTE applies the GTE predicate on the "origin" field.
func OriginGTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGTE(FieldOrigin, v))
}

// OriginLT applies the LT predicate on the "origin" field.
func OriginLT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLT(FieldOrigin, v))
}

// OriginLTE applies the LTE predicate on the "origin" field.
func OriginLTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLTE(FieldOrigin, v))
}

// OriginContains applies the Contains predicate on the "origin" field.
func OriginContains(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContains(FieldOrigin, v))
}

// OriginHasPrefix applies the HasPrefix predicate on the "origin" field.
func OriginHasPrefix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasPrefix(FieldOrigin, v))
}

// OriginHasSuffix applies the HasSuffix predicate on the "origin" field.
func OriginHasSuffix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasSuffix(FieldOrigin, v))
}

// OriginEqualFold applies the EqualFold predicate on the "origin" field.
func OriginEqualFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEqualFold(FieldOrigin, v))
}

// OriginContainsFold applies the ContainsFold predicate on the "origin" field.
func OriginContainsFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContainsFold(FieldOrigin, v))
}

// CollectorEQ applies the EQ predicate on the "collector" field.
func CollectorEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldCollector, v))
}

// CollectorNEQ applies the NEQ predicate on the "collector" field.
func CollectorNEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldCollector, v))
}

// CollectorIn applies the In predicate on the "collector" field.
func CollectorIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldCollector, vs...))
}

// CollectorNotIn applies the NotIn predicate on the "collector" field.
func CollectorNotIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldCollector, vs...))
}

// CollectorGT applies the GT predicate on the "collector" field.
func CollectorGT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGT(FieldCollector, v))
}

// CollectorGTE applies the GTE predicate on the "collector" field.
func CollectorGTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGTE(FieldCollector, v))
}

// CollectorLT applies the LT predicate on the "collector" field.
func CollectorLT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLT(FieldCollector, v))
}

// CollectorLTE applies the LTE predicate on the "collector" field.
func CollectorLTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLTE(FieldCollector, v))
}

// CollectorContains applies the Contains predicate on the "collector" field.
func CollectorContains(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContains(FieldCollector, v))
}

// CollectorHasPrefix applies the HasPrefix predicate on the "collector" field.
func CollectorHasPrefix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasPrefix(FieldCollector, v))
}

// CollectorHasSuffix applies the HasSuffix predicate on the "collector" field.
func CollectorHasSuffix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasSuffix(FieldCollector, v))
}

// CollectorEqualFold applies the EqualFold predicate on the "collector" field.
func CollectorEqualFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEqualFold(FieldCollector, v))
}

// CollectorContainsFold applies the ContainsFold predicate on the "collector" field.
func CollectorContainsFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContainsFold(FieldCollector, v))
}

// JustificationEQ applies the EQ predicate on the "justification" field.
func JustificationEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEQ(FieldJustification, v))
}

// JustificationNEQ applies the NEQ predicate on the "justification" field.
func JustificationNEQ(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNEQ(FieldJustification, v))
}

// JustificationIn applies the In predicate on the "justification" field.
func JustificationIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldIn(FieldJustification, vs...))
}

// JustificationNotIn applies the NotIn predicate on the "justification" field.
func JustificationNotIn(vs ...string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldNotIn(FieldJustification, vs...))
}

// JustificationGT applies the GT predicate on the "justification" field.
func JustificationGT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGT(FieldJustification, v))
}

// JustificationGTE applies the GTE predicate on the "justification" field.
func JustificationGTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldGTE(FieldJustification, v))
}

// JustificationLT applies the LT predicate on the "justification" field.
func JustificationLT(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLT(FieldJustification, v))
}

// JustificationLTE applies the LTE predicate on the "justification" field.
func JustificationLTE(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldLTE(FieldJustification, v))
}

// JustificationContains applies the Contains predicate on the "justification" field.
func JustificationContains(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContains(FieldJustification, v))
}

// JustificationHasPrefix applies the HasPrefix predicate on the "justification" field.
func JustificationHasPrefix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasPrefix(FieldJustification, v))
}

// JustificationHasSuffix applies the HasSuffix predicate on the "justification" field.
func JustificationHasSuffix(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldHasSuffix(FieldJustification, v))
}

// JustificationEqualFold applies the EqualFold predicate on the "justification" field.
func JustificationEqualFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldEqualFold(FieldJustification, v))
}

// JustificationContainsFold applies the ContainsFold predicate on the "justification" field.
func JustificationContainsFold(v string) predicate.PkgEqual {
	return predicate.PkgEqual(sql.FieldContainsFold(FieldJustification, v))
}

// HasPackageA applies the HasEdge predicate on the "package_a" edge.
func HasPackageA() predicate.PkgEqual {
	return predicate.PkgEqual(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PackageATable, PackageAColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPackageAWith applies the HasEdge predicate on the "package_a" edge with a given conditions (other predicates).
func HasPackageAWith(preds ...predicate.PackageVersion) predicate.PkgEqual {
	return predicate.PkgEqual(func(s *sql.Selector) {
		step := newPackageAStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPackageB applies the HasEdge predicate on the "package_b" edge.
func HasPackageB() predicate.PkgEqual {
	return predicate.PkgEqual(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PackageBTable, PackageBColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPackageBWith applies the HasEdge predicate on the "package_b" edge with a given conditions (other predicates).
func HasPackageBWith(preds ...predicate.PackageVersion) predicate.PkgEqual {
	return predicate.PkgEqual(func(s *sql.Selector) {
		step := newPackageBStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PkgEqual) predicate.PkgEqual {
	return predicate.PkgEqual(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PkgEqual) predicate.PkgEqual {
	return predicate.PkgEqual(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PkgEqual) predicate.PkgEqual {
	return predicate.PkgEqual(sql.NotPredicates(p))
}