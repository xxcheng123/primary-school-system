// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xxcheng123/primary-school-system/ent/predicate"
	"github.com/xxcheng123/primary-school-system/ent/xconfig"
)

// XConfigQuery is the builder for querying XConfig entities.
type XConfigQuery struct {
	config
	ctx        *QueryContext
	order      []xconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.XConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the XConfigQuery builder.
func (xq *XConfigQuery) Where(ps ...predicate.XConfig) *XConfigQuery {
	xq.predicates = append(xq.predicates, ps...)
	return xq
}

// Limit the number of records to be returned by this query.
func (xq *XConfigQuery) Limit(limit int) *XConfigQuery {
	xq.ctx.Limit = &limit
	return xq
}

// Offset to start from.
func (xq *XConfigQuery) Offset(offset int) *XConfigQuery {
	xq.ctx.Offset = &offset
	return xq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (xq *XConfigQuery) Unique(unique bool) *XConfigQuery {
	xq.ctx.Unique = &unique
	return xq
}

// Order specifies how the records should be ordered.
func (xq *XConfigQuery) Order(o ...xconfig.OrderOption) *XConfigQuery {
	xq.order = append(xq.order, o...)
	return xq
}

// First returns the first XConfig entity from the query.
// Returns a *NotFoundError when no XConfig was found.
func (xq *XConfigQuery) First(ctx context.Context) (*XConfig, error) {
	nodes, err := xq.Limit(1).All(setContextOp(ctx, xq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{xconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (xq *XConfigQuery) FirstX(ctx context.Context) *XConfig {
	node, err := xq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first XConfig ID from the query.
// Returns a *NotFoundError when no XConfig ID was found.
func (xq *XConfigQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = xq.Limit(1).IDs(setContextOp(ctx, xq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{xconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (xq *XConfigQuery) FirstIDX(ctx context.Context) int64 {
	id, err := xq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single XConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one XConfig entity is found.
// Returns a *NotFoundError when no XConfig entities are found.
func (xq *XConfigQuery) Only(ctx context.Context) (*XConfig, error) {
	nodes, err := xq.Limit(2).All(setContextOp(ctx, xq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{xconfig.Label}
	default:
		return nil, &NotSingularError{xconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (xq *XConfigQuery) OnlyX(ctx context.Context) *XConfig {
	node, err := xq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only XConfig ID in the query.
// Returns a *NotSingularError when more than one XConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (xq *XConfigQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = xq.Limit(2).IDs(setContextOp(ctx, xq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{xconfig.Label}
	default:
		err = &NotSingularError{xconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (xq *XConfigQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := xq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of XConfigs.
func (xq *XConfigQuery) All(ctx context.Context) ([]*XConfig, error) {
	ctx = setContextOp(ctx, xq.ctx, "All")
	if err := xq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*XConfig, *XConfigQuery]()
	return withInterceptors[[]*XConfig](ctx, xq, qr, xq.inters)
}

// AllX is like All, but panics if an error occurs.
func (xq *XConfigQuery) AllX(ctx context.Context) []*XConfig {
	nodes, err := xq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of XConfig IDs.
func (xq *XConfigQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if xq.ctx.Unique == nil && xq.path != nil {
		xq.Unique(true)
	}
	ctx = setContextOp(ctx, xq.ctx, "IDs")
	if err = xq.Select(xconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (xq *XConfigQuery) IDsX(ctx context.Context) []int64 {
	ids, err := xq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (xq *XConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, xq.ctx, "Count")
	if err := xq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, xq, querierCount[*XConfigQuery](), xq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (xq *XConfigQuery) CountX(ctx context.Context) int {
	count, err := xq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (xq *XConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, xq.ctx, "Exist")
	switch _, err := xq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (xq *XConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := xq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the XConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (xq *XConfigQuery) Clone() *XConfigQuery {
	if xq == nil {
		return nil
	}
	return &XConfigQuery{
		config:     xq.config,
		ctx:        xq.ctx.Clone(),
		order:      append([]xconfig.OrderOption{}, xq.order...),
		inters:     append([]Interceptor{}, xq.inters...),
		predicates: append([]predicate.XConfig{}, xq.predicates...),
		// clone intermediate query.
		sql:  xq.sql.Clone(),
		path: xq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.XConfig.Query().
//		GroupBy(xconfig.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (xq *XConfigQuery) GroupBy(field string, fields ...string) *XConfigGroupBy {
	xq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &XConfigGroupBy{build: xq}
	grbuild.flds = &xq.ctx.Fields
	grbuild.label = xconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.XConfig.Query().
//		Select(xconfig.FieldCreateTime).
//		Scan(ctx, &v)
func (xq *XConfigQuery) Select(fields ...string) *XConfigSelect {
	xq.ctx.Fields = append(xq.ctx.Fields, fields...)
	sbuild := &XConfigSelect{XConfigQuery: xq}
	sbuild.label = xconfig.Label
	sbuild.flds, sbuild.scan = &xq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a XConfigSelect configured with the given aggregations.
func (xq *XConfigQuery) Aggregate(fns ...AggregateFunc) *XConfigSelect {
	return xq.Select().Aggregate(fns...)
}

func (xq *XConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range xq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, xq); err != nil {
				return err
			}
		}
	}
	for _, f := range xq.ctx.Fields {
		if !xconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if xq.path != nil {
		prev, err := xq.path(ctx)
		if err != nil {
			return err
		}
		xq.sql = prev
	}
	return nil
}

func (xq *XConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*XConfig, error) {
	var (
		nodes = []*XConfig{}
		_spec = xq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*XConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &XConfig{config: xq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, xq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (xq *XConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := xq.querySpec()
	_spec.Node.Columns = xq.ctx.Fields
	if len(xq.ctx.Fields) > 0 {
		_spec.Unique = xq.ctx.Unique != nil && *xq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, xq.driver, _spec)
}

func (xq *XConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(xconfig.Table, xconfig.Columns, sqlgraph.NewFieldSpec(xconfig.FieldID, field.TypeInt64))
	_spec.From = xq.sql
	if unique := xq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if xq.path != nil {
		_spec.Unique = true
	}
	if fields := xq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xconfig.FieldID)
		for i := range fields {
			if fields[i] != xconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := xq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := xq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := xq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := xq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (xq *XConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(xq.driver.Dialect())
	t1 := builder.Table(xconfig.Table)
	columns := xq.ctx.Fields
	if len(columns) == 0 {
		columns = xconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if xq.sql != nil {
		selector = xq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if xq.ctx.Unique != nil && *xq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range xq.predicates {
		p(selector)
	}
	for _, p := range xq.order {
		p(selector)
	}
	if offset := xq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := xq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// XConfigGroupBy is the group-by builder for XConfig entities.
type XConfigGroupBy struct {
	selector
	build *XConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (xgb *XConfigGroupBy) Aggregate(fns ...AggregateFunc) *XConfigGroupBy {
	xgb.fns = append(xgb.fns, fns...)
	return xgb
}

// Scan applies the selector query and scans the result into the given value.
func (xgb *XConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xgb.build.ctx, "GroupBy")
	if err := xgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XConfigQuery, *XConfigGroupBy](ctx, xgb.build, xgb, xgb.build.inters, v)
}

func (xgb *XConfigGroupBy) sqlScan(ctx context.Context, root *XConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(xgb.fns))
	for _, fn := range xgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*xgb.flds)+len(xgb.fns))
		for _, f := range *xgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*xgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// XConfigSelect is the builder for selecting fields of XConfig entities.
type XConfigSelect struct {
	*XConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (xs *XConfigSelect) Aggregate(fns ...AggregateFunc) *XConfigSelect {
	xs.fns = append(xs.fns, fns...)
	return xs
}

// Scan applies the selector query and scans the result into the given value.
func (xs *XConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xs.ctx, "Select")
	if err := xs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XConfigQuery, *XConfigSelect](ctx, xs.XConfigQuery, xs, xs.inters, v)
}

func (xs *XConfigSelect) sqlScan(ctx context.Context, root *XConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(xs.fns))
	for _, fn := range xs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*xs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
