package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func GetBaseSchema() []ent.Field {
	fields := []ent.Field{
		field.String("id").StructTag("pk").Default(uuid.NewString()),
		field.Time("created_at").Optional().Default(time.Now()),
		field.String("created_by").NotEmpty(),
		field.Time("updated_at").Optional(),
		field.String("updated_by").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("deleted_by").Optional(),
		field.Bool("is_deleted").Default(false),
	}
	return fields
}
