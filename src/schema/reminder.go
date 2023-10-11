package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Reminder holds the schema definition for the Reminder entity.
type Reminder struct {
	ent.Schema
}

// Fields of the Reminder.
func (Reminder) Fields() []ent.Field {
	fields := []ent.Field{
		field.Time("in"),
		field.Time("out"),
		field.Int("day").Positive(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the Reminder.
func (Reminder) Edges() []ent.Edge {
	return nil
}
