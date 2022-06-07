package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type WhoMixin struct {
	mixin.Schema
}

func (WhoMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("create_by").Immutable(),
		field.String("update_by"),
	}
}
