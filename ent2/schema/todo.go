package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Bool("done").
			Optional(),
		field.Bool("clean").Default(false),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
