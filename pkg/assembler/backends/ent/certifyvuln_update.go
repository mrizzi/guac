// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifyvuln"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/vulnerability"
)

// CertifyVulnUpdate is the builder for updating CertifyVuln entities.
type CertifyVulnUpdate struct {
	config
	hooks    []Hook
	mutation *CertifyVulnMutation
}

// Where appends a list predicates to the CertifyVulnUpdate builder.
func (cvu *CertifyVulnUpdate) Where(ps ...predicate.CertifyVuln) *CertifyVulnUpdate {
	cvu.mutation.Where(ps...)
	return cvu
}

// SetVulnerabilityID sets the "vulnerability_id" field.
func (cvu *CertifyVulnUpdate) SetVulnerabilityID(i int) *CertifyVulnUpdate {
	cvu.mutation.SetVulnerabilityID(i)
	return cvu
}

// SetNillableVulnerabilityID sets the "vulnerability_id" field if the given value is not nil.
func (cvu *CertifyVulnUpdate) SetNillableVulnerabilityID(i *int) *CertifyVulnUpdate {
	if i != nil {
		cvu.SetVulnerabilityID(*i)
	}
	return cvu
}

// ClearVulnerabilityID clears the value of the "vulnerability_id" field.
func (cvu *CertifyVulnUpdate) ClearVulnerabilityID() *CertifyVulnUpdate {
	cvu.mutation.ClearVulnerabilityID()
	return cvu
}

// SetPackageID sets the "package_id" field.
func (cvu *CertifyVulnUpdate) SetPackageID(i int) *CertifyVulnUpdate {
	cvu.mutation.SetPackageID(i)
	return cvu
}

// SetTimeScanned sets the "time_scanned" field.
func (cvu *CertifyVulnUpdate) SetTimeScanned(t time.Time) *CertifyVulnUpdate {
	cvu.mutation.SetTimeScanned(t)
	return cvu
}

// SetDbURI sets the "db_uri" field.
func (cvu *CertifyVulnUpdate) SetDbURI(s string) *CertifyVulnUpdate {
	cvu.mutation.SetDbURI(s)
	return cvu
}

// SetDbVersion sets the "db_version" field.
func (cvu *CertifyVulnUpdate) SetDbVersion(s string) *CertifyVulnUpdate {
	cvu.mutation.SetDbVersion(s)
	return cvu
}

// SetScannerURI sets the "scanner_uri" field.
func (cvu *CertifyVulnUpdate) SetScannerURI(s string) *CertifyVulnUpdate {
	cvu.mutation.SetScannerURI(s)
	return cvu
}

// SetScannerVersion sets the "scanner_version" field.
func (cvu *CertifyVulnUpdate) SetScannerVersion(s string) *CertifyVulnUpdate {
	cvu.mutation.SetScannerVersion(s)
	return cvu
}

// SetOrigin sets the "origin" field.
func (cvu *CertifyVulnUpdate) SetOrigin(s string) *CertifyVulnUpdate {
	cvu.mutation.SetOrigin(s)
	return cvu
}

// SetCollector sets the "collector" field.
func (cvu *CertifyVulnUpdate) SetCollector(s string) *CertifyVulnUpdate {
	cvu.mutation.SetCollector(s)
	return cvu
}

// SetVulnerability sets the "vulnerability" edge to the Vulnerability entity.
func (cvu *CertifyVulnUpdate) SetVulnerability(v *Vulnerability) *CertifyVulnUpdate {
	return cvu.SetVulnerabilityID(v.ID)
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (cvu *CertifyVulnUpdate) SetPackage(p *PackageVersion) *CertifyVulnUpdate {
	return cvu.SetPackageID(p.ID)
}

// Mutation returns the CertifyVulnMutation object of the builder.
func (cvu *CertifyVulnUpdate) Mutation() *CertifyVulnMutation {
	return cvu.mutation
}

// ClearVulnerability clears the "vulnerability" edge to the Vulnerability entity.
func (cvu *CertifyVulnUpdate) ClearVulnerability() *CertifyVulnUpdate {
	cvu.mutation.ClearVulnerability()
	return cvu
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (cvu *CertifyVulnUpdate) ClearPackage() *CertifyVulnUpdate {
	cvu.mutation.ClearPackage()
	return cvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cvu *CertifyVulnUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cvu.sqlSave, cvu.mutation, cvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cvu *CertifyVulnUpdate) SaveX(ctx context.Context) int {
	affected, err := cvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cvu *CertifyVulnUpdate) Exec(ctx context.Context) error {
	_, err := cvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvu *CertifyVulnUpdate) ExecX(ctx context.Context) {
	if err := cvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cvu *CertifyVulnUpdate) check() error {
	if _, ok := cvu.mutation.PackageID(); cvu.mutation.PackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CertifyVuln.package"`)
	}
	return nil
}

func (cvu *CertifyVulnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(certifyvuln.Table, certifyvuln.Columns, sqlgraph.NewFieldSpec(certifyvuln.FieldID, field.TypeInt))
	if ps := cvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cvu.mutation.TimeScanned(); ok {
		_spec.SetField(certifyvuln.FieldTimeScanned, field.TypeTime, value)
	}
	if value, ok := cvu.mutation.DbURI(); ok {
		_spec.SetField(certifyvuln.FieldDbURI, field.TypeString, value)
	}
	if value, ok := cvu.mutation.DbVersion(); ok {
		_spec.SetField(certifyvuln.FieldDbVersion, field.TypeString, value)
	}
	if value, ok := cvu.mutation.ScannerURI(); ok {
		_spec.SetField(certifyvuln.FieldScannerURI, field.TypeString, value)
	}
	if value, ok := cvu.mutation.ScannerVersion(); ok {
		_spec.SetField(certifyvuln.FieldScannerVersion, field.TypeString, value)
	}
	if value, ok := cvu.mutation.Origin(); ok {
		_spec.SetField(certifyvuln.FieldOrigin, field.TypeString, value)
	}
	if value, ok := cvu.mutation.Collector(); ok {
		_spec.SetField(certifyvuln.FieldCollector, field.TypeString, value)
	}
	if cvu.mutation.VulnerabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.VulnerabilityTable,
			Columns: []string{certifyvuln.VulnerabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vulnerability.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cvu.mutation.VulnerabilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.VulnerabilityTable,
			Columns: []string{certifyvuln.VulnerabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vulnerability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cvu.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.PackageTable,
			Columns: []string{certifyvuln.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cvu.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.PackageTable,
			Columns: []string{certifyvuln.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certifyvuln.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cvu.mutation.done = true
	return n, nil
}

// CertifyVulnUpdateOne is the builder for updating a single CertifyVuln entity.
type CertifyVulnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CertifyVulnMutation
}

// SetVulnerabilityID sets the "vulnerability_id" field.
func (cvuo *CertifyVulnUpdateOne) SetVulnerabilityID(i int) *CertifyVulnUpdateOne {
	cvuo.mutation.SetVulnerabilityID(i)
	return cvuo
}

// SetNillableVulnerabilityID sets the "vulnerability_id" field if the given value is not nil.
func (cvuo *CertifyVulnUpdateOne) SetNillableVulnerabilityID(i *int) *CertifyVulnUpdateOne {
	if i != nil {
		cvuo.SetVulnerabilityID(*i)
	}
	return cvuo
}

// ClearVulnerabilityID clears the value of the "vulnerability_id" field.
func (cvuo *CertifyVulnUpdateOne) ClearVulnerabilityID() *CertifyVulnUpdateOne {
	cvuo.mutation.ClearVulnerabilityID()
	return cvuo
}

// SetPackageID sets the "package_id" field.
func (cvuo *CertifyVulnUpdateOne) SetPackageID(i int) *CertifyVulnUpdateOne {
	cvuo.mutation.SetPackageID(i)
	return cvuo
}

// SetTimeScanned sets the "time_scanned" field.
func (cvuo *CertifyVulnUpdateOne) SetTimeScanned(t time.Time) *CertifyVulnUpdateOne {
	cvuo.mutation.SetTimeScanned(t)
	return cvuo
}

// SetDbURI sets the "db_uri" field.
func (cvuo *CertifyVulnUpdateOne) SetDbURI(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetDbURI(s)
	return cvuo
}

// SetDbVersion sets the "db_version" field.
func (cvuo *CertifyVulnUpdateOne) SetDbVersion(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetDbVersion(s)
	return cvuo
}

// SetScannerURI sets the "scanner_uri" field.
func (cvuo *CertifyVulnUpdateOne) SetScannerURI(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetScannerURI(s)
	return cvuo
}

// SetScannerVersion sets the "scanner_version" field.
func (cvuo *CertifyVulnUpdateOne) SetScannerVersion(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetScannerVersion(s)
	return cvuo
}

// SetOrigin sets the "origin" field.
func (cvuo *CertifyVulnUpdateOne) SetOrigin(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetOrigin(s)
	return cvuo
}

// SetCollector sets the "collector" field.
func (cvuo *CertifyVulnUpdateOne) SetCollector(s string) *CertifyVulnUpdateOne {
	cvuo.mutation.SetCollector(s)
	return cvuo
}

// SetVulnerability sets the "vulnerability" edge to the Vulnerability entity.
func (cvuo *CertifyVulnUpdateOne) SetVulnerability(v *Vulnerability) *CertifyVulnUpdateOne {
	return cvuo.SetVulnerabilityID(v.ID)
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (cvuo *CertifyVulnUpdateOne) SetPackage(p *PackageVersion) *CertifyVulnUpdateOne {
	return cvuo.SetPackageID(p.ID)
}

// Mutation returns the CertifyVulnMutation object of the builder.
func (cvuo *CertifyVulnUpdateOne) Mutation() *CertifyVulnMutation {
	return cvuo.mutation
}

// ClearVulnerability clears the "vulnerability" edge to the Vulnerability entity.
func (cvuo *CertifyVulnUpdateOne) ClearVulnerability() *CertifyVulnUpdateOne {
	cvuo.mutation.ClearVulnerability()
	return cvuo
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (cvuo *CertifyVulnUpdateOne) ClearPackage() *CertifyVulnUpdateOne {
	cvuo.mutation.ClearPackage()
	return cvuo
}

// Where appends a list predicates to the CertifyVulnUpdate builder.
func (cvuo *CertifyVulnUpdateOne) Where(ps ...predicate.CertifyVuln) *CertifyVulnUpdateOne {
	cvuo.mutation.Where(ps...)
	return cvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cvuo *CertifyVulnUpdateOne) Select(field string, fields ...string) *CertifyVulnUpdateOne {
	cvuo.fields = append([]string{field}, fields...)
	return cvuo
}

// Save executes the query and returns the updated CertifyVuln entity.
func (cvuo *CertifyVulnUpdateOne) Save(ctx context.Context) (*CertifyVuln, error) {
	return withHooks(ctx, cvuo.sqlSave, cvuo.mutation, cvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cvuo *CertifyVulnUpdateOne) SaveX(ctx context.Context) *CertifyVuln {
	node, err := cvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cvuo *CertifyVulnUpdateOne) Exec(ctx context.Context) error {
	_, err := cvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvuo *CertifyVulnUpdateOne) ExecX(ctx context.Context) {
	if err := cvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cvuo *CertifyVulnUpdateOne) check() error {
	if _, ok := cvuo.mutation.PackageID(); cvuo.mutation.PackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CertifyVuln.package"`)
	}
	return nil
}

func (cvuo *CertifyVulnUpdateOne) sqlSave(ctx context.Context) (_node *CertifyVuln, err error) {
	if err := cvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(certifyvuln.Table, certifyvuln.Columns, sqlgraph.NewFieldSpec(certifyvuln.FieldID, field.TypeInt))
	id, ok := cvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CertifyVuln.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certifyvuln.FieldID)
		for _, f := range fields {
			if !certifyvuln.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != certifyvuln.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cvuo.mutation.TimeScanned(); ok {
		_spec.SetField(certifyvuln.FieldTimeScanned, field.TypeTime, value)
	}
	if value, ok := cvuo.mutation.DbURI(); ok {
		_spec.SetField(certifyvuln.FieldDbURI, field.TypeString, value)
	}
	if value, ok := cvuo.mutation.DbVersion(); ok {
		_spec.SetField(certifyvuln.FieldDbVersion, field.TypeString, value)
	}
	if value, ok := cvuo.mutation.ScannerURI(); ok {
		_spec.SetField(certifyvuln.FieldScannerURI, field.TypeString, value)
	}
	if value, ok := cvuo.mutation.ScannerVersion(); ok {
		_spec.SetField(certifyvuln.FieldScannerVersion, field.TypeString, value)
	}
	if value, ok := cvuo.mutation.Origin(); ok {
		_spec.SetField(certifyvuln.FieldOrigin, field.TypeString, value)
	}
	if value, ok := cvuo.mutation.Collector(); ok {
		_spec.SetField(certifyvuln.FieldCollector, field.TypeString, value)
	}
	if cvuo.mutation.VulnerabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.VulnerabilityTable,
			Columns: []string{certifyvuln.VulnerabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vulnerability.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cvuo.mutation.VulnerabilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.VulnerabilityTable,
			Columns: []string{certifyvuln.VulnerabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vulnerability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cvuo.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.PackageTable,
			Columns: []string{certifyvuln.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cvuo.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certifyvuln.PackageTable,
			Columns: []string{certifyvuln.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CertifyVuln{config: cvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certifyvuln.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cvuo.mutation.done = true
	return _node, nil
}
