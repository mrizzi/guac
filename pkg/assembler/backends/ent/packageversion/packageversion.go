// Code generated by ent, DO NOT EDIT.

package packageversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the packageversion type in the database.
	Label = "package_version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNameID holds the string denoting the name_id field in the database.
	FieldNameID = "name_id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldSubpath holds the string denoting the subpath field in the database.
	FieldSubpath = "subpath"
	// FieldQualifiers holds the string denoting the qualifiers field in the database.
	FieldQualifiers = "qualifiers"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// EdgeName holds the string denoting the name edge name in mutations.
	EdgeName = "name"
	// EdgeOccurrences holds the string denoting the occurrences edge name in mutations.
	EdgeOccurrences = "occurrences"
	// EdgeSbom holds the string denoting the sbom edge name in mutations.
	EdgeSbom = "sbom"
	// Table holds the table name of the packageversion in the database.
	Table = "package_versions"
	// NameTable is the table that holds the name relation/edge.
	NameTable = "package_versions"
	// NameInverseTable is the table name for the PackageName entity.
	// It exists in this package in order to avoid circular dependency with the "packagename" package.
	NameInverseTable = "package_names"
	// NameColumn is the table column denoting the name relation/edge.
	NameColumn = "name_id"
	// OccurrencesTable is the table that holds the occurrences relation/edge.
	OccurrencesTable = "occurrences"
	// OccurrencesInverseTable is the table name for the Occurrence entity.
	// It exists in this package in order to avoid circular dependency with the "occurrence" package.
	OccurrencesInverseTable = "occurrences"
	// OccurrencesColumn is the table column denoting the occurrences relation/edge.
	OccurrencesColumn = "package_id"
	// SbomTable is the table that holds the sbom relation/edge.
	SbomTable = "sbo_ms"
	// SbomInverseTable is the table name for the SBOM entity.
	// It exists in this package in order to avoid circular dependency with the "sbom" package.
	SbomInverseTable = "sbo_ms"
	// SbomColumn is the table column denoting the sbom relation/edge.
	SbomColumn = "package_id"
)

// Columns holds all SQL columns for packageversion fields.
var Columns = []string{
	FieldID,
	FieldNameID,
	FieldVersion,
	FieldSubpath,
	FieldQualifiers,
	FieldHash,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion string
	// DefaultSubpath holds the default value on creation for the "subpath" field.
	DefaultSubpath string
)

// OrderOption defines the ordering options for the PackageVersion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNameID orders the results by the name_id field.
func ByNameID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameID, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// BySubpath orders the results by the subpath field.
func BySubpath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubpath, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByNameField orders the results by name field.
func ByNameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNameStep(), sql.OrderByField(field, opts...))
	}
}

// ByOccurrencesCount orders the results by occurrences count.
func ByOccurrencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOccurrencesStep(), opts...)
	}
}

// ByOccurrences orders the results by occurrences terms.
func ByOccurrences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOccurrencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySbomCount orders the results by sbom count.
func BySbomCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSbomStep(), opts...)
	}
}

// BySbom orders the results by sbom terms.
func BySbom(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSbomStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NameInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NameTable, NameColumn),
	)
}
func newOccurrencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OccurrencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OccurrencesTable, OccurrencesColumn),
	)
}
func newSbomStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SbomInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SbomTable, SbomColumn),
	)
}