// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/attendance"
)

// AttendanceCreate is the builder for creating a Attendance entity.
type AttendanceCreate struct {
	config
	mutation *AttendanceMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ac *AttendanceCreate) SetUserID(s string) *AttendanceCreate {
	ac.mutation.SetUserID(s)
	return ac
}

// SetIn sets the "in" field.
func (ac *AttendanceCreate) SetIn(t time.Time) *AttendanceCreate {
	ac.mutation.SetIn(t)
	return ac
}

// SetOut sets the "out" field.
func (ac *AttendanceCreate) SetOut(t time.Time) *AttendanceCreate {
	ac.mutation.SetOut(t)
	return ac
}

// SetNillableOut sets the "out" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableOut(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetOut(*t)
	}
	return ac
}

// SetIsPresent sets the "is_present" field.
func (ac *AttendanceCreate) SetIsPresent(b bool) *AttendanceCreate {
	ac.mutation.SetIsPresent(b)
	return ac
}

// SetNillableIsPresent sets the "is_present" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableIsPresent(b *bool) *AttendanceCreate {
	if b != nil {
		ac.SetIsPresent(*b)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttendanceCreate) SetCreatedAt(t time.Time) *AttendanceCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableCreatedAt(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *AttendanceCreate) SetCreatedBy(s string) *AttendanceCreate {
	ac.mutation.SetCreatedBy(s)
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttendanceCreate) SetUpdatedAt(t time.Time) *AttendanceCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableUpdatedAt(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *AttendanceCreate) SetUpdatedBy(s string) *AttendanceCreate {
	ac.mutation.SetUpdatedBy(s)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableUpdatedBy(s *string) *AttendanceCreate {
	if s != nil {
		ac.SetUpdatedBy(*s)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AttendanceCreate) SetDeletedAt(t time.Time) *AttendanceCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableDeletedAt(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetDeletedBy sets the "deleted_by" field.
func (ac *AttendanceCreate) SetDeletedBy(s string) *AttendanceCreate {
	ac.mutation.SetDeletedBy(s)
	return ac
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableDeletedBy(s *string) *AttendanceCreate {
	if s != nil {
		ac.SetDeletedBy(*s)
	}
	return ac
}

// SetIsDeleted sets the "is_deleted" field.
func (ac *AttendanceCreate) SetIsDeleted(b bool) *AttendanceCreate {
	ac.mutation.SetIsDeleted(b)
	return ac
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableIsDeleted(b *bool) *AttendanceCreate {
	if b != nil {
		ac.SetIsDeleted(*b)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AttendanceCreate) SetID(s string) *AttendanceCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableID(s *string) *AttendanceCreate {
	if s != nil {
		ac.SetID(*s)
	}
	return ac
}

// Mutation returns the AttendanceMutation object of the builder.
func (ac *AttendanceCreate) Mutation() *AttendanceMutation {
	return ac.mutation
}

// Save creates the Attendance in the database.
func (ac *AttendanceCreate) Save(ctx context.Context) (*Attendance, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttendanceCreate) SaveX(ctx context.Context) *Attendance {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttendanceCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttendanceCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttendanceCreate) defaults() {
	if _, ok := ac.mutation.IsPresent(); !ok {
		v := attendance.DefaultIsPresent
		ac.mutation.SetIsPresent(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attendance.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.IsDeleted(); !ok {
		v := attendance.DefaultIsDeleted
		ac.mutation.SetIsDeleted(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := attendance.DefaultID
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttendanceCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "Attendance.user_id"`)}
	}
	if v, ok := ac.mutation.UserID(); ok {
		if err := attendance.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`generated: validator failed for field "Attendance.user_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.In(); !ok {
		return &ValidationError{Name: "in", err: errors.New(`generated: missing required field "Attendance.in"`)}
	}
	if _, ok := ac.mutation.IsPresent(); !ok {
		return &ValidationError{Name: "is_present", err: errors.New(`generated: missing required field "Attendance.is_present"`)}
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`generated: missing required field "Attendance.created_by"`)}
	}
	if v, ok := ac.mutation.CreatedBy(); ok {
		if err := attendance.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Attendance.created_by": %w`, err)}
		}
	}
	if _, ok := ac.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`generated: missing required field "Attendance.is_deleted"`)}
	}
	return nil
}

func (ac *AttendanceCreate) sqlSave(ctx context.Context) (*Attendance, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Attendance.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttendanceCreate) createSpec() (*Attendance, *sqlgraph.CreateSpec) {
	var (
		_node = &Attendance{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attendance.Table, sqlgraph.NewFieldSpec(attendance.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.UserID(); ok {
		_spec.SetField(attendance.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := ac.mutation.In(); ok {
		_spec.SetField(attendance.FieldIn, field.TypeTime, value)
		_node.In = value
	}
	if value, ok := ac.mutation.Out(); ok {
		_spec.SetField(attendance.FieldOut, field.TypeTime, value)
		_node.Out = value
	}
	if value, ok := ac.mutation.IsPresent(); ok {
		_spec.SetField(attendance.FieldIsPresent, field.TypeBool, value)
		_node.IsPresent = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attendance.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(attendance.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(attendance.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(attendance.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(attendance.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.DeletedBy(); ok {
		_spec.SetField(attendance.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ac.mutation.IsDeleted(); ok {
		_spec.SetField(attendance.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	return _node, _spec
}

// AttendanceCreateBulk is the builder for creating many Attendance entities in bulk.
type AttendanceCreateBulk struct {
	config
	err      error
	builders []*AttendanceCreate
}

// Save creates the Attendance entities in the database.
func (acb *AttendanceCreateBulk) Save(ctx context.Context) ([]*Attendance, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attendance, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttendanceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttendanceCreateBulk) SaveX(ctx context.Context) []*Attendance {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttendanceCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttendanceCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
