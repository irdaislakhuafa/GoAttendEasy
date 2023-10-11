package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("user_id").NotEmpty(),
		field.String("role_id").NotEmpty(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}
