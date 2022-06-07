// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcmvp/iam.go/ent/account"
	"github.com/kcmvp/iam.go/ent/oauth"
)

// OAuthCreate is the builder for creating a OAuth entity.
type OAuthCreate struct {
	config
	mutation *OAuthMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oc *OAuthCreate) SetCreateTime(t time.Time) *OAuthCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *OAuthCreate) SetNillableCreateTime(t *time.Time) *OAuthCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *OAuthCreate) SetUpdateTime(t time.Time) *OAuthCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *OAuthCreate) SetNillableUpdateTime(t *time.Time) *OAuthCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetCreateBy sets the "create_by" field.
func (oc *OAuthCreate) SetCreateBy(s string) *OAuthCreate {
	oc.mutation.SetCreateBy(s)
	return oc
}

// SetUpdateBy sets the "update_by" field.
func (oc *OAuthCreate) SetUpdateBy(s string) *OAuthCreate {
	oc.mutation.SetUpdateBy(s)
	return oc
}

// SetOauthType sets the "oauth_type" field.
func (oc *OAuthCreate) SetOauthType(i int) *OAuthCreate {
	oc.mutation.SetOauthType(i)
	return oc
}

// SetOpenID sets the "open_id" field.
func (oc *OAuthCreate) SetOpenID(s string) *OAuthCreate {
	oc.mutation.SetOpenID(s)
	return oc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (oc *OAuthCreate) SetAccountID(id int) *OAuthCreate {
	oc.mutation.SetAccountID(id)
	return oc
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (oc *OAuthCreate) SetNillableAccountID(id *int) *OAuthCreate {
	if id != nil {
		oc = oc.SetAccountID(*id)
	}
	return oc
}

// SetAccount sets the "account" edge to the Account entity.
func (oc *OAuthCreate) SetAccount(a *Account) *OAuthCreate {
	return oc.SetAccountID(a.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (oc *OAuthCreate) Mutation() *OAuthMutation {
	return oc.mutation
}

// Save creates the OAuth in the database.
func (oc *OAuthCreate) Save(ctx context.Context) (*OAuth, error) {
	var (
		err  error
		node *OAuth
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OAuthCreate) SaveX(ctx context.Context) *OAuth {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OAuthCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OAuthCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OAuthCreate) defaults() {
	if _, ok := oc.mutation.CreateTime(); !ok {
		v := oauth.DefaultCreateTime()
		oc.mutation.SetCreateTime(v)
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		v := oauth.DefaultUpdateTime()
		oc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OAuthCreate) check() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "OAuth.create_time"`)}
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "OAuth.update_time"`)}
	}
	if _, ok := oc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "create_by", err: errors.New(`ent: missing required field "OAuth.create_by"`)}
	}
	if _, ok := oc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "OAuth.update_by"`)}
	}
	if _, ok := oc.mutation.OauthType(); !ok {
		return &ValidationError{Name: "oauth_type", err: errors.New(`ent: missing required field "OAuth.oauth_type"`)}
	}
	if _, ok := oc.mutation.OpenID(); !ok {
		return &ValidationError{Name: "open_id", err: errors.New(`ent: missing required field "OAuth.open_id"`)}
	}
	return nil
}

func (oc *OAuthCreate) sqlSave(ctx context.Context) (*OAuth, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OAuthCreate) createSpec() (*OAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuth{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: oauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := oc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := oc.mutation.OauthType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauth.FieldOauthType,
		})
		_node.OauthType = value
	}
	if value, ok := oc.mutation.OpenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldOpenID,
		})
		_node.OpenID = value
	}
	if nodes := oc.mutation.AccountIDs(); len(nodes) > 0 {
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
		_node.account_oauth = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OAuthCreateBulk is the builder for creating many OAuth entities in bulk.
type OAuthCreateBulk struct {
	config
	builders []*OAuthCreate
}

// Save creates the OAuth entities in the database.
func (ocb *OAuthCreateBulk) Save(ctx context.Context) ([]*OAuth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*OAuth, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OAuthCreateBulk) SaveX(ctx context.Context) []*OAuth {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OAuthCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}