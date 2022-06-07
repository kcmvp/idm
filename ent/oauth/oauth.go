// Code generated by entc, DO NOT EDIT.

package oauth

import (
	"time"
)

const (
	// Label holds the string label denoting the oauth type in the database.
	Label = "oauth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldUpdateBy holds the string denoting the update_by field in the database.
	FieldUpdateBy = "update_by"
	// FieldOauthType holds the string denoting the oauth_type field in the database.
	FieldOauthType = "oauth_type"
	// FieldOpenID holds the string denoting the open_id field in the database.
	FieldOpenID = "open_id"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the oauth in the database.
	Table = "oauths"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "oauths"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_oauth"
)

// Columns holds all SQL columns for oauth fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCreateBy,
	FieldUpdateBy,
	FieldOauthType,
	FieldOpenID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "oauths"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_oauth",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
