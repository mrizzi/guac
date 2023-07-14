// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// BillOfMaterialsQuery is the builder for querying BillOfMaterials entities.
type BillOfMaterialsQuery struct {
	config
	ctx          *QueryContext
	order        []billofmaterials.OrderOption
	inters       []Interceptor
	predicates   []predicate.BillOfMaterials
	withPackage  *PackageVersionQuery
	withArtifact *ArtifactQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillOfMaterialsQuery builder.
func (bomq *BillOfMaterialsQuery) Where(ps ...predicate.BillOfMaterials) *BillOfMaterialsQuery {
	bomq.predicates = append(bomq.predicates, ps...)
	return bomq
}

// Limit the number of records to be returned by this query.
func (bomq *BillOfMaterialsQuery) Limit(limit int) *BillOfMaterialsQuery {
	bomq.ctx.Limit = &limit
	return bomq
}

// Offset to start from.
func (bomq *BillOfMaterialsQuery) Offset(offset int) *BillOfMaterialsQuery {
	bomq.ctx.Offset = &offset
	return bomq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bomq *BillOfMaterialsQuery) Unique(unique bool) *BillOfMaterialsQuery {
	bomq.ctx.Unique = &unique
	return bomq
}

// Order specifies how the records should be ordered.
func (bomq *BillOfMaterialsQuery) Order(o ...billofmaterials.OrderOption) *BillOfMaterialsQuery {
	bomq.order = append(bomq.order, o...)
	return bomq
}

// QueryPackage chains the current query on the "package" edge.
func (bomq *BillOfMaterialsQuery) QueryPackage() *PackageVersionQuery {
	query := (&PackageVersionClient{config: bomq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bomq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bomq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billofmaterials.Table, billofmaterials.FieldID, selector),
			sqlgraph.To(packageversion.Table, packageversion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, billofmaterials.PackageTable, billofmaterials.PackageColumn),
		)
		fromU = sqlgraph.SetNeighbors(bomq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArtifact chains the current query on the "artifact" edge.
func (bomq *BillOfMaterialsQuery) QueryArtifact() *ArtifactQuery {
	query := (&ArtifactClient{config: bomq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bomq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bomq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billofmaterials.Table, billofmaterials.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, billofmaterials.ArtifactTable, billofmaterials.ArtifactColumn),
		)
		fromU = sqlgraph.SetNeighbors(bomq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillOfMaterials entity from the query.
// Returns a *NotFoundError when no BillOfMaterials was found.
func (bomq *BillOfMaterialsQuery) First(ctx context.Context) (*BillOfMaterials, error) {
	nodes, err := bomq.Limit(1).All(setContextOp(ctx, bomq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billofmaterials.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) FirstX(ctx context.Context) *BillOfMaterials {
	node, err := bomq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillOfMaterials ID from the query.
// Returns a *NotFoundError when no BillOfMaterials ID was found.
func (bomq *BillOfMaterialsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bomq.Limit(1).IDs(setContextOp(ctx, bomq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billofmaterials.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) FirstIDX(ctx context.Context) int {
	id, err := bomq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillOfMaterials entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillOfMaterials entity is found.
// Returns a *NotFoundError when no BillOfMaterials entities are found.
func (bomq *BillOfMaterialsQuery) Only(ctx context.Context) (*BillOfMaterials, error) {
	nodes, err := bomq.Limit(2).All(setContextOp(ctx, bomq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billofmaterials.Label}
	default:
		return nil, &NotSingularError{billofmaterials.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) OnlyX(ctx context.Context) *BillOfMaterials {
	node, err := bomq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillOfMaterials ID in the query.
// Returns a *NotSingularError when more than one BillOfMaterials ID is found.
// Returns a *NotFoundError when no entities are found.
func (bomq *BillOfMaterialsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bomq.Limit(2).IDs(setContextOp(ctx, bomq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billofmaterials.Label}
	default:
		err = &NotSingularError{billofmaterials.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) OnlyIDX(ctx context.Context) int {
	id, err := bomq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillOfMaterialsSlice.
func (bomq *BillOfMaterialsQuery) All(ctx context.Context) ([]*BillOfMaterials, error) {
	ctx = setContextOp(ctx, bomq.ctx, "All")
	if err := bomq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillOfMaterials, *BillOfMaterialsQuery]()
	return withInterceptors[[]*BillOfMaterials](ctx, bomq, qr, bomq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) AllX(ctx context.Context) []*BillOfMaterials {
	nodes, err := bomq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillOfMaterials IDs.
func (bomq *BillOfMaterialsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bomq.ctx.Unique == nil && bomq.path != nil {
		bomq.Unique(true)
	}
	ctx = setContextOp(ctx, bomq.ctx, "IDs")
	if err = bomq.Select(billofmaterials.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) IDsX(ctx context.Context) []int {
	ids, err := bomq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bomq *BillOfMaterialsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bomq.ctx, "Count")
	if err := bomq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bomq, querierCount[*BillOfMaterialsQuery](), bomq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) CountX(ctx context.Context) int {
	count, err := bomq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bomq *BillOfMaterialsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bomq.ctx, "Exist")
	switch _, err := bomq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bomq *BillOfMaterialsQuery) ExistX(ctx context.Context) bool {
	exist, err := bomq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillOfMaterialsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bomq *BillOfMaterialsQuery) Clone() *BillOfMaterialsQuery {
	if bomq == nil {
		return nil
	}
	return &BillOfMaterialsQuery{
		config:       bomq.config,
		ctx:          bomq.ctx.Clone(),
		order:        append([]billofmaterials.OrderOption{}, bomq.order...),
		inters:       append([]Interceptor{}, bomq.inters...),
		predicates:   append([]predicate.BillOfMaterials{}, bomq.predicates...),
		withPackage:  bomq.withPackage.Clone(),
		withArtifact: bomq.withArtifact.Clone(),
		// clone intermediate query.
		sql:  bomq.sql.Clone(),
		path: bomq.path,
	}
}

// WithPackage tells the query-builder to eager-load the nodes that are connected to
// the "package" edge. The optional arguments are used to configure the query builder of the edge.
func (bomq *BillOfMaterialsQuery) WithPackage(opts ...func(*PackageVersionQuery)) *BillOfMaterialsQuery {
	query := (&PackageVersionClient{config: bomq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bomq.withPackage = query
	return bomq
}

// WithArtifact tells the query-builder to eager-load the nodes that are connected to
// the "artifact" edge. The optional arguments are used to configure the query builder of the edge.
func (bomq *BillOfMaterialsQuery) WithArtifact(opts ...func(*ArtifactQuery)) *BillOfMaterialsQuery {
	query := (&ArtifactClient{config: bomq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bomq.withArtifact = query
	return bomq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PackageID int `json:"package_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillOfMaterials.Query().
//		GroupBy(billofmaterials.FieldPackageID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bomq *BillOfMaterialsQuery) GroupBy(field string, fields ...string) *BillOfMaterialsGroupBy {
	bomq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillOfMaterialsGroupBy{build: bomq}
	grbuild.flds = &bomq.ctx.Fields
	grbuild.label = billofmaterials.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PackageID int `json:"package_id,omitempty"`
//	}
//
//	client.BillOfMaterials.Query().
//		Select(billofmaterials.FieldPackageID).
//		Scan(ctx, &v)
func (bomq *BillOfMaterialsQuery) Select(fields ...string) *BillOfMaterialsSelect {
	bomq.ctx.Fields = append(bomq.ctx.Fields, fields...)
	sbuild := &BillOfMaterialsSelect{BillOfMaterialsQuery: bomq}
	sbuild.label = billofmaterials.Label
	sbuild.flds, sbuild.scan = &bomq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillOfMaterialsSelect configured with the given aggregations.
func (bomq *BillOfMaterialsQuery) Aggregate(fns ...AggregateFunc) *BillOfMaterialsSelect {
	return bomq.Select().Aggregate(fns...)
}

func (bomq *BillOfMaterialsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bomq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bomq); err != nil {
				return err
			}
		}
	}
	for _, f := range bomq.ctx.Fields {
		if !billofmaterials.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bomq.path != nil {
		prev, err := bomq.path(ctx)
		if err != nil {
			return err
		}
		bomq.sql = prev
	}
	return nil
}

func (bomq *BillOfMaterialsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillOfMaterials, error) {
	var (
		nodes       = []*BillOfMaterials{}
		_spec       = bomq.querySpec()
		loadedTypes = [2]bool{
			bomq.withPackage != nil,
			bomq.withArtifact != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillOfMaterials).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillOfMaterials{config: bomq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bomq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bomq.withPackage; query != nil {
		if err := bomq.loadPackage(ctx, query, nodes, nil,
			func(n *BillOfMaterials, e *PackageVersion) { n.Edges.Package = e }); err != nil {
			return nil, err
		}
	}
	if query := bomq.withArtifact; query != nil {
		if err := bomq.loadArtifact(ctx, query, nodes, nil,
			func(n *BillOfMaterials, e *Artifact) { n.Edges.Artifact = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bomq *BillOfMaterialsQuery) loadPackage(ctx context.Context, query *PackageVersionQuery, nodes []*BillOfMaterials, init func(*BillOfMaterials), assign func(*BillOfMaterials, *PackageVersion)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BillOfMaterials)
	for i := range nodes {
		if nodes[i].PackageID == nil {
			continue
		}
		fk := *nodes[i].PackageID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(packageversion.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "package_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bomq *BillOfMaterialsQuery) loadArtifact(ctx context.Context, query *ArtifactQuery, nodes []*BillOfMaterials, init func(*BillOfMaterials), assign func(*BillOfMaterials, *Artifact)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BillOfMaterials)
	for i := range nodes {
		if nodes[i].ArtifactID == nil {
			continue
		}
		fk := *nodes[i].ArtifactID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(artifact.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "artifact_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bomq *BillOfMaterialsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bomq.querySpec()
	_spec.Node.Columns = bomq.ctx.Fields
	if len(bomq.ctx.Fields) > 0 {
		_spec.Unique = bomq.ctx.Unique != nil && *bomq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bomq.driver, _spec)
}

func (bomq *BillOfMaterialsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billofmaterials.Table, billofmaterials.Columns, sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeInt))
	_spec.From = bomq.sql
	if unique := bomq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bomq.path != nil {
		_spec.Unique = true
	}
	if fields := bomq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billofmaterials.FieldID)
		for i := range fields {
			if fields[i] != billofmaterials.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bomq.withPackage != nil {
			_spec.Node.AddColumnOnce(billofmaterials.FieldPackageID)
		}
		if bomq.withArtifact != nil {
			_spec.Node.AddColumnOnce(billofmaterials.FieldArtifactID)
		}
	}
	if ps := bomq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bomq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bomq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bomq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bomq *BillOfMaterialsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bomq.driver.Dialect())
	t1 := builder.Table(billofmaterials.Table)
	columns := bomq.ctx.Fields
	if len(columns) == 0 {
		columns = billofmaterials.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bomq.sql != nil {
		selector = bomq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bomq.ctx.Unique != nil && *bomq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bomq.predicates {
		p(selector)
	}
	for _, p := range bomq.order {
		p(selector)
	}
	if offset := bomq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bomq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BillOfMaterialsGroupBy is the group-by builder for BillOfMaterials entities.
type BillOfMaterialsGroupBy struct {
	selector
	build *BillOfMaterialsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bomgb *BillOfMaterialsGroupBy) Aggregate(fns ...AggregateFunc) *BillOfMaterialsGroupBy {
	bomgb.fns = append(bomgb.fns, fns...)
	return bomgb
}

// Scan applies the selector query and scans the result into the given value.
func (bomgb *BillOfMaterialsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bomgb.build.ctx, "GroupBy")
	if err := bomgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillOfMaterialsQuery, *BillOfMaterialsGroupBy](ctx, bomgb.build, bomgb, bomgb.build.inters, v)
}

func (bomgb *BillOfMaterialsGroupBy) sqlScan(ctx context.Context, root *BillOfMaterialsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bomgb.fns))
	for _, fn := range bomgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bomgb.flds)+len(bomgb.fns))
		for _, f := range *bomgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bomgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bomgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillOfMaterialsSelect is the builder for selecting fields of BillOfMaterials entities.
type BillOfMaterialsSelect struct {
	*BillOfMaterialsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (boms *BillOfMaterialsSelect) Aggregate(fns ...AggregateFunc) *BillOfMaterialsSelect {
	boms.fns = append(boms.fns, fns...)
	return boms
}

// Scan applies the selector query and scans the result into the given value.
func (boms *BillOfMaterialsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, boms.ctx, "Select")
	if err := boms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillOfMaterialsQuery, *BillOfMaterialsSelect](ctx, boms.BillOfMaterialsQuery, boms, boms.inters, v)
}

func (boms *BillOfMaterialsSelect) sqlScan(ctx context.Context, root *BillOfMaterialsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(boms.fns))
	for _, fn := range boms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*boms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := boms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}