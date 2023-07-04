// Code generated by ent, DO NOT EDIT.

package builder

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the builder type in the database.
	Label = "builder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// EdgeSlsaAttestation holds the string denoting the slsa_attestation edge name in mutations.
	EdgeSlsaAttestation = "slsa_attestation"
	// Table holds the table name of the builder in the database.
	Table = "builders"
	// SlsaAttestationTable is the table that holds the slsa_attestation relation/edge.
	SlsaAttestationTable = "builders"
	// SlsaAttestationInverseTable is the table name for the SLSAAttestation entity.
	// It exists in this package in order to avoid circular dependency with the "slsaattestation" package.
	SlsaAttestationInverseTable = "slsa_attestations"
	// SlsaAttestationColumn is the table column denoting the slsa_attestation relation/edge.
	SlsaAttestationColumn = "slsa_attestation_built_by"
)

// Columns holds all SQL columns for builder fields.
var Columns = []string{
	FieldID,
	FieldURI,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "builders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"slsa_attestation_built_by",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Builder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// BySlsaAttestationField orders the results by slsa_attestation field.
func BySlsaAttestationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSlsaAttestationStep(), sql.OrderByField(field, opts...))
	}
}
func newSlsaAttestationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SlsaAttestationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SlsaAttestationTable, SlsaAttestationColumn),
	)
}