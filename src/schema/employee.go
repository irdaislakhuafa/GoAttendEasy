package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("user_id").NotEmpty(),
		field.String("division_id").NotEmpty(),
	}
	fields = append(fields, GetBaseSchema()...)
	return fields
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
