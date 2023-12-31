// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/division"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
)

// DivisionQuery is the builder for querying Division entities.
type DivisionQuery struct {
	config
	ctx        *QueryContext
	order      []division.OrderOption
	inters     []Interceptor
	predicates []predicate.Division
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DivisionQuery builder.
func (dq *DivisionQuery) Where(ps ...predicate.Division) *DivisionQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DivisionQuery) Limit(limit int) *DivisionQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DivisionQuery) Offset(offset int) *DivisionQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DivisionQuery) Unique(unique bool) *DivisionQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DivisionQuery) Order(o ...division.OrderOption) *DivisionQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Division entity from the query.
// Returns a *NotFoundError when no Division was found.
func (dq *DivisionQuery) First(ctx context.Context) (*Division, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{division.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DivisionQuery) FirstX(ctx context.Context) *Division {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Division ID from the query.
// Returns a *NotFoundError when no Division ID was found.
func (dq *DivisionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{division.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DivisionQuery) FirstIDX(ctx context.Context) string {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Division entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Division entity is found.
// Returns a *NotFoundError when no Division entities are found.
func (dq *DivisionQuery) Only(ctx context.Context) (*Division, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{division.Label}
	default:
		return nil, &NotSingularError{division.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DivisionQuery) OnlyX(ctx context.Context) *Division {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Division ID in the query.
// Returns a *NotSingularError when more than one Division ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DivisionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{division.Label}
	default:
		err = &NotSingularError{division.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DivisionQuery) OnlyIDX(ctx context.Context) string {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Divisions.
func (dq *DivisionQuery) All(ctx context.Context) ([]*Division, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Division, *DivisionQuery]()
	return withInterceptors[[]*Division](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DivisionQuery) AllX(ctx context.Context) []*Division {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Division IDs.
func (dq *DivisionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(division.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DivisionQuery) IDsX(ctx context.Context) []string {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DivisionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DivisionQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DivisionQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DivisionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DivisionQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DivisionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DivisionQuery) Clone() *DivisionQuery {
	if dq == nil {
		return nil
	}
	return &DivisionQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]division.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Division{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Division.Query().
//		GroupBy(division.FieldName).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (dq *DivisionQuery) GroupBy(field string, fields ...string) *DivisionGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DivisionGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = division.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Division.Query().
//		Select(division.FieldName).
//		Scan(ctx, &v)
func (dq *DivisionQuery) Select(fields ...string) *DivisionSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DivisionSelect{DivisionQuery: dq}
	sbuild.label = division.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DivisionSelect configured with the given aggregations.
func (dq *DivisionQuery) Aggregate(fns ...AggregateFunc) *DivisionSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DivisionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !division.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DivisionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Division, error) {
	var (
		nodes = []*Division{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Division).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Division{config: dq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dq *DivisionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DivisionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(division.Table, division.Columns, sqlgraph.NewFieldSpec(division.FieldID, field.TypeString))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, division.FieldID)
		for i := range fields {
			if fields[i] != division.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DivisionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(division.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = division.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DivisionGroupBy is the group-by builder for Division entities.
type DivisionGroupBy struct {
	selector
	build *DivisionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DivisionGroupBy) Aggregate(fns ...AggregateFunc) *DivisionGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DivisionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionQuery, *DivisionGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DivisionGroupBy) sqlScan(ctx context.Context, root *DivisionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DivisionSelect is the builder for selecting fields of Division entities.
type DivisionSelect struct {
	*DivisionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DivisionSelect) Aggregate(fns ...AggregateFunc) *DivisionSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DivisionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionQuery, *DivisionSelect](ctx, ds.DivisionQuery, ds, ds.inters, v)
}

func (ds *DivisionSelect) sqlScan(ctx context.Context, root *DivisionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
