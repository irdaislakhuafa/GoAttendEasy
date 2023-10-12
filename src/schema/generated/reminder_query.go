// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/reminder"
)

// ReminderQuery is the builder for querying Reminder entities.
type ReminderQuery struct {
	config
	ctx        *QueryContext
	order      []reminder.OrderOption
	inters     []Interceptor
	predicates []predicate.Reminder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReminderQuery builder.
func (rq *ReminderQuery) Where(ps ...predicate.Reminder) *ReminderQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReminderQuery) Limit(limit int) *ReminderQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReminderQuery) Offset(offset int) *ReminderQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReminderQuery) Unique(unique bool) *ReminderQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReminderQuery) Order(o ...reminder.OrderOption) *ReminderQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// First returns the first Reminder entity from the query.
// Returns a *NotFoundError when no Reminder was found.
func (rq *ReminderQuery) First(ctx context.Context) (*Reminder, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reminder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReminderQuery) FirstX(ctx context.Context) *Reminder {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Reminder ID from the query.
// Returns a *NotFoundError when no Reminder ID was found.
func (rq *ReminderQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reminder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReminderQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Reminder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Reminder entity is found.
// Returns a *NotFoundError when no Reminder entities are found.
func (rq *ReminderQuery) Only(ctx context.Context) (*Reminder, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reminder.Label}
	default:
		return nil, &NotSingularError{reminder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReminderQuery) OnlyX(ctx context.Context) *Reminder {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Reminder ID in the query.
// Returns a *NotSingularError when more than one Reminder ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReminderQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reminder.Label}
	default:
		err = &NotSingularError{reminder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReminderQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Reminders.
func (rq *ReminderQuery) All(ctx context.Context) ([]*Reminder, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Reminder, *ReminderQuery]()
	return withInterceptors[[]*Reminder](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReminderQuery) AllX(ctx context.Context) []*Reminder {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Reminder IDs.
func (rq *ReminderQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(reminder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReminderQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReminderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReminderQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReminderQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReminderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReminderQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReminderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReminderQuery) Clone() *ReminderQuery {
	if rq == nil {
		return nil
	}
	return &ReminderQuery{
		config:     rq.config,
		ctx:        rq.ctx.Clone(),
		order:      append([]reminder.OrderOption{}, rq.order...),
		inters:     append([]Interceptor{}, rq.inters...),
		predicates: append([]predicate.Reminder{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		In time.Time `json:"in,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Reminder.Query().
//		GroupBy(reminder.FieldIn).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (rq *ReminderQuery) GroupBy(field string, fields ...string) *ReminderGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReminderGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = reminder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		In time.Time `json:"in,omitempty"`
//	}
//
//	client.Reminder.Query().
//		Select(reminder.FieldIn).
//		Scan(ctx, &v)
func (rq *ReminderQuery) Select(fields ...string) *ReminderSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReminderSelect{ReminderQuery: rq}
	sbuild.label = reminder.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReminderSelect configured with the given aggregations.
func (rq *ReminderQuery) Aggregate(fns ...AggregateFunc) *ReminderSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReminderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !reminder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReminderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Reminder, error) {
	var (
		nodes = []*Reminder{}
		_spec = rq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Reminder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Reminder{config: rq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rq *ReminderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReminderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reminder.Table, reminder.Columns, sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reminder.FieldID)
		for i := range fields {
			if fields[i] != reminder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReminderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(reminder.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = reminder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReminderGroupBy is the group-by builder for Reminder entities.
type ReminderGroupBy struct {
	selector
	build *ReminderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReminderGroupBy) Aggregate(fns ...AggregateFunc) *ReminderGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReminderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReminderQuery, *ReminderGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReminderGroupBy) sqlScan(ctx context.Context, root *ReminderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReminderSelect is the builder for selecting fields of Reminder entities.
type ReminderSelect struct {
	*ReminderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReminderSelect) Aggregate(fns ...AggregateFunc) *ReminderSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReminderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReminderQuery, *ReminderSelect](ctx, rs.ReminderQuery, rs, rs.inters, v)
}

func (rs *ReminderSelect) sqlScan(ctx context.Context, root *ReminderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
