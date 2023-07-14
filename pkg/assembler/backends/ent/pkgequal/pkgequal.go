// Code generated by ent, DO NOT EDIT.

package pkgequal

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pkgequal type in the database.
	Label = "pkg_equal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPackageVersionID holds the string denoting the package_version_id field in the database.
	FieldPackageVersionID = "package_version_id"
	// FieldSimilarID holds the string denoting the similar_id field in the database.
	FieldSimilarID = "similar_id"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// EdgePackageA holds the string denoting the package_a edge name in mutations.
	EdgePackageA = "package_a"
	// EdgePackageB holds the string denoting the package_b edge name in mutations.
	EdgePackageB = "package_b"
	// Table holds the table name of the pkgequal in the database.
	Table = "pkg_equals"
	// PackageATable is the table that holds the package_a relation/edge.
	PackageATable = "pkg_equals"
	// PackageAInverseTable is the table name for the PackageVersion entity.
	// It exists in this package in order to avoid circular dependency with the "packageversion" package.
	PackageAInverseTable = "package_versions"
	// PackageAColumn is the table column denoting the package_a relation/edge.
	PackageAColumn = "package_version_id"
	// PackageBTable is the table that holds the package_b relation/edge.
	PackageBTable = "pkg_equals"
	// PackageBInverseTable is the table name for the PackageVersion entity.
	// It exists in this package in order to avoid circular dependency with the "packageversion" package.
	PackageBInverseTable = "package_versions"
	// PackageBColumn is the table column denoting the package_b relation/edge.
	PackageBColumn = "similar_id"
)

// Columns holds all SQL columns for pkgequal fields.
var Columns = []string{
	FieldID,
	FieldPackageVersionID,
	FieldSimilarID,
	FieldOrigin,
	FieldCollector,
	FieldJustification,
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

// OrderOption defines the ordering options for the PkgEqual queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPackageVersionID orders the results by the package_version_id field.
func ByPackageVersionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPackageVersionID, opts...).ToFunc()
}

// BySimilarID orders the results by the similar_id field.
func BySimilarID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSimilarID, opts...).ToFunc()
}

// ByOrigin orders the results by the origin field.
func ByOrigin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrigin, opts...).ToFunc()
}

// ByCollector orders the results by the collector field.
func ByCollector(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollector, opts...).ToFunc()
}

// ByJustification orders the results by the justification field.
func ByJustification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJustification, opts...).ToFunc()
}

// ByPackageAField orders the results by package_a field.
func ByPackageAField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackageAStep(), sql.OrderByField(field, opts...))
	}
}

// ByPackageBField orders the results by package_b field.
func ByPackageBField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackageBStep(), sql.OrderByField(field, opts...))
	}
}
func newPackageAStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackageAInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PackageATable, PackageAColumn),
	)
}
func newPackageBStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackageBInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PackageBTable, PackageBColumn),
	)
}