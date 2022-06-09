// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kcmvp/idm/ent/fun"
)

// Fun is the model entity for the Fun schema.
type Fun struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FunQuery when eager-loading is set.
	Edges FunEdges `json:"edges"`
}

// FunEdges holds the relations/edges for other nodes in the graph.
type FunEdges struct {
	// Role holds the value of the role edge.
	Role []*Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e FunEdges) RoleOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fun) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fun.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Fun", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fun fields.
func (f *Fun) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fun.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		}
	}
	return nil
}

// QueryRole queries the "role" edge of the Fun entity.
func (f *Fun) QueryRole() *RoleQuery {
	return (&FunClient{config: f.config}).QueryRole(f)
}

// Update returns a builder for updating this Fun.
// Note that you need to call Fun.Unwrap() before calling this method if this Fun
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fun) Update() *FunUpdateOne {
	return (&FunClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Fun entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fun) Unwrap() *Fun {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fun is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fun) String() string {
	var builder strings.Builder
	builder.WriteString("Fun(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Funs is a parsable slice of Fun.
type Funs []*Fun

func (f Funs) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}