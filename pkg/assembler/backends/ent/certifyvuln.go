// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifyvuln"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/securityadvisory"
)

// CertifyVuln is the model entity for the CertifyVuln schema.
type CertifyVuln struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Advisory is one of OSV, GHSA, or CVE, or nil if not vulnerable
	VulnerabilityID *int `json:"vulnerability_id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID int `json:"package_id,omitempty"`
	// TimeScanned holds the value of the "time_scanned" field.
	TimeScanned time.Time `json:"time_scanned,omitempty"`
	// DbURI holds the value of the "db_uri" field.
	DbURI string `json:"db_uri,omitempty"`
	// DbVersion holds the value of the "db_version" field.
	DbVersion string `json:"db_version,omitempty"`
	// ScannerURI holds the value of the "scanner_uri" field.
	ScannerURI string `json:"scanner_uri,omitempty"`
	// ScannerVersion holds the value of the "scanner_version" field.
	ScannerVersion string `json:"scanner_version,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
	// Collector holds the value of the "collector" field.
	Collector string `json:"collector,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CertifyVulnQuery when eager-loading is set.
	Edges        CertifyVulnEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CertifyVulnEdges holds the relations/edges for other nodes in the graph.
type CertifyVulnEdges struct {
	// Vulnerability is one of OSV, GHSA, or CVE
	Vulnerability *SecurityAdvisory `json:"vulnerability,omitempty"`
	// Package holds the value of the package edge.
	Package *PackageVersion `json:"package,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// VulnerabilityOrErr returns the Vulnerability value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertifyVulnEdges) VulnerabilityOrErr() (*SecurityAdvisory, error) {
	if e.loadedTypes[0] {
		if e.Vulnerability == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: securityadvisory.Label}
		}
		return e.Vulnerability, nil
	}
	return nil, &NotLoadedError{edge: "vulnerability"}
}

// PackageOrErr returns the Package value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertifyVulnEdges) PackageOrErr() (*PackageVersion, error) {
	if e.loadedTypes[1] {
		if e.Package == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.Package, nil
	}
	return nil, &NotLoadedError{edge: "package"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CertifyVuln) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certifyvuln.FieldID, certifyvuln.FieldVulnerabilityID, certifyvuln.FieldPackageID:
			values[i] = new(sql.NullInt64)
		case certifyvuln.FieldDbURI, certifyvuln.FieldDbVersion, certifyvuln.FieldScannerURI, certifyvuln.FieldScannerVersion, certifyvuln.FieldOrigin, certifyvuln.FieldCollector:
			values[i] = new(sql.NullString)
		case certifyvuln.FieldTimeScanned:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertifyVuln fields.
func (cv *CertifyVuln) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certifyvuln.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cv.ID = int(value.Int64)
		case certifyvuln.FieldVulnerabilityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vulnerability_id", values[i])
			} else if value.Valid {
				cv.VulnerabilityID = new(int)
				*cv.VulnerabilityID = int(value.Int64)
			}
		case certifyvuln.FieldPackageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				cv.PackageID = int(value.Int64)
			}
		case certifyvuln.FieldTimeScanned:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time_scanned", values[i])
			} else if value.Valid {
				cv.TimeScanned = value.Time
			}
		case certifyvuln.FieldDbURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field db_uri", values[i])
			} else if value.Valid {
				cv.DbURI = value.String
			}
		case certifyvuln.FieldDbVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field db_version", values[i])
			} else if value.Valid {
				cv.DbVersion = value.String
			}
		case certifyvuln.FieldScannerURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scanner_uri", values[i])
			} else if value.Valid {
				cv.ScannerURI = value.String
			}
		case certifyvuln.FieldScannerVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scanner_version", values[i])
			} else if value.Valid {
				cv.ScannerVersion = value.String
			}
		case certifyvuln.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				cv.Origin = value.String
			}
		case certifyvuln.FieldCollector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collector", values[i])
			} else if value.Valid {
				cv.Collector = value.String
			}
		default:
			cv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CertifyVuln.
// This includes values selected through modifiers, order, etc.
func (cv *CertifyVuln) Value(name string) (ent.Value, error) {
	return cv.selectValues.Get(name)
}

// QueryVulnerability queries the "vulnerability" edge of the CertifyVuln entity.
func (cv *CertifyVuln) QueryVulnerability() *SecurityAdvisoryQuery {
	return NewCertifyVulnClient(cv.config).QueryVulnerability(cv)
}

// QueryPackage queries the "package" edge of the CertifyVuln entity.
func (cv *CertifyVuln) QueryPackage() *PackageVersionQuery {
	return NewCertifyVulnClient(cv.config).QueryPackage(cv)
}

// Update returns a builder for updating this CertifyVuln.
// Note that you need to call CertifyVuln.Unwrap() before calling this method if this CertifyVuln
// was returned from a transaction, and the transaction was committed or rolled back.
func (cv *CertifyVuln) Update() *CertifyVulnUpdateOne {
	return NewCertifyVulnClient(cv.config).UpdateOne(cv)
}

// Unwrap unwraps the CertifyVuln entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cv *CertifyVuln) Unwrap() *CertifyVuln {
	_tx, ok := cv.config.driver.(*txDriver)
	if !ok {
		panic("ent: CertifyVuln is not a transactional entity")
	}
	cv.config.driver = _tx.drv
	return cv
}

// String implements the fmt.Stringer.
func (cv *CertifyVuln) String() string {
	var builder strings.Builder
	builder.WriteString("CertifyVuln(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cv.ID))
	if v := cv.VulnerabilityID; v != nil {
		builder.WriteString("vulnerability_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("package_id=")
	builder.WriteString(fmt.Sprintf("%v", cv.PackageID))
	builder.WriteString(", ")
	builder.WriteString("time_scanned=")
	builder.WriteString(cv.TimeScanned.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("db_uri=")
	builder.WriteString(cv.DbURI)
	builder.WriteString(", ")
	builder.WriteString("db_version=")
	builder.WriteString(cv.DbVersion)
	builder.WriteString(", ")
	builder.WriteString("scanner_uri=")
	builder.WriteString(cv.ScannerURI)
	builder.WriteString(", ")
	builder.WriteString("scanner_version=")
	builder.WriteString(cv.ScannerVersion)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(cv.Origin)
	builder.WriteString(", ")
	builder.WriteString("collector=")
	builder.WriteString(cv.Collector)
	builder.WriteByte(')')
	return builder.String()
}

// CertifyVulns is a parsable slice of CertifyVuln.
type CertifyVulns []*CertifyVuln
