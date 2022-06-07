// Code generated by entc, DO NOT EDIT.

package role

import (
	"time"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeApplication holds the string denoting the application edge name in mutations.
	EdgeApplication = "application"
	// EdgeFuncs holds the string denoting the funcs edge name in mutations.
	EdgeFuncs = "funcs"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// ApplicationTable is the table that holds the application relation/edge. The primary key declared below.
	ApplicationTable = "application_roles"
	// ApplicationInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationInverseTable = "applications"
	// FuncsTable is the table that holds the funcs relation/edge. The primary key declared below.
	FuncsTable = "role_funcs"
	// FuncsInverseTable is the table name for the RoleFunc entity.
	// It exists in this package in order to avoid circular dependency with the "rolefunc" package.
	FuncsInverseTable = "role_funcs"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCreateBy,
	FieldUpdateBy,
	FieldDeleted,
	FieldName,
	FieldDesc,
}

var (
	// ApplicationPrimaryKey and ApplicationColumn2 are the table columns denoting the
	// primary key for the application relation (M2M).
	ApplicationPrimaryKey = []string{"application_id", "role_id"}
	// FuncsPrimaryKey and FuncsColumn2 are the table columns denoting the
	// primary key for the funcs relation (M2M).
	FuncsPrimaryKey = []string{"role_id", "role_func_id"}
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