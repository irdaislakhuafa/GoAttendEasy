package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Division holds the schema definition for the Division entity.
type Division struct {
	ent.Schema
}

// Fields of the Division.
func (Division) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").NotEmpty(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the Division.
func (Division) Edges() []ent.Edge {
	return nil
}
