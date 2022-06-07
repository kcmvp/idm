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
	"github.com/kcmvp/idm.go/ent/account"
	"github.com/kcmvp/idm.go/ent/predicate"
	"github.com/kcmvp/idm.go/ent/schema"
	"github.com/kcmvp/idm.go/ent/subaccount"
)

// SubAccountUpdate is the builder for updating SubAccount entities.
type SubAccountUpdate struct {
	config
	hooks    []Hook
	mutation *SubAccountMutation
}

// Where appends a list predicates to the SubAccountUpdate builder.
func (sau *SubAccountUpdate) Where(ps ...predicate.SubAccount) *SubAccountUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetUpdateTime sets the "update_time" field.
func (sau *SubAccountUpdate) SetUpdateTime(t time.Time) *SubAccountUpdate {
	sau.mutation.SetUpdateTime(t)
	return sau
}

// SetUpdateBy sets the "update_by" field.
func (sau *SubAccountUpdate) SetUpdateBy(s string) *SubAccountUpdate {
	sau.mutation.SetUpdateBy(s)
	return sau
}

// SetDeleted sets the "deleted" field.
func (sau *SubAccountUpdate) SetDeleted(b bool) *SubAccountUpdate {
	sau.mutation.SetDeleted(b)
	return sau
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (sau *SubAccountUpdate) SetNillableDeleted(b *bool) *SubAccountUpdate {
	if b != nil {
		sau.SetDeleted(*b)
	}
	return sau
}

// SetAcctType sets the "acct_type" field.
func (sau *SubAccountUpdate) SetAcctType(st schema.AcctType) *SubAccountUpdate {
	sau.mutation.ResetAcctType()
	sau.mutation.SetAcctType(st)
	return sau
}

// AddAcctType adds st to the "acct_type" field.
func (sau *SubAccountUpdate) AddAcctType(st schema.AcctType) *SubAccountUpdate {
	sau.mutation.AddAcctType(st)
	return sau
}

// SetSubAcct sets the "sub_acct" field.
func (sau *SubAccountUpdate) SetSubAcct(s string) *SubAccountUpdate {
	sau.mutation.SetSubAcct(s)
	return sau
}

// AddAccounIDs adds the "accoun" edge to the Account entity by IDs.
func (sau *SubAccountUpdate) AddAccounIDs(ids ...int) *SubAccountUpdate {
	sau.mutation.AddAccounIDs(ids...)
	return sau
}

// AddAccoun adds the "accoun" edges to the Account entity.
func (sau *SubAccountUpdate) AddAccoun(a ...*Account) *SubAccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.AddAccounIDs(ids...)
}

// Mutation returns the SubAccountMutation object of the builder.
func (sau *SubAccountUpdate) Mutation() *SubAccountMutation {
	return sau.mutation
}

// ClearAccoun clears all "accoun" edges to the Account entity.
func (sau *SubAccountUpdate) ClearAccoun() *SubAccountUpdate {
	sau.mutation.ClearAccoun()
	return sau
}

// RemoveAccounIDs removes the "accoun" edge to Account entities by IDs.
func (sau *SubAccountUpdate) RemoveAccounIDs(ids ...int) *SubAccountUpdate {
	sau.mutation.RemoveAccounIDs(ids...)
	return sau
}

// RemoveAccoun removes "accoun" edges to Account entities.
func (sau *SubAccountUpdate) RemoveAccoun(a ...*Account) *SubAccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.RemoveAccounIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SubAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sau.defaults()
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SubAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SubAccountUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SubAccountUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *SubAccountUpdate) defaults() {
	if _, ok := sau.mutation.UpdateTime(); !ok {
		v := subaccount.UpdateDefaultUpdateTime()
		sau.mutation.SetUpdateTime(v)
	}
}

func (sau *SubAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subaccount.Table,
			Columns: subaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subaccount.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subaccount.FieldUpdateTime,
		})
	}
	if value, ok := sau.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subaccount.FieldUpdateBy,
		})
	}
	if value, ok := sau.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subaccount.FieldDeleted,
		})
	}
	if value, ok := sau.mutation.AcctType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subaccount.FieldAcctType,
		})
	}
	if value, ok := sau.mutation.AddedAcctType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subaccount.FieldAcctType,
		})
	}
	if value, ok := sau.mutation.SubAcct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subaccount.FieldSubAcct,
		})
	}
	if sau.mutation.AccounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
	if nodes := sau.mutation.RemovedAccounIDs(); len(nodes) > 0 && !sau.mutation.AccounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.AccounIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SubAccountUpdateOne is the builder for updating a single SubAccount entity.
type SubAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubAccountMutation
}

// SetUpdateTime sets the "update_time" field.
func (sauo *SubAccountUpdateOne) SetUpdateTime(t time.Time) *SubAccountUpdateOne {
	sauo.mutation.SetUpdateTime(t)
	return sauo
}

// SetUpdateBy sets the "update_by" field.
func (sauo *SubAccountUpdateOne) SetUpdateBy(s string) *SubAccountUpdateOne {
	sauo.mutation.SetUpdateBy(s)
	return sauo
}

// SetDeleted sets the "deleted" field.
func (sauo *SubAccountUpdateOne) SetDeleted(b bool) *SubAccountUpdateOne {
	sauo.mutation.SetDeleted(b)
	return sauo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (sauo *SubAccountUpdateOne) SetNillableDeleted(b *bool) *SubAccountUpdateOne {
	if b != nil {
		sauo.SetDeleted(*b)
	}
	return sauo
}

// SetAcctType sets the "acct_type" field.
func (sauo *SubAccountUpdateOne) SetAcctType(st schema.AcctType) *SubAccountUpdateOne {
	sauo.mutation.ResetAcctType()
	sauo.mutation.SetAcctType(st)
	return sauo
}

// AddAcctType adds st to the "acct_type" field.
func (sauo *SubAccountUpdateOne) AddAcctType(st schema.AcctType) *SubAccountUpdateOne {
	sauo.mutation.AddAcctType(st)
	return sauo
}

// SetSubAcct sets the "sub_acct" field.
func (sauo *SubAccountUpdateOne) SetSubAcct(s string) *SubAccountUpdateOne {
	sauo.mutation.SetSubAcct(s)
	return sauo
}

// AddAccounIDs adds the "accoun" edge to the Account entity by IDs.
func (sauo *SubAccountUpdateOne) AddAccounIDs(ids ...int) *SubAccountUpdateOne {
	sauo.mutation.AddAccounIDs(ids...)
	return sauo
}

// AddAccoun adds the "accoun" edges to the Account entity.
func (sauo *SubAccountUpdateOne) AddAccoun(a ...*Account) *SubAccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.AddAccounIDs(ids...)
}

// Mutation returns the SubAccountMutation object of the builder.
func (sauo *SubAccountUpdateOne) Mutation() *SubAccountMutation {
	return sauo.mutation
}

// ClearAccoun clears all "accoun" edges to the Account entity.
func (sauo *SubAccountUpdateOne) ClearAccoun() *SubAccountUpdateOne {
	sauo.mutation.ClearAccoun()
	return sauo
}

// RemoveAccounIDs removes the "accoun" edge to Account entities by IDs.
func (sauo *SubAccountUpdateOne) RemoveAccounIDs(ids ...int) *SubAccountUpdateOne {
	sauo.mutation.RemoveAccounIDs(ids...)
	return sauo
}

// RemoveAccoun removes "accoun" edges to Account entities.
func (sauo *SubAccountUpdateOne) RemoveAccoun(a ...*Account) *SubAccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.RemoveAccounIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SubAccountUpdateOne) Select(field string, fields ...string) *SubAccountUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SubAccount entity.
func (sauo *SubAccountUpdateOne) Save(ctx context.Context) (*SubAccount, error) {
	var (
		err  error
		node *SubAccount
	)
	sauo.defaults()
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SubAccountUpdateOne) SaveX(ctx context.Context) *SubAccount {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SubAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SubAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *SubAccountUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdateTime(); !ok {
		v := subaccount.UpdateDefaultUpdateTime()
		sauo.mutation.SetUpdateTime(v)
	}
}

func (sauo *SubAccountUpdateOne) sqlSave(ctx context.Context) (_node *SubAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subaccount.Table,
			Columns: subaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subaccount.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subaccount.FieldID)
		for _, f := range fields {
			if !subaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subaccount.FieldUpdateTime,
		})
	}
	if value, ok := sauo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subaccount.FieldUpdateBy,
		})
	}
	if value, ok := sauo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subaccount.FieldDeleted,
		})
	}
	if value, ok := sauo.mutation.AcctType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subaccount.FieldAcctType,
		})
	}
	if value, ok := sauo.mutation.AddedAcctType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subaccount.FieldAcctType,
		})
	}
	if value, ok := sauo.mutation.SubAcct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subaccount.FieldSubAcct,
		})
	}
	if sauo.mutation.AccounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
	if nodes := sauo.mutation.RemovedAccounIDs(); len(nodes) > 0 && !sauo.mutation.AccounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.AccounIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subaccount.AccounTable,
			Columns: subaccount.AccounPrimaryKey,
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
	_node = &SubAccount{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
