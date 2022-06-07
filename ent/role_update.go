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
	"github.com/kcmvp/iam.go/ent/application"
	"github.com/kcmvp/iam.go/ent/predicate"
	"github.com/kcmvp/iam.go/ent/role"
	"github.com/kcmvp/iam.go/ent/rolefunc"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RoleUpdate) SetUpdateTime(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetUpdateBy sets the "update_by" field.
func (ru *RoleUpdate) SetUpdateBy(s string) *RoleUpdate {
	ru.mutation.SetUpdateBy(s)
	return ru
}

// SetDeleted sets the "deleted" field.
func (ru *RoleUpdate) SetDeleted(b bool) *RoleUpdate {
	ru.mutation.SetDeleted(b)
	return ru
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDeleted(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetDeleted(*b)
	}
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDesc sets the "desc" field.
func (ru *RoleUpdate) SetDesc(s string) *RoleUpdate {
	ru.mutation.SetDesc(s)
	return ru
}

// AddApplicationIDs adds the "application" edge to the Application entity by IDs.
func (ru *RoleUpdate) AddApplicationIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddApplicationIDs(ids...)
	return ru
}

// AddApplication adds the "application" edges to the Application entity.
func (ru *RoleUpdate) AddApplication(a ...*Application) *RoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddApplicationIDs(ids...)
}

// AddFuncIDs adds the "funcs" edge to the RoleFunc entity by IDs.
func (ru *RoleUpdate) AddFuncIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddFuncIDs(ids...)
	return ru
}

// AddFuncs adds the "funcs" edges to the RoleFunc entity.
func (ru *RoleUpdate) AddFuncs(r ...*RoleFunc) *RoleUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddFuncIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearApplication clears all "application" edges to the Application entity.
func (ru *RoleUpdate) ClearApplication() *RoleUpdate {
	ru.mutation.ClearApplication()
	return ru
}

// RemoveApplicationIDs removes the "application" edge to Application entities by IDs.
func (ru *RoleUpdate) RemoveApplicationIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveApplicationIDs(ids...)
	return ru
}

// RemoveApplication removes "application" edges to Application entities.
func (ru *RoleUpdate) RemoveApplication(a ...*Application) *RoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveApplicationIDs(ids...)
}

// ClearFuncs clears all "funcs" edges to the RoleFunc entity.
func (ru *RoleUpdate) ClearFuncs() *RoleUpdate {
	ru.mutation.ClearFuncs()
	return ru
}

// RemoveFuncIDs removes the "funcs" edge to RoleFunc entities by IDs.
func (ru *RoleUpdate) RemoveFuncIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveFuncIDs(ids...)
	return ru
}

// RemoveFuncs removes "funcs" edges to RoleFunc entities.
func (ru *RoleUpdate) RemoveFuncs(r ...*RoleFunc) *RoleUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveFuncIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := role.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdateTime,
		})
	}
	if value, ok := ru.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldUpdateBy,
		})
	}
	if value, ok := ru.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: role.FieldDeleted,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ru.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDesc,
		})
	}
	if ru.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedApplicationIDs(); len(nodes) > 0 && !ru.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ApplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.FuncsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedFuncsIDs(); len(nodes) > 0 && !ru.mutation.FuncsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.FuncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RoleUpdateOne) SetUpdateTime(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetUpdateBy sets the "update_by" field.
func (ruo *RoleUpdateOne) SetUpdateBy(s string) *RoleUpdateOne {
	ruo.mutation.SetUpdateBy(s)
	return ruo
}

// SetDeleted sets the "deleted" field.
func (ruo *RoleUpdateOne) SetDeleted(b bool) *RoleUpdateOne {
	ruo.mutation.SetDeleted(b)
	return ruo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDeleted(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetDeleted(*b)
	}
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDesc sets the "desc" field.
func (ruo *RoleUpdateOne) SetDesc(s string) *RoleUpdateOne {
	ruo.mutation.SetDesc(s)
	return ruo
}

// AddApplicationIDs adds the "application" edge to the Application entity by IDs.
func (ruo *RoleUpdateOne) AddApplicationIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddApplicationIDs(ids...)
	return ruo
}

// AddApplication adds the "application" edges to the Application entity.
func (ruo *RoleUpdateOne) AddApplication(a ...*Application) *RoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddApplicationIDs(ids...)
}

// AddFuncIDs adds the "funcs" edge to the RoleFunc entity by IDs.
func (ruo *RoleUpdateOne) AddFuncIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddFuncIDs(ids...)
	return ruo
}

// AddFuncs adds the "funcs" edges to the RoleFunc entity.
func (ruo *RoleUpdateOne) AddFuncs(r ...*RoleFunc) *RoleUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddFuncIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearApplication clears all "application" edges to the Application entity.
func (ruo *RoleUpdateOne) ClearApplication() *RoleUpdateOne {
	ruo.mutation.ClearApplication()
	return ruo
}

// RemoveApplicationIDs removes the "application" edge to Application entities by IDs.
func (ruo *RoleUpdateOne) RemoveApplicationIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveApplicationIDs(ids...)
	return ruo
}

// RemoveApplication removes "application" edges to Application entities.
func (ruo *RoleUpdateOne) RemoveApplication(a ...*Application) *RoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveApplicationIDs(ids...)
}

// ClearFuncs clears all "funcs" edges to the RoleFunc entity.
func (ruo *RoleUpdateOne) ClearFuncs() *RoleUpdateOne {
	ruo.mutation.ClearFuncs()
	return ruo
}

// RemoveFuncIDs removes the "funcs" edge to RoleFunc entities by IDs.
func (ruo *RoleUpdateOne) RemoveFuncIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveFuncIDs(ids...)
	return ruo
}

// RemoveFuncs removes "funcs" edges to RoleFunc entities.
func (ruo *RoleUpdateOne) RemoveFuncs(r ...*RoleFunc) *RoleUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveFuncIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	var (
		err  error
		node *Role
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := role.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdateTime,
		})
	}
	if value, ok := ruo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldUpdateBy,
		})
	}
	if value, ok := ruo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: role.FieldDeleted,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ruo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDesc,
		})
	}
	if ruo.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedApplicationIDs(); len(nodes) > 0 && !ruo.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ApplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.ApplicationTable,
			Columns: role.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.FuncsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedFuncsIDs(); len(nodes) > 0 && !ruo.mutation.FuncsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.FuncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.FuncsTable,
			Columns: role.FuncsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolefunc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
