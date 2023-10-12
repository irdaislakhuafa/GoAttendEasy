// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/division"
)

// Division is the model entity for the Division schema.
type Division struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty" pk`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted    bool `json:"is_deleted,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Division) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case division.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case division.FieldID, division.FieldName, division.FieldDescription, division.FieldCreatedBy, division.FieldUpdatedBy, division.FieldDeletedBy:
			values[i] = new(sql.NullString)
		case division.FieldCreatedAt, division.FieldUpdatedAt, division.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Division fields.
func (d *Division) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case division.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case division.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case division.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case division.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case division.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				d.CreatedBy = value.String
			}
		case division.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case division.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				d.UpdatedBy = value.String
			}
		case division.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = value.Time
			}
		case division.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				d.DeletedBy = value.String
			}
		case division.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				d.IsDeleted = value.Bool
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Division.
// This includes values selected through modifiers, order, etc.
func (d *Division) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this Division.
// Note that you need to call Division.Unwrap() before calling this method if this Division
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Division) Update() *DivisionUpdateOne {
	return NewDivisionClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Division entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Division) Unwrap() *Division {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("generated: Division is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Division) String() string {
	var builder strings.Builder
	builder.WriteString("Division(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(d.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(d.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(d.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(d.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", d.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// Divisions is a parsable slice of Division.
type Divisions []*Division