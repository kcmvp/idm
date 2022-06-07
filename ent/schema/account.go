package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.Bool("email_confirmed").Default(false),
		field.String("mobile"),
		field.String("first_name"),
		field.String("last_name"),
		field.String("passwd"),
		field.Bool("disabled").Default(false),
		field.String("source"),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		CommonFields{},
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subAccounts", SubAccount.Type),
	}
}
