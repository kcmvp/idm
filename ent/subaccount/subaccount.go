// Code generated by entc, DO NOT EDIT.

package subaccount

import (
	"time"
)

const (
	// Label holds the string label denoting the subaccount type in the database.
	Label = "sub_account"
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
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldAcctType holds the string denoting the acct_type field in the database.
	FieldAcctType = "acct_type"
	// FieldSubAcct holds the string denoting the sub_acct field in the database.
	FieldSubAcct = "sub_acct"
	// EdgeAccoun holds the string denoting the accoun edge name in mutations.
	EdgeAccoun = "accoun"
	// Table holds the table name of the subaccount in the database.
	Table = "sub_accounts"
	// AccounTable is the table that holds the accoun relation/edge. The primary key declared below.
	AccounTable = "account_subAccounts"
	// AccounInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccounInverseTable = "accounts"
)

// Columns holds all SQL columns for subaccount fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCreateBy,
	FieldUpdateBy,
	FieldDeleted,
	FieldAcctType,
	FieldSubAcct,
}

var (
	// AccounPrimaryKey and AccounColumn2 are the table columns denoting the
	// primary key for the accoun relation (M2M).
	AccounPrimaryKey = []string{"account_id", "sub_account_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted bool
)