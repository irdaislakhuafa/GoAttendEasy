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
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/employee"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUserID sets the "user_id" field.
func (eu *EmployeeUpdate) SetUserID(s string) *EmployeeUpdate {
	eu.mutation.SetUserID(s)
	return eu
}

// SetDivisionID sets the "division_id" field.
func (eu *EmployeeUpdate) SetDivisionID(s string) *EmployeeUpdate {
	eu.mutation.SetDivisionID(s)
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EmployeeUpdate) SetCreatedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableCreatedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (eu *EmployeeUpdate) ClearCreatedAt() *EmployeeUpdate {
	eu.mutation.ClearCreatedAt()
	return eu
}

// SetCreatedBy sets the "created_by" field.
func (eu *EmployeeUpdate) SetCreatedBy(s string) *EmployeeUpdate {
	eu.mutation.SetCreatedBy(s)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EmployeeUpdate) SetUpdatedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableUpdatedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetUpdatedAt(*t)
	}
	return eu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (eu *EmployeeUpdate) ClearUpdatedAt() *EmployeeUpdate {
	eu.mutation.ClearUpdatedAt()
	return eu
}

// SetUpdatedBy sets the "updated_by" field.
func (eu *EmployeeUpdate) SetUpdatedBy(s string) *EmployeeUpdate {
	eu.mutation.SetUpdatedBy(s)
	return eu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableUpdatedBy(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetUpdatedBy(*s)
	}
	return eu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (eu *EmployeeUpdate) ClearUpdatedBy() *EmployeeUpdate {
	eu.mutation.ClearUpdatedBy()
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EmployeeUpdate) SetDeletedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDeletedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (eu *EmployeeUpdate) ClearDeletedAt() *EmployeeUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// SetDeletedBy sets the "deleted_by" field.
func (eu *EmployeeUpdate) SetDeletedBy(s string) *EmployeeUpdate {
	eu.mutation.SetDeletedBy(s)
	return eu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDeletedBy(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetDeletedBy(*s)
	}
	return eu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (eu *EmployeeUpdate) ClearDeletedBy() *EmployeeUpdate {
	eu.mutation.ClearDeletedBy()
	return eu
}

// SetIsDeleted sets the "is_deleted" field.
func (eu *EmployeeUpdate) SetIsDeleted(b bool) *EmployeeUpdate {
	eu.mutation.SetIsDeleted(b)
	return eu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableIsDeleted(b *bool) *EmployeeUpdate {
	if b != nil {
		eu.SetIsDeleted(*b)
	}
	return eu
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmployeeUpdate) check() error {
	if v, ok := eu.mutation.UserID(); ok {
		if err := employee.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`generated: validator failed for field "Employee.user_id": %w`, err)}
		}
	}
	if v, ok := eu.mutation.DivisionID(); ok {
		if err := employee.DivisionIDValidator(v); err != nil {
			return &ValidationError{Name: "division_id", err: fmt.Errorf(`generated: validator failed for field "Employee.division_id": %w`, err)}
		}
	}
	if v, ok := eu.mutation.CreatedBy(); ok {
		if err := employee.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Employee.created_by": %w`, err)}
		}
	}
	return nil
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UserID(); ok {
		_spec.SetField(employee.FieldUserID, field.TypeString, value)
	}
	if value, ok := eu.mutation.DivisionID(); ok {
		_spec.SetField(employee.FieldDivisionID, field.TypeString, value)
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.SetField(employee.FieldCreatedAt, field.TypeTime, value)
	}
	if eu.mutation.CreatedAtCleared() {
		_spec.ClearField(employee.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.CreatedBy(); ok {
		_spec.SetField(employee.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if eu.mutation.UpdatedAtCleared() {
		_spec.ClearField(employee.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.UpdatedBy(); ok {
		_spec.SetField(employee.FieldUpdatedBy, field.TypeString, value)
	}
	if eu.mutation.UpdatedByCleared() {
		_spec.ClearField(employee.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.DeletedBy(); ok {
		_spec.SetField(employee.FieldDeletedBy, field.TypeString, value)
	}
	if eu.mutation.DeletedByCleared() {
		_spec.ClearField(employee.FieldDeletedBy, field.TypeString)
	}
	if value, ok := eu.mutation.IsDeleted(); ok {
		_spec.SetField(employee.FieldIsDeleted, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetUserID sets the "user_id" field.
func (euo *EmployeeUpdateOne) SetUserID(s string) *EmployeeUpdateOne {
	euo.mutation.SetUserID(s)
	return euo
}

// SetDivisionID sets the "division_id" field.
func (euo *EmployeeUpdateOne) SetDivisionID(s string) *EmployeeUpdateOne {
	euo.mutation.SetDivisionID(s)
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EmployeeUpdateOne) SetCreatedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableCreatedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (euo *EmployeeUpdateOne) ClearCreatedAt() *EmployeeUpdateOne {
	euo.mutation.ClearCreatedAt()
	return euo
}

// SetCreatedBy sets the "created_by" field.
func (euo *EmployeeUpdateOne) SetCreatedBy(s string) *EmployeeUpdateOne {
	euo.mutation.SetCreatedBy(s)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EmployeeUpdateOne) SetUpdatedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableUpdatedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetUpdatedAt(*t)
	}
	return euo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (euo *EmployeeUpdateOne) ClearUpdatedAt() *EmployeeUpdateOne {
	euo.mutation.ClearUpdatedAt()
	return euo
}

// SetUpdatedBy sets the "updated_by" field.
func (euo *EmployeeUpdateOne) SetUpdatedBy(s string) *EmployeeUpdateOne {
	euo.mutation.SetUpdatedBy(s)
	return euo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableUpdatedBy(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetUpdatedBy(*s)
	}
	return euo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (euo *EmployeeUpdateOne) ClearUpdatedBy() *EmployeeUpdateOne {
	euo.mutation.ClearUpdatedBy()
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EmployeeUpdateOne) SetDeletedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDeletedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (euo *EmployeeUpdateOne) ClearDeletedAt() *EmployeeUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// SetDeletedBy sets the "deleted_by" field.
func (euo *EmployeeUpdateOne) SetDeletedBy(s string) *EmployeeUpdateOne {
	euo.mutation.SetDeletedBy(s)
	return euo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDeletedBy(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetDeletedBy(*s)
	}
	return euo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (euo *EmployeeUpdateOne) ClearDeletedBy() *EmployeeUpdateOne {
	euo.mutation.ClearDeletedBy()
	return euo
}

// SetIsDeleted sets the "is_deleted" field.
func (euo *EmployeeUpdateOne) SetIsDeleted(b bool) *EmployeeUpdateOne {
	euo.mutation.SetIsDeleted(b)
	return euo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableIsDeleted(b *bool) *EmployeeUpdateOne {
	if b != nil {
		euo.SetIsDeleted(*b)
	}
	return euo
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmployeeUpdateOne) check() error {
	if v, ok := euo.mutation.UserID(); ok {
		if err := employee.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`generated: validator failed for field "Employee.user_id": %w`, err)}
		}
	}
	if v, ok := euo.mutation.DivisionID(); ok {
		if err := employee.DivisionIDValidator(v); err != nil {
			return &ValidationError{Name: "division_id", err: fmt.Errorf(`generated: validator failed for field "Employee.division_id": %w`, err)}
		}
	}
	if v, ok := euo.mutation.CreatedBy(); ok {
		if err := employee.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Employee.created_by": %w`, err)}
		}
	}
	return nil
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UserID(); ok {
		_spec.SetField(employee.FieldUserID, field.TypeString, value)
	}
	if value, ok := euo.mutation.DivisionID(); ok {
		_spec.SetField(employee.FieldDivisionID, field.TypeString, value)
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.SetField(employee.FieldCreatedAt, field.TypeTime, value)
	}
	if euo.mutation.CreatedAtCleared() {
		_spec.ClearField(employee.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.CreatedBy(); ok {
		_spec.SetField(employee.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if euo.mutation.UpdatedAtCleared() {
		_spec.ClearField(employee.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.UpdatedBy(); ok {
		_spec.SetField(employee.FieldUpdatedBy, field.TypeString, value)
	}
	if euo.mutation.UpdatedByCleared() {
		_spec.ClearField(employee.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.DeletedBy(); ok {
		_spec.SetField(employee.FieldDeletedBy, field.TypeString, value)
	}
	if euo.mutation.DeletedByCleared() {
		_spec.ClearField(employee.FieldDeletedBy, field.TypeString)
	}
	if value, ok := euo.mutation.IsDeleted(); ok {
		_spec.SetField(employee.FieldIsDeleted, field.TypeBool, value)
	}
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
