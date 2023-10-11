package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").MaxLen(1000).Optional(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
