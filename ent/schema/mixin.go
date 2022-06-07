package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CommonFields struct {
	mixin.Schema
}

func (CommonFields) Fields() []ent.Field {
	return []ent.Field{
		field.String("create_by").Immutable(),
		field.String("update_by"),
		field.Bool("deleted").Default(false),
	}
}
