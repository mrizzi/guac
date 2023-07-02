// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/buildernode"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// BuilderNodeQuery is the builder for querying BuilderNode entities.
type BuilderNodeQuery struct {
	config
	ctx        *QueryContext
	order      []buildernode.OrderOption
	inters     []Interceptor
	predicates []predicate.BuilderNode
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BuilderNodeQuery builder.
func (bnq *BuilderNodeQuery) Where(ps ...predicate.BuilderNode) *BuilderNodeQuery {
	bnq.predicates = append(bnq.predicates, ps...)
	return bnq
}

// Limit the number of records to be returned by this query.
func (bnq *BuilderNodeQuery) Limit(limit int) *BuilderNodeQuery {
	bnq.ctx.Limit = &limit
	return bnq
}

// Offset to start from.
func (bnq *BuilderNodeQuery) Offset(offset int) *BuilderNodeQuery {
	bnq.ctx.Offset = &offset
	return bnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bnq *BuilderNodeQuery) Unique(unique bool) *BuilderNodeQuery {
	bnq.ctx.Unique = &unique
	return bnq
}

// Order specifies how the records should be ordered.
func (bnq *BuilderNodeQuery) Order(o ...buildernode.OrderOption) *BuilderNodeQuery {
	bnq.order = append(bnq.order, o...)
	return bnq
}

// First returns the first BuilderNode entity from the query.
// Returns a *NotFoundError when no BuilderNode was found.
func (bnq *BuilderNodeQuery) First(ctx context.Context) (*BuilderNode, error) {
	nodes, err := bnq.Limit(1).All(setContextOp(ctx, bnq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{buildernode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bnq *BuilderNodeQuery) FirstX(ctx context.Context) *BuilderNode {
	node, err := bnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BuilderNode ID from the query.
// Returns a *NotFoundError when no BuilderNode ID was found.
func (bnq *BuilderNodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bnq.Limit(1).IDs(setContextOp(ctx, bnq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{buildernode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bnq *BuilderNodeQuery) FirstIDX(ctx context.Context) int {
	id, err := bnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BuilderNode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BuilderNode entity is found.
// Returns a *NotFoundError when no BuilderNode entities are found.
func (bnq *BuilderNodeQuery) Only(ctx context.Context) (*BuilderNode, error) {
	nodes, err := bnq.Limit(2).All(setContextOp(ctx, bnq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{buildernode.Label}
	default:
		return nil, &NotSingularError{buildernode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bnq *BuilderNodeQuery) OnlyX(ctx context.Context) *BuilderNode {
	node, err := bnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BuilderNode ID in the query.
// Returns a *NotSingularError when more than one BuilderNode ID is found.
// Returns a *NotFoundError when no entities are found.
func (bnq *BuilderNodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bnq.Limit(2).IDs(setContextOp(ctx, bnq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{buildernode.Label}
	default:
		err = &NotSingularError{buildernode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bnq *BuilderNodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := bnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BuilderNodes.
func (bnq *BuilderNodeQuery) All(ctx context.Context) ([]*BuilderNode, error) {
	ctx = setContextOp(ctx, bnq.ctx, "All")
	if err := bnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BuilderNode, *BuilderNodeQuery]()
	return withInterceptors[[]*BuilderNode](ctx, bnq, qr, bnq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bnq *BuilderNodeQuery) AllX(ctx context.Context) []*BuilderNode {
	nodes, err := bnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BuilderNode IDs.
func (bnq *BuilderNodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bnq.ctx.Unique == nil && bnq.path != nil {
		bnq.Unique(true)
	}
	ctx = setContextOp(ctx, bnq.ctx, "IDs")
	if err = bnq.Select(buildernode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bnq *BuilderNodeQuery) IDsX(ctx context.Context) []int {
	ids, err := bnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bnq *BuilderNodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bnq.ctx, "Count")
	if err := bnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bnq, querierCount[*BuilderNodeQuery](), bnq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bnq *BuilderNodeQuery) CountX(ctx context.Context) int {
	count, err := bnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bnq *BuilderNodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bnq.ctx, "Exist")
	switch _, err := bnq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bnq *BuilderNodeQuery) ExistX(ctx context.Context) bool {
	exist, err := bnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BuilderNodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bnq *BuilderNodeQuery) Clone() *BuilderNodeQuery {
	if bnq == nil {
		return nil
	}
	return &BuilderNodeQuery{
		config:     bnq.config,
		ctx:        bnq.ctx.Clone(),
		order:      append([]buildernode.OrderOption{}, bnq.order...),
		inters:     append([]Interceptor{}, bnq.inters...),
		predicates: append([]predicate.BuilderNode{}, bnq.predicates...),
		// clone intermediate query.
		sql:  bnq.sql.Clone(),
		path: bnq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URI string `json:"uri,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BuilderNode.Query().
//		GroupBy(buildernode.FieldURI).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bnq *BuilderNodeQuery) GroupBy(field string, fields ...string) *BuilderNodeGroupBy {
	bnq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BuilderNodeGroupBy{build: bnq}
	grbuild.flds = &bnq.ctx.Fields
	grbuild.label = buildernode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URI string `json:"uri,omitempty"`
//	}
//
//	client.BuilderNode.Query().
//		Select(buildernode.FieldURI).
//		Scan(ctx, &v)
func (bnq *BuilderNodeQuery) Select(fields ...string) *BuilderNodeSelect {
	bnq.ctx.Fields = append(bnq.ctx.Fields, fields...)
	sbuild := &BuilderNodeSelect{BuilderNodeQuery: bnq}
	sbuild.label = buildernode.Label
	sbuild.flds, sbuild.scan = &bnq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BuilderNodeSelect configured with the given aggregations.
func (bnq *BuilderNodeQuery) Aggregate(fns ...AggregateFunc) *BuilderNodeSelect {
	return bnq.Select().Aggregate(fns...)
}

func (bnq *BuilderNodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bnq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bnq); err != nil {
				return err
			}
		}
	}
	for _, f := range bnq.ctx.Fields {
		if !buildernode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bnq.path != nil {
		prev, err := bnq.path(ctx)
		if err != nil {
			return err
		}
		bnq.sql = prev
	}
	return nil
}

func (bnq *BuilderNodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BuilderNode, error) {
	var (
		nodes = []*BuilderNode{}
		_spec = bnq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BuilderNode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BuilderNode{config: bnq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bnq *BuilderNodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bnq.querySpec()
	_spec.Node.Columns = bnq.ctx.Fields
	if len(bnq.ctx.Fields) > 0 {
		_spec.Unique = bnq.ctx.Unique != nil && *bnq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bnq.driver, _spec)
}

func (bnq *BuilderNodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(buildernode.Table, buildernode.Columns, sqlgraph.NewFieldSpec(buildernode.FieldID, field.TypeInt))
	_spec.From = bnq.sql
	if unique := bnq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bnq.path != nil {
		_spec.Unique = true
	}
	if fields := bnq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, buildernode.FieldID)
		for i := range fields {
			if fields[i] != buildernode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bnq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bnq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bnq *BuilderNodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bnq.driver.Dialect())
	t1 := builder.Table(buildernode.Table)
	columns := bnq.ctx.Fields
	if len(columns) == 0 {
		columns = buildernode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bnq.sql != nil {
		selector = bnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bnq.ctx.Unique != nil && *bnq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bnq.predicates {
		p(selector)
	}
	for _, p := range bnq.order {
		p(selector)
	}
	if offset := bnq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bnq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BuilderNodeGroupBy is the group-by builder for BuilderNode entities.
type BuilderNodeGroupBy struct {
	selector
	build *BuilderNodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bngb *BuilderNodeGroupBy) Aggregate(fns ...AggregateFunc) *BuilderNodeGroupBy {
	bngb.fns = append(bngb.fns, fns...)
	return bngb
}

// Scan applies the selector query and scans the result into the given value.
func (bngb *BuilderNodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bngb.build.ctx, "GroupBy")
	if err := bngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderNodeQuery, *BuilderNodeGroupBy](ctx, bngb.build, bngb, bngb.build.inters, v)
}

func (bngb *BuilderNodeGroupBy) sqlScan(ctx context.Context, root *BuilderNodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bngb.fns))
	for _, fn := range bngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bngb.flds)+len(bngb.fns))
		for _, f := range *bngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BuilderNodeSelect is the builder for selecting fields of BuilderNode entities.
type BuilderNodeSelect struct {
	*BuilderNodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bns *BuilderNodeSelect) Aggregate(fns ...AggregateFunc) *BuilderNodeSelect {
	bns.fns = append(bns.fns, fns...)
	return bns
}

// Scan applies the selector query and scans the result into the given value.
func (bns *BuilderNodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bns.ctx, "Select")
	if err := bns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderNodeQuery, *BuilderNodeSelect](ctx, bns.BuilderNodeQuery, bns, bns.inters, v)
}

func (bns *BuilderNodeSelect) sqlScan(ctx context.Context, root *BuilderNodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bns.fns))
	for _, fn := range bns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}