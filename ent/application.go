// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kcmvp/iam.go/ent/application"
)

// Application is the model entity for the Application schema.
type Application struct {
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
	// AppName holds the value of the "app_name" field.
	AppName string `json:"app_name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationQuery when eager-loading is set.
	Edges ApplicationEdges `json:"edges"`
}

// ApplicationEdges holds the relations/edges for other nodes in the graph.
type ApplicationEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldDeleted:
			values[i] = new(sql.NullBool)
		case application.FieldID:
			values[i] = new(sql.NullInt64)
		case application.FieldCreateBy, application.FieldUpdateBy, application.FieldAppName, application.FieldURL:
			values[i] = new(sql.NullString)
		case application.FieldCreateTime, application.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Application", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case application.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case application.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case application.FieldCreateBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				a.CreateBy = value.String
			}
		case application.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				a.UpdateBy = value.String
			}
		case application.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				a.Deleted = value.Bool
			}
		case application.FieldAppName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_name", values[i])
			} else if value.Valid {
				a.AppName = value.String
			}
		case application.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		}
	}
	return nil
}

// QueryRoles queries the "roles" edge of the Application entity.
func (a *Application) QueryRoles() *RoleQuery {
	return (&ApplicationClient{config: a.config}).QueryRoles(a)
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return (&ApplicationClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
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
	builder.WriteString(", app_name=")
	builder.WriteString(a.AppName)
	builder.WriteString(", url=")
	builder.WriteString(a.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application

func (a Applications) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
