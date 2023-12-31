package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("role_id").NotEmpty(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
