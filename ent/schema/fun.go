package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Fun holds the schema definition for the Fun entity.
type Fun struct {
	ent.Schema
}

// Fields of the Fun.
func (Fun) Fields() []ent.Field {
	return nil
}

// Edges of the Fun.
func (Fun) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("funcs"),
	}
}
