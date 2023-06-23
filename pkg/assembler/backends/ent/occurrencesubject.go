// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrencesubject"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// OccurrenceSubject is the model entity for the OccurrenceSubject schema.
type OccurrenceSubject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID *int `json:"source_id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID *int `json:"package_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OccurrenceSubjectQuery when eager-loading is set.
	Edges        OccurrenceSubjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OccurrenceSubjectEdges holds the relations/edges for other nodes in the graph.
type OccurrenceSubjectEdges struct {
	// Occurrence holds the value of the occurrence edge.
	Occurrence *Occurrence `json:"occurrence,omitempty"`
	// Package holds the value of the package edge.
	Package *PackageVersion `json:"package,omitempty"`
	// Source holds the value of the source edge.
	Source *SourceName `json:"source,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OccurrenceOrErr returns the Occurrence value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OccurrenceSubjectEdges) OccurrenceOrErr() (*Occurrence, error) {
	if e.loadedTypes[0] {
		if e.Occurrence == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: occurrence.Label}
		}
		return e.Occurrence, nil
	}
	return nil, &NotLoadedError{edge: "occurrence"}
}

// PackageOrErr returns the Package value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OccurrenceSubjectEdges) PackageOrErr() (*PackageVersion, error) {
	if e.loadedTypes[1] {
		if e.Package == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.Package, nil
	}
	return nil, &NotLoadedError{edge: "package"}
}

// SourceOrErr returns the Source value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OccurrenceSubjectEdges) SourceOrErr() (*SourceName, error) {
	if e.loadedTypes[2] {
		if e.Source == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sourcename.Label}
		}
		return e.Source, nil
	}
	return nil, &NotLoadedError{edge: "source"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OccurrenceSubject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case occurrencesubject.FieldID, occurrencesubject.FieldSourceID, occurrencesubject.FieldPackageID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OccurrenceSubject fields.
func (os *OccurrenceSubject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case occurrencesubject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int(value.Int64)
		case occurrencesubject.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				os.SourceID = new(int)
				*os.SourceID = int(value.Int64)
			}
		case occurrencesubject.FieldPackageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				os.PackageID = new(int)
				*os.PackageID = int(value.Int64)
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OccurrenceSubject.
// This includes values selected through modifiers, order, etc.
func (os *OccurrenceSubject) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryOccurrence queries the "occurrence" edge of the OccurrenceSubject entity.
func (os *OccurrenceSubject) QueryOccurrence() *OccurrenceQuery {
	return NewOccurrenceSubjectClient(os.config).QueryOccurrence(os)
}

// QueryPackage queries the "package" edge of the OccurrenceSubject entity.
func (os *OccurrenceSubject) QueryPackage() *PackageVersionQuery {
	return NewOccurrenceSubjectClient(os.config).QueryPackage(os)
}

// QuerySource queries the "source" edge of the OccurrenceSubject entity.
func (os *OccurrenceSubject) QuerySource() *SourceNameQuery {
	return NewOccurrenceSubjectClient(os.config).QuerySource(os)
}

// Update returns a builder for updating this OccurrenceSubject.
// Note that you need to call OccurrenceSubject.Unwrap() before calling this method if this OccurrenceSubject
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OccurrenceSubject) Update() *OccurrenceSubjectUpdateOne {
	return NewOccurrenceSubjectClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OccurrenceSubject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OccurrenceSubject) Unwrap() *OccurrenceSubject {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OccurrenceSubject is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OccurrenceSubject) String() string {
	var builder strings.Builder
	builder.WriteString("OccurrenceSubject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	if v := os.SourceID; v != nil {
		builder.WriteString("source_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := os.PackageID; v != nil {
		builder.WriteString("package_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// OccurrenceSubjects is a parsable slice of OccurrenceSubject.
type OccurrenceSubjects []*OccurrenceSubject
