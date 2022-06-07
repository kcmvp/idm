package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// OAuth holds the schema definition for the OAuth entity.
type OAuth struct {
	ent.Schema
}

func (OAuth) Mixin() []ent.Mixin {

	return []ent.Mixin{
		mixin.Time{},
		WhoMixin{},
	}
}

// Fields of the OAuth.
func (OAuth) Fields() []ent.Field {
	return []ent.Field{
		field.Int("oauth_type"),
		field.String("open_id"),
	}
}

// Edges of the OAuth.
func (OAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("oauth").Unique(),
	}
}
