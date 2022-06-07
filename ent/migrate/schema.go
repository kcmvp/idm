// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeString},
		{Name: "update_by", Type: field.TypeString},
		{Name: "app_id", Type: field.TypeString},
		{Name: "mobile", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "passwd", Type: field.TypeString},
		{Name: "disabled", Type: field.TypeBool},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// OauthsColumns holds the columns for the "oauths" table.
	OauthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeString},
		{Name: "update_by", Type: field.TypeString},
		{Name: "oauth_type", Type: field.TypeInt},
		{Name: "open_id", Type: field.TypeString},
		{Name: "account_oauth", Type: field.TypeInt, Nullable: true},
	}
	// OauthsTable holds the schema information for the "oauths" table.
	OauthsTable = &schema.Table{
		Name:       "oauths",
		Columns:    OauthsColumns,
		PrimaryKey: []*schema.Column{OauthsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "oauths_accounts_oauth",
				Columns:    []*schema.Column{OauthsColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		OauthsTable,
	}
)

func init() {
	OauthsTable.ForeignKeys[0].RefTable = AccountsTable
}