// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcmvp/iam.go/ent/role"
	"github.com/kcmvp/iam.go/ent/rolefunc"
)

// RoleFuncCreate is the builder for creating a RoleFunc entity.
type RoleFuncCreate struct {
	config
	mutation *RoleFuncMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rfc *RoleFuncCreate) SetCreateTime(t time.Time) *RoleFuncCreate {
	rfc.mutation.SetCreateTime(t)
	return rfc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rfc *RoleFuncCreate) SetNillableCreateTime(t *time.Time) *RoleFuncCreate {
	if t != nil {
		rfc.SetCreateTime(*t)
	}
	return rfc
}

// SetUpdateTime sets the "update_time" field.
func (rfc *RoleFuncCreate) SetUpdateTime(t time.Time) *RoleFuncCreate {
	rfc.mutation.SetUpdateTime(t)
	return rfc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rfc *RoleFuncCreate) SetNillableUpdateTime(t *time.Time) *RoleFuncCreate {
	if t != nil {
		rfc.SetUpdateTime(*t)
	}
	return rfc
}

// SetCreateBy sets the "create_by" field.
func (rfc *RoleFuncCreate) SetCreateBy(s string) *RoleFuncCreate {
	rfc.mutation.SetCreateBy(s)
	return rfc
}

// SetUpdateBy sets the "update_by" field.
func (rfc *RoleFuncCreate) SetUpdateBy(s string) *RoleFuncCreate {
	rfc.mutation.SetUpdateBy(s)
	return rfc
}

// SetDeleted sets the "deleted" field.
func (rfc *RoleFuncCreate) SetDeleted(b bool) *RoleFuncCreate {
	rfc.mutation.SetDeleted(b)
	return rfc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (rfc *RoleFuncCreate) SetNillableDeleted(b *bool) *RoleFuncCreate {
	if b != nil {
		rfc.SetDeleted(*b)
	}
	return rfc
}

// SetURLPattern sets the "url_pattern" field.
func (rfc *RoleFuncCreate) SetURLPattern(s string) *RoleFuncCreate {
	rfc.mutation.SetURLPattern(s)
	return rfc
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (rfc *RoleFuncCreate) AddRoleIDs(ids ...int) *RoleFuncCreate {
	rfc.mutation.AddRoleIDs(ids...)
	return rfc
}

// AddRole adds the "role" edges to the Role entity.
func (rfc *RoleFuncCreate) AddRole(r ...*Role) *RoleFuncCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rfc.AddRoleIDs(ids...)
}

// Mutation returns the RoleFuncMutation object of the builder.
func (rfc *RoleFuncCreate) Mutation() *RoleFuncMutation {
	return rfc.mutation
}

// Save creates the RoleFunc in the database.
func (rfc *RoleFuncCreate) Save(ctx context.Context) (*RoleFunc, error) {
	var (
		err  error
		node *RoleFunc
	)
	rfc.defaults()
	if len(rfc.hooks) == 0 {
		if err = rfc.check(); err != nil {
			return nil, err
		}
		node, err = rfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleFuncMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rfc.check(); err != nil {
				return nil, err
			}
			rfc.mutation = mutation
			if node, err = rfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rfc.hooks) - 1; i >= 0; i-- {
			if rfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rfc *RoleFuncCreate) SaveX(ctx context.Context) *RoleFunc {
	v, err := rfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rfc *RoleFuncCreate) Exec(ctx context.Context) error {
	_, err := rfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfc *RoleFuncCreate) ExecX(ctx context.Context) {
	if err := rfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rfc *RoleFuncCreate) defaults() {
	if _, ok := rfc.mutation.CreateTime(); !ok {
		v := rolefunc.DefaultCreateTime()
		rfc.mutation.SetCreateTime(v)
	}
	if _, ok := rfc.mutation.UpdateTime(); !ok {
		v := rolefunc.DefaultUpdateTime()
		rfc.mutation.SetUpdateTime(v)
	}
	if _, ok := rfc.mutation.Deleted(); !ok {
		v := rolefunc.DefaultDeleted
		rfc.mutation.SetDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rfc *RoleFuncCreate) check() error {
	if _, ok := rfc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "RoleFunc.create_time"`)}
	}
	if _, ok := rfc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "RoleFunc.update_time"`)}
	}
	if _, ok := rfc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "create_by", err: errors.New(`ent: missing required field "RoleFunc.create_by"`)}
	}
	if _, ok := rfc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "RoleFunc.update_by"`)}
	}
	if _, ok := rfc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "RoleFunc.deleted"`)}
	}
	if _, ok := rfc.mutation.URLPattern(); !ok {
		return &ValidationError{Name: "url_pattern", err: errors.New(`ent: missing required field "RoleFunc.url_pattern"`)}
	}
	if v, ok := rfc.mutation.URLPattern(); ok {
		if err := rolefunc.URLPatternValidator(v); err != nil {
			return &ValidationError{Name: "url_pattern", err: fmt.Errorf(`ent: validator failed for field "RoleFunc.url_pattern": %w`, err)}
		}
	}
	return nil
}

func (rfc *RoleFuncCreate) sqlSave(ctx context.Context) (*RoleFunc, error) {
	_node, _spec := rfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rfc *RoleFuncCreate) createSpec() (*RoleFunc, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleFunc{config: rfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rolefunc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolefunc.FieldID,
			},
		}
	)
	if value, ok := rfc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rolefunc.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := rfc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rolefunc.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := rfc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolefunc.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := rfc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolefunc.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := rfc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rolefunc.FieldDeleted,
		})
		_node.Deleted = value
	}
	if value, ok := rfc.mutation.URLPattern(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolefunc.FieldURLPattern,
		})
		_node.URLPattern = value
	}
	if nodes := rfc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rolefunc.RoleTable,
			Columns: rolefunc.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleFuncCreateBulk is the builder for creating many RoleFunc entities in bulk.
type RoleFuncCreateBulk struct {
	config
	builders []*RoleFuncCreate
}

// Save creates the RoleFunc entities in the database.
func (rfcb *RoleFuncCreateBulk) Save(ctx context.Context) ([]*RoleFunc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rfcb.builders))
	nodes := make([]*RoleFunc, len(rfcb.builders))
	mutators := make([]Mutator, len(rfcb.builders))
	for i := range rfcb.builders {
		func(i int, root context.Context) {
			builder := rfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleFuncMutation)
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
					_, err = mutators[i+1].Mutate(root, rfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rfcb *RoleFuncCreateBulk) SaveX(ctx context.Context) []*RoleFunc {
	v, err := rfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rfcb *RoleFuncCreateBulk) Exec(ctx context.Context) error {
	_, err := rfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfcb *RoleFuncCreateBulk) ExecX(ctx context.Context) {
	if err := rfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
