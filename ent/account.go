// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kcmvp/idm.go/ent/account"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// CreateBy holds the value of the "create_by" field.
	CreateBy string `json:"create_by,omitempty"`
	// UpdateBy holds the value of the "update_by" field.
	UpdateBy string `json:"update_by,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	// Mobile holds the value of the "mobile" field.
	Mobile string `json:"mobile,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Passwd holds the value of the "passwd" field.
	Passwd string `json:"passwd,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges AccountEdges `json:"edges"`
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// SubAccounts holds the value of the subAccounts edge.
	SubAccounts []*SubAccount `json:"subAccounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubAccountsOrErr returns the SubAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) SubAccountsOrErr() ([]*SubAccount, error) {
	if e.loadedTypes[0] {
		return e.SubAccounts, nil
	}
	return nil, &NotLoadedError{edge: "subAccounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldDeleted, account.FieldEmailConfirmed, account.FieldDisabled:
			values[i] = new(sql.NullBool)
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.FieldCreateBy, account.FieldUpdateBy, account.FieldEmail, account.FieldMobile, account.FieldFirstName, account.FieldLastName, account.FieldPasswd, account.FieldSource:
			values[i] = new(sql.NullString)
		case account.FieldCreateTime, account.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case account.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case account.FieldCreateBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				a.CreateBy = value.String
			}
		case account.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				a.UpdateBy = value.String
			}
		case account.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				a.Deleted = value.Bool
			}
		case account.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case account.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				a.EmailConfirmed = value.Bool
			}
		case account.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				a.Mobile = value.String
			}
		case account.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				a.FirstName = value.String
			}
		case account.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				a.LastName = value.String
			}
		case account.FieldPasswd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passwd", values[i])
			} else if value.Valid {
				a.Passwd = value.String
			}
		case account.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				a.Disabled = value.Bool
			}
		case account.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				a.Source = value.String
			}
		}
	}
	return nil
}

// QuerySubAccounts queries the "subAccounts" edge of the Account entity.
func (a *Account) QuerySubAccounts() *SubAccountQuery {
	return (&AccountClient{config: a.config}).QuerySubAccounts(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(a.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", create_by=")
	builder.WriteString(a.CreateBy)
	builder.WriteString(", update_by=")
	builder.WriteString(a.UpdateBy)
	builder.WriteString(", deleted=")
	builder.WriteString(fmt.Sprintf("%v", a.Deleted))
	builder.WriteString(", email=")
	builder.WriteString(a.Email)
	builder.WriteString(", email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", a.EmailConfirmed))
	builder.WriteString(", mobile=")
	builder.WriteString(a.Mobile)
	builder.WriteString(", first_name=")
	builder.WriteString(a.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(a.LastName)
	builder.WriteString(", passwd=")
	builder.WriteString(a.Passwd)
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", a.Disabled))
	builder.WriteString(", source=")
	builder.WriteString(a.Source)
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
