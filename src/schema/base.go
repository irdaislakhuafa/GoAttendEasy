package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func GetBaseSchema() []ent.Field {
	fields := []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.Time("created_at").Default(time.Now()),
		field.String("created_by").NotEmpty(),
		field.Time("updated_at").Optional(),
		field.String("updated_by"),
		field.Time("deleted_at").Optional(),
		field.String("deleted_by"),
		field.Bool("is_deleted").Default(false),
	}
	return fields
}
