// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/reminder"
)

// ReminderCreate is the builder for creating a Reminder entity.
type ReminderCreate struct {
	config
	mutation *ReminderMutation
	hooks    []Hook
}

// SetIn sets the "in" field.
func (rc *ReminderCreate) SetIn(t time.Time) *ReminderCreate {
	rc.mutation.SetIn(t)
	return rc
}

// SetOut sets the "out" field.
func (rc *ReminderCreate) SetOut(t time.Time) *ReminderCreate {
	rc.mutation.SetOut(t)
	return rc
}

// SetDay sets the "day" field.
func (rc *ReminderCreate) SetDay(i int) *ReminderCreate {
	rc.mutation.SetDay(i)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReminderCreate) SetCreatedAt(t time.Time) *ReminderCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReminderCreate) SetNillableCreatedAt(t *time.Time) *ReminderCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetCreatedBy sets the "created_by" field.
func (rc *ReminderCreate) SetCreatedBy(s string) *ReminderCreate {
	rc.mutation.SetCreatedBy(s)
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReminderCreate) SetUpdatedAt(t time.Time) *ReminderCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ReminderCreate) SetNillableUpdatedAt(t *time.Time) *ReminderCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetUpdatedBy sets the "updated_by" field.
func (rc *ReminderCreate) SetUpdatedBy(s string) *ReminderCreate {
	rc.mutation.SetUpdatedBy(s)
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *ReminderCreate) SetDeletedAt(t time.Time) *ReminderCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *ReminderCreate) SetNillableDeletedAt(t *time.Time) *ReminderCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetDeletedBy sets the "deleted_by" field.
func (rc *ReminderCreate) SetDeletedBy(s string) *ReminderCreate {
	rc.mutation.SetDeletedBy(s)
	return rc
}

// SetIsDeleted sets the "is_deleted" field.
func (rc *ReminderCreate) SetIsDeleted(b bool) *ReminderCreate {
	rc.mutation.SetIsDeleted(b)
	return rc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (rc *ReminderCreate) SetNillableIsDeleted(b *bool) *ReminderCreate {
	if b != nil {
		rc.SetIsDeleted(*b)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReminderCreate) SetID(u uuid.UUID) *ReminderCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ReminderCreate) SetNillableID(u *uuid.UUID) *ReminderCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// Mutation returns the ReminderMutation object of the builder.
func (rc *ReminderCreate) Mutation() *ReminderMutation {
	return rc.mutation
}

// Save creates the Reminder in the database.
func (rc *ReminderCreate) Save(ctx context.Context) (*Reminder, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReminderCreate) SaveX(ctx context.Context) *Reminder {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReminderCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReminderCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReminderCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := reminder.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.IsDeleted(); !ok {
		v := reminder.DefaultIsDeleted
		rc.mutation.SetIsDeleted(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := reminder.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReminderCreate) check() error {
	if _, ok := rc.mutation.In(); !ok {
		return &ValidationError{Name: "in", err: errors.New(`generated: missing required field "Reminder.in"`)}
	}
	if _, ok := rc.mutation.Out(); !ok {
		return &ValidationError{Name: "out", err: errors.New(`generated: missing required field "Reminder.out"`)}
	}
	if _, ok := rc.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`generated: missing required field "Reminder.day"`)}
	}
	if v, ok := rc.mutation.Day(); ok {
		if err := reminder.DayValidator(v); err != nil {
			return &ValidationError{Name: "day", err: fmt.Errorf(`generated: validator failed for field "Reminder.day": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Reminder.created_at"`)}
	}
	if _, ok := rc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`generated: missing required field "Reminder.created_by"`)}
	}
	if v, ok := rc.mutation.CreatedBy(); ok {
		if err := reminder.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`generated: validator failed for field "Reminder.created_by": %w`, err)}
		}
	}
	if _, ok := rc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`generated: missing required field "Reminder.updated_by"`)}
	}
	if _, ok := rc.mutation.DeletedBy(); !ok {
		return &ValidationError{Name: "deleted_by", err: errors.New(`generated: missing required field "Reminder.deleted_by"`)}
	}
	if _, ok := rc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`generated: missing required field "Reminder.is_deleted"`)}
	}
	return nil
}

func (rc *ReminderCreate) sqlSave(ctx context.Context) (*Reminder, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReminderCreate) createSpec() (*Reminder, *sqlgraph.CreateSpec) {
	var (
		_node = &Reminder{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(reminder.Table, sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.In(); ok {
		_spec.SetField(reminder.FieldIn, field.TypeTime, value)
		_node.In = value
	}
	if value, ok := rc.mutation.Out(); ok {
		_spec.SetField(reminder.FieldOut, field.TypeTime, value)
		_node.Out = value
	}
	if value, ok := rc.mutation.Day(); ok {
		_spec.SetField(reminder.FieldDay, field.TypeInt, value)
		_node.Day = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(reminder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.CreatedBy(); ok {
		_spec.SetField(reminder.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(reminder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.UpdatedBy(); ok {
		_spec.SetField(reminder.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(reminder.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := rc.mutation.DeletedBy(); ok {
		_spec.SetField(reminder.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := rc.mutation.IsDeleted(); ok {
		_spec.SetField(reminder.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	return _node, _spec
}

// ReminderCreateBulk is the builder for creating many Reminder entities in bulk.
type ReminderCreateBulk struct {
	config
	err      error
	builders []*ReminderCreate
}

// Save creates the Reminder entities in the database.
func (rcb *ReminderCreateBulk) Save(ctx context.Context) ([]*Reminder, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reminder, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReminderMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReminderCreateBulk) SaveX(ctx context.Context) []*Reminder {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReminderCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReminderCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
