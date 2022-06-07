// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcmvp/iam.go/ent/predicate"
	"github.com/kcmvp/iam.go/ent/rolefunc"
)

// RoleFuncDelete is the builder for deleting a RoleFunc entity.
type RoleFuncDelete struct {
	config
	hooks    []Hook
	mutation *RoleFuncMutation
}

// Where appends a list predicates to the RoleFuncDelete builder.
func (rfd *RoleFuncDelete) Where(ps ...predicate.RoleFunc) *RoleFuncDelete {
	rfd.mutation.Where(ps...)
	return rfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rfd *RoleFuncDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rfd.hooks) == 0 {
		affected, err = rfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleFuncMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfd.mutation = mutation
			affected, err = rfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rfd.hooks) - 1; i >= 0; i-- {
			if rfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfd *RoleFuncDelete) ExecX(ctx context.Context) int {
	n, err := rfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rfd *RoleFuncDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: rolefunc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolefunc.FieldID,
			},
		},
	}
	if ps := rfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rfd.driver, _spec)
}

// RoleFuncDeleteOne is the builder for deleting a single RoleFunc entity.
type RoleFuncDeleteOne struct {
	rfd *RoleFuncDelete
}

// Exec executes the deletion query.
func (rfdo *RoleFuncDeleteOne) Exec(ctx context.Context) error {
	n, err := rfdo.rfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rolefunc.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rfdo *RoleFuncDeleteOne) ExecX(ctx context.Context) {
	rfdo.rfd.ExecX(ctx)
}