package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type AcctType int

const (
	WeChat AcctType = iota
	QQ
	GMAIL
)

// SubAccount holds the schema definition for the SubAccount entity.
type SubAccount struct {
	ent.Schema
}

// Fields of the SubAccount.
func (SubAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("acct_type").GoType(AcctType(0)),
		field.String("sub_acct"),
	}
}

// Edges of the SubAccount.
func (SubAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("subAccounts").Unique(),
	}
}

func (SubAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		CommonFields{},
	}
}
