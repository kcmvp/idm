package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// RoleFunc holds the schema definition for the RoleFunc entity.
type RoleFunc struct {
	ent.Schema
}

// Fields of the RoleFunc.
func (RoleFunc) Fields() []ent.Field {
	return []ent.Field{
		field.String("url_pattern").NotEmpty(),
	}
}

// Edges of the RoleFunc.
func (RoleFunc) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("funcs"),
	}
}

func (RoleFunc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		CommonFields{},
	}
}
