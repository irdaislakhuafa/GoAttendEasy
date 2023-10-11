package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("user_id").NotEmpty(),
		field.Time("in"),
		field.Time("out"),
		field.Bool("is_present").Default(false),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return nil
}
