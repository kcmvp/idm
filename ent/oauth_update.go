// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcmvp/iam.go/ent/account"
	"github.com/kcmvp/iam.go/ent/oauth"
	"github.com/kcmvp/iam.go/ent/predicate"
)

// OAuthUpdate is the builder for updating OAuth entities.
type OAuthUpdate struct {
	config
	hooks    []Hook
	mutation *OAuthMutation
}

// Where appends a list predicates to the OAuthUpdate builder.
func (ou *OAuthUpdate) Where(ps ...predicate.OAuth) *OAuthUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdateTime sets the "update_time" field.
func (ou *OAuthUpdate) SetUpdateTime(t time.Time) *OAuthUpdate {
	ou.mutation.SetUpdateTime(t)
	return ou
}

// SetUpdateBy sets the "update_by" field.
func (ou *OAuthUpdate) SetUpdateBy(s string) *OAuthUpdate {
	ou.mutation.SetUpdateBy(s)
	return ou
}

// SetOauthType sets the "oauth_type" field.
func (ou *OAuthUpdate) SetOauthType(i int) *OAuthUpdate {
	ou.mutation.ResetOauthType()
	ou.mutation.SetOauthType(i)
	return ou
}

// AddOauthType adds i to the "oauth_type" field.
func (ou *OAuthUpdate) AddOauthType(i int) *OAuthUpdate {
	ou.mutation.AddOauthType(i)
	return ou
}

// SetOpenID sets the "open_id" field.
func (ou *OAuthUpdate) SetOpenID(s string) *OAuthUpdate {
	ou.mutation.SetOpenID(s)
	return ou
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (ou *OAuthUpdate) SetAccountID(id int) *OAuthUpdate {
	ou.mutation.SetAccountID(id)
	return ou
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (ou *OAuthUpdate) SetNillableAccountID(id *int) *OAuthUpdate {
	if id != nil {
		ou = ou.SetAccountID(*id)
	}
	return ou
}

// SetAccount sets the "account" edge to the Account entity.
func (ou *OAuthUpdate) SetAccount(a *Account) *OAuthUpdate {
	return ou.SetAccountID(a.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (ou *OAuthUpdate) Mutation() *OAuthMutation {
	return ou.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (ou *OAuthUpdate) ClearAccount() *OAuthUpdate {
	ou.mutation.ClearAccount()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OAuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OAuthUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OAuthUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OAuthUpdate) defaults() {
	if _, ok := ou.mutation.UpdateTime(); !ok {
		v := oauth.UpdateDefaultUpdateTime()
		ou.mutation.SetUpdateTime(v)
	}
}

func (ou *OAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth.Table,
			Columns: oauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth.FieldUpdateTime,
		})
	}
	if value, ok := ou.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldUpdateBy,
		})
	}
	if value, ok := ou.mutation.OauthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauth.FieldOauthType,
		})
	}
	if value, ok := ou.mutation.AddedOauthType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauth.FieldOauthType,
		})
	}
	if value, ok := ou.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldOpenID,
		})
	}
	if ou.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauth.AccountTable,
			Columns: []string{oauth.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauth.AccountTable,
			Columns: []string{oauth.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OAuthUpdateOne is the builder for updating a single OAuth entity.
type OAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OAuthMutation
}

// SetUpdateTime sets the "update_time" field.
func (ouo *OAuthUpdateOne) SetUpdateTime(t time.Time) *OAuthUpdateOne {
	ouo.mutation.SetUpdateTime(t)
	return ouo
}

// SetUpdateBy sets the "update_by" field.
func (ouo *OAuthUpdateOne) SetUpdateBy(s string) *OAuthUpdateOne {
	ouo.mutation.SetUpdateBy(s)
	return ouo
}

// SetOauthType sets the "oauth_type" field.
func (ouo *OAuthUpdateOne) SetOauthType(i int) *OAuthUpdateOne {
	ouo.mutation.ResetOauthType()
	ouo.mutation.SetOauthType(i)
	return ouo
}

// AddOauthType adds i to the "oauth_type" field.
func (ouo *OAuthUpdateOne) AddOauthType(i int) *OAuthUpdateOne {
	ouo.mutation.AddOauthType(i)
	return ouo
}

// SetOpenID sets the "open_id" field.
func (ouo *OAuthUpdateOne) SetOpenID(s string) *OAuthUpdateOne {
	ouo.mutation.SetOpenID(s)
	return ouo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (ouo *OAuthUpdateOne) SetAccountID(id int) *OAuthUpdateOne {
	ouo.mutation.SetAccountID(id)
	return ouo
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (ouo *OAuthUpdateOne) SetNillableAccountID(id *int) *OAuthUpdateOne {
	if id != nil {
		ouo = ouo.SetAccountID(*id)
	}
	return ouo
}

// SetAccount sets the "account" edge to the Account entity.
func (ouo *OAuthUpdateOne) SetAccount(a *Account) *OAuthUpdateOne {
	return ouo.SetAccountID(a.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (ouo *OAuthUpdateOne) Mutation() *OAuthMutation {
	return ouo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (ouo *OAuthUpdateOne) ClearAccount() *OAuthUpdateOne {
	ouo.mutation.ClearAccount()
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OAuthUpdateOne) Select(field string, fields ...string) *OAuthUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated OAuth entity.
func (ouo *OAuthUpdateOne) Save(ctx context.Context) (*OAuth, error) {
	var (
		err  error
		node *OAuth
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OAuthUpdateOne) SaveX(ctx context.Context) *OAuth {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OAuthUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OAuthUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdateTime(); !ok {
		v := oauth.UpdateDefaultUpdateTime()
		ouo.mutation.SetUpdateTime(v)
	}
}

func (ouo *OAuthUpdateOne) sqlSave(ctx context.Context) (_node *OAuth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth.Table,
			Columns: oauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth.FieldID)
		for _, f := range fields {
			if !oauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth.FieldUpdateTime,
		})
	}
	if value, ok := ouo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldUpdateBy,
		})
	}
	if value, ok := ouo.mutation.OauthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauth.FieldOauthType,
		})
	}
	if value, ok := ouo.mutation.AddedOauthType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauth.FieldOauthType,
		})
	}
	if value, ok := ouo.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldOpenID,
		})
	}
	if ouo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauth.AccountTable,
			Columns: []string{oauth.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauth.AccountTable,
			Columns: []string{oauth.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OAuth{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
