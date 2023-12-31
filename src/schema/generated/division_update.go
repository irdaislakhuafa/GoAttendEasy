// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/division"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
)

// DivisionUpdate is the builder for updating Division entities.
type DivisionUpdate struct {
	config
	hooks    []Hook
	mutation *DivisionMutation
}

// Where appends a list predicates to the DivisionUpdate builder.
func (du *DivisionUpdate) Where(ps ...predicate.Division) *DivisionUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DivisionUpdate) SetName(s string) *DivisionUpdate {
	du.mutation.SetName(s)
	return du
}

// SetDescription sets the "description" field.
func (du *DivisionUpdate) SetDescription(s string) *DivisionUpdate {
	du.mutation.SetDescription(s)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DivisionUpdate) SetCreatedAt(t time.Time) *DivisionUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableCreatedAt(t *time.Time) *DivisionUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// ClearCreatedAt clears the value of the "created_at" field.
func (du *DivisionUpdate) ClearCreatedAt() *DivisionUpdate {
	du.mutation.ClearCreatedAt()
	return du
}

// SetCreatedBy sets the "created_by" field.
func (du *DivisionUpdate) SetCreatedBy(s string) *DivisionUpdate {
	du.mutation.SetCreatedBy(s)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DivisionUpdate) SetUpdatedAt(t time.Time) *DivisionUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableUpdatedAt(t *time.Time) *DivisionUpdate {
	if t != nil {
		du.SetUpdatedAt(*t)
	}
	return du
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (du *DivisionUpdate) ClearUpdatedAt() *DivisionUpdate {
	du.mutation.ClearUpdatedAt()
	return du
}

// SetUpdatedBy sets the "updated_by" field.
func (du *DivisionUpdate) SetUpdatedBy(s string) *DivisionUpdate {
	du.mutation.SetUpdatedBy(s)
	return du
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableUpdatedBy(s *string) *DivisionUpdate {
	if s != nil {
		du.SetUpdatedBy(*s)
	}
	return du
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (du *DivisionUpdate) ClearUpdatedBy() *DivisionUpdate {
	du.mutation.ClearUpdatedBy()
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DivisionUpdate) SetDeletedAt(t time.Time) *DivisionUpdate {
	du.mutation.SetDeletedAt(t)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableDeletedAt(t *time.Time) *DivisionUpdate {
	if t != nil {
		du.SetDeletedAt(*t)
	}
	return du
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (du *DivisionUpdate) ClearDeletedAt() *DivisionUpdate {
	du.mutation.ClearDeletedAt()
	return du
}

// SetDeletedBy sets the "deleted_by" field.
func (du *DivisionUpdate) SetDeletedBy(s string) *DivisionUpdate {
	du.mutation.SetDeletedBy(s)
	return du
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableDeletedBy(s *string) *DivisionUpdate {
	if s != nil {
		du.SetDeletedBy(*s)
	}
	return du
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (du *DivisionUpdate) ClearDeletedBy() *DivisionUpdate {
	du.mutation.ClearDeletedBy()
	return du
}

// SetIsDeleted sets the "is_deleted" field.
func (du *DivisionUpdate) SetIsDeleted(b bool) *DivisionUpdate {
	du.mutation.SetIsDeleted(b)
	return du
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (du *DivisionUpdate) SetNillableIsDeleted(b *bool) *DivisionUpdate {
	if b != nil {
		du.SetIsDeleted(*b)
	}
	return du
}

// Mutation returns the DivisionMutation object of the builder.
func (du *DivisionUpdate) Mutation() *DivisionMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DivisionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DivisionUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DivisionUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DivisionUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DivisionUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := division.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Division.name": %w`, err)}
		}
	}
	if v, ok := du.mutation.Description(); ok {
		if err := division.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Division.description": %w`, err)}
		}
	}
	if v, ok := du.mutation.CreatedBy(); ok {
		if err := division.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Division.created_by": %w`, err)}
		}
	}
	return nil
}

func (du *DivisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(division.Table, division.Columns, sqlgraph.NewFieldSpec(division.FieldID, field.TypeString))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(division.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(division.FieldDescription, field.TypeString, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(division.FieldCreatedAt, field.TypeTime, value)
	}
	if du.mutation.CreatedAtCleared() {
		_spec.ClearField(division.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := du.mutation.CreatedBy(); ok {
		_spec.SetField(division.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(division.FieldUpdatedAt, field.TypeTime, value)
	}
	if du.mutation.UpdatedAtCleared() {
		_spec.ClearField(division.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := du.mutation.UpdatedBy(); ok {
		_spec.SetField(division.FieldUpdatedBy, field.TypeString, value)
	}
	if du.mutation.UpdatedByCleared() {
		_spec.ClearField(division.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.SetField(division.FieldDeletedAt, field.TypeTime, value)
	}
	if du.mutation.DeletedAtCleared() {
		_spec.ClearField(division.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := du.mutation.DeletedBy(); ok {
		_spec.SetField(division.FieldDeletedBy, field.TypeString, value)
	}
	if du.mutation.DeletedByCleared() {
		_spec.ClearField(division.FieldDeletedBy, field.TypeString)
	}
	if value, ok := du.mutation.IsDeleted(); ok {
		_spec.SetField(division.FieldIsDeleted, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{division.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DivisionUpdateOne is the builder for updating a single Division entity.
type DivisionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DivisionMutation
}

// SetName sets the "name" field.
func (duo *DivisionUpdateOne) SetName(s string) *DivisionUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetDescription sets the "description" field.
func (duo *DivisionUpdateOne) SetDescription(s string) *DivisionUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DivisionUpdateOne) SetCreatedAt(t time.Time) *DivisionUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableCreatedAt(t *time.Time) *DivisionUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (duo *DivisionUpdateOne) ClearCreatedAt() *DivisionUpdateOne {
	duo.mutation.ClearCreatedAt()
	return duo
}

// SetCreatedBy sets the "created_by" field.
func (duo *DivisionUpdateOne) SetCreatedBy(s string) *DivisionUpdateOne {
	duo.mutation.SetCreatedBy(s)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DivisionUpdateOne) SetUpdatedAt(t time.Time) *DivisionUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableUpdatedAt(t *time.Time) *DivisionUpdateOne {
	if t != nil {
		duo.SetUpdatedAt(*t)
	}
	return duo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (duo *DivisionUpdateOne) ClearUpdatedAt() *DivisionUpdateOne {
	duo.mutation.ClearUpdatedAt()
	return duo
}

// SetUpdatedBy sets the "updated_by" field.
func (duo *DivisionUpdateOne) SetUpdatedBy(s string) *DivisionUpdateOne {
	duo.mutation.SetUpdatedBy(s)
	return duo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableUpdatedBy(s *string) *DivisionUpdateOne {
	if s != nil {
		duo.SetUpdatedBy(*s)
	}
	return duo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (duo *DivisionUpdateOne) ClearUpdatedBy() *DivisionUpdateOne {
	duo.mutation.ClearUpdatedBy()
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DivisionUpdateOne) SetDeletedAt(t time.Time) *DivisionUpdateOne {
	duo.mutation.SetDeletedAt(t)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableDeletedAt(t *time.Time) *DivisionUpdateOne {
	if t != nil {
		duo.SetDeletedAt(*t)
	}
	return duo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (duo *DivisionUpdateOne) ClearDeletedAt() *DivisionUpdateOne {
	duo.mutation.ClearDeletedAt()
	return duo
}

// SetDeletedBy sets the "deleted_by" field.
func (duo *DivisionUpdateOne) SetDeletedBy(s string) *DivisionUpdateOne {
	duo.mutation.SetDeletedBy(s)
	return duo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableDeletedBy(s *string) *DivisionUpdateOne {
	if s != nil {
		duo.SetDeletedBy(*s)
	}
	return duo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (duo *DivisionUpdateOne) ClearDeletedBy() *DivisionUpdateOne {
	duo.mutation.ClearDeletedBy()
	return duo
}

// SetIsDeleted sets the "is_deleted" field.
func (duo *DivisionUpdateOne) SetIsDeleted(b bool) *DivisionUpdateOne {
	duo.mutation.SetIsDeleted(b)
	return duo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (duo *DivisionUpdateOne) SetNillableIsDeleted(b *bool) *DivisionUpdateOne {
	if b != nil {
		duo.SetIsDeleted(*b)
	}
	return duo
}

// Mutation returns the DivisionMutation object of the builder.
func (duo *DivisionUpdateOne) Mutation() *DivisionMutation {
	return duo.mutation
}

// Where appends a list predicates to the DivisionUpdate builder.
func (duo *DivisionUpdateOne) Where(ps ...predicate.Division) *DivisionUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DivisionUpdateOne) Select(field string, fields ...string) *DivisionUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Division entity.
func (duo *DivisionUpdateOne) Save(ctx context.Context) (*Division, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DivisionUpdateOne) SaveX(ctx context.Context) *Division {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DivisionUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DivisionUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DivisionUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := division.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Division.name": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Description(); ok {
		if err := division.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Division.description": %w`, err)}
		}
	}
	if v, ok := duo.mutation.CreatedBy(); ok {
		if err := division.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Division.created_by": %w`, err)}
		}
	}
	return nil
}

func (duo *DivisionUpdateOne) sqlSave(ctx context.Context) (_node *Division, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(division.Table, division.Columns, sqlgraph.NewFieldSpec(division.FieldID, field.TypeString))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Division.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, division.FieldID)
		for _, f := range fields {
			if !division.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != division.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(division.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(division.FieldDescription, field.TypeString, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(division.FieldCreatedAt, field.TypeTime, value)
	}
	if duo.mutation.CreatedAtCleared() {
		_spec.ClearField(division.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := duo.mutation.CreatedBy(); ok {
		_spec.SetField(division.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(division.FieldUpdatedAt, field.TypeTime, value)
	}
	if duo.mutation.UpdatedAtCleared() {
		_spec.ClearField(division.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := duo.mutation.UpdatedBy(); ok {
		_spec.SetField(division.FieldUpdatedBy, field.TypeString, value)
	}
	if duo.mutation.UpdatedByCleared() {
		_spec.ClearField(division.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.SetField(division.FieldDeletedAt, field.TypeTime, value)
	}
	if duo.mutation.DeletedAtCleared() {
		_spec.ClearField(division.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := duo.mutation.DeletedBy(); ok {
		_spec.SetField(division.FieldDeletedBy, field.TypeString, value)
	}
	if duo.mutation.DeletedByCleared() {
		_spec.ClearField(division.FieldDeletedBy, field.TypeString)
	}
	if value, ok := duo.mutation.IsDeleted(); ok {
		_spec.SetField(division.FieldIsDeleted, field.TypeBool, value)
	}
	_node = &Division{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{division.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
