// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcmvp/idm/ent/predicate"
	"github.com/kcmvp/idm/ent/role"
	"github.com/kcmvp/idm/ent/rolefunc"
)

// RoleFuncQuery is the builder for querying RoleFunc entities.
type RoleFuncQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RoleFunc
	// eager-loading edges.
	withRole *RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleFuncQuery builder.
func (rfq *RoleFuncQuery) Where(ps ...predicate.RoleFunc) *RoleFuncQuery {
	rfq.predicates = append(rfq.predicates, ps...)
	return rfq
}

// Limit adds a limit step to the query.
func (rfq *RoleFuncQuery) Limit(limit int) *RoleFuncQuery {
	rfq.limit = &limit
	return rfq
}

// Offset adds an offset step to the query.
func (rfq *RoleFuncQuery) Offset(offset int) *RoleFuncQuery {
	rfq.offset = &offset
	return rfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rfq *RoleFuncQuery) Unique(unique bool) *RoleFuncQuery {
	rfq.unique = &unique
	return rfq
}

// Order adds an order step to the query.
func (rfq *RoleFuncQuery) Order(o ...OrderFunc) *RoleFuncQuery {
	rfq.order = append(rfq.order, o...)
	return rfq
}

// QueryRole chains the current query on the "role" edge.
func (rfq *RoleFuncQuery) QueryRole() *RoleQuery {
	query := &RoleQuery{config: rfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolefunc.Table, rolefunc.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, rolefunc.RoleTable, rolefunc.RolePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoleFunc entity from the query.
// Returns a *NotFoundError when no RoleFunc was found.
func (rfq *RoleFuncQuery) First(ctx context.Context) (*RoleFunc, error) {
	nodes, err := rfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolefunc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rfq *RoleFuncQuery) FirstX(ctx context.Context) *RoleFunc {
	node, err := rfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoleFunc ID from the query.
// Returns a *NotFoundError when no RoleFunc ID was found.
func (rfq *RoleFuncQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rolefunc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rfq *RoleFuncQuery) FirstIDX(ctx context.Context) int {
	id, err := rfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoleFunc entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoleFunc entity is found.
// Returns a *NotFoundError when no RoleFunc entities are found.
func (rfq *RoleFuncQuery) Only(ctx context.Context) (*RoleFunc, error) {
	nodes, err := rfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolefunc.Label}
	default:
		return nil, &NotSingularError{rolefunc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rfq *RoleFuncQuery) OnlyX(ctx context.Context) *RoleFunc {
	node, err := rfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoleFunc ID in the query.
// Returns a *NotSingularError when more than one RoleFunc ID is found.
// Returns a *NotFoundError when no entities are found.
func (rfq *RoleFuncQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = &NotSingularError{rolefunc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rfq *RoleFuncQuery) OnlyIDX(ctx context.Context) int {
	id, err := rfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoleFuncs.
func (rfq *RoleFuncQuery) All(ctx context.Context) ([]*RoleFunc, error) {
	if err := rfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rfq *RoleFuncQuery) AllX(ctx context.Context) []*RoleFunc {
	nodes, err := rfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoleFunc IDs.
func (rfq *RoleFuncQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rfq.Select(rolefunc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rfq *RoleFuncQuery) IDsX(ctx context.Context) []int {
	ids, err := rfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rfq *RoleFuncQuery) Count(ctx context.Context) (int, error) {
	if err := rfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rfq *RoleFuncQuery) CountX(ctx context.Context) int {
	count, err := rfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rfq *RoleFuncQuery) Exist(ctx context.Context) (bool, error) {
	if err := rfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rfq *RoleFuncQuery) ExistX(ctx context.Context) bool {
	exist, err := rfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleFuncQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rfq *RoleFuncQuery) Clone() *RoleFuncQuery {
	if rfq == nil {
		return nil
	}
	return &RoleFuncQuery{
		config:     rfq.config,
		limit:      rfq.limit,
		offset:     rfq.offset,
		order:      append([]OrderFunc{}, rfq.order...),
		predicates: append([]predicate.RoleFunc{}, rfq.predicates...),
		withRole:   rfq.withRole.Clone(),
		// clone intermediate query.
		sql:    rfq.sql.Clone(),
		path:   rfq.path,
		unique: rfq.unique,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (rfq *RoleFuncQuery) WithRole(opts ...func(*RoleQuery)) *RoleFuncQuery {
	query := &RoleQuery{config: rfq.config}
	for _, opt := range opts {
		opt(query)
	}
	rfq.withRole = query
	return rfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoleFunc.Query().
//		GroupBy(rolefunc.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rfq *RoleFuncQuery) GroupBy(field string, fields ...string) *RoleFuncGroupBy {
	group := &RoleFuncGroupBy{config: rfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.RoleFunc.Query().
//		Select(rolefunc.FieldCreateTime).
//		Scan(ctx, &v)
//
func (rfq *RoleFuncQuery) Select(fields ...string) *RoleFuncSelect {
	rfq.fields = append(rfq.fields, fields...)
	return &RoleFuncSelect{RoleFuncQuery: rfq}
}

func (rfq *RoleFuncQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rfq.fields {
		if !rolefunc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rfq.path != nil {
		prev, err := rfq.path(ctx)
		if err != nil {
			return err
		}
		rfq.sql = prev
	}
	return nil
}

func (rfq *RoleFuncQuery) sqlAll(ctx context.Context) ([]*RoleFunc, error) {
	var (
		nodes       = []*RoleFunc{}
		_spec       = rfq.querySpec()
		loadedTypes = [1]bool{
			rfq.withRole != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RoleFunc{config: rfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rfq.withRole; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*RoleFunc, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Role = []*Role{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*RoleFunc)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   rolefunc.RoleTable,
				Columns: rolefunc.RolePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(rolefunc.RolePrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, rfq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "role": %w`, err)
		}
		query.Where(role.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "role" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Role = append(nodes[i].Edges.Role, n)
			}
		}
	}

	return nodes, nil
}

func (rfq *RoleFuncQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rfq.querySpec()
	_spec.Node.Columns = rfq.fields
	if len(rfq.fields) > 0 {
		_spec.Unique = rfq.unique != nil && *rfq.unique
	}
	return sqlgraph.CountNodes(ctx, rfq.driver, _spec)
}

func (rfq *RoleFuncQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rfq *RoleFuncQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rolefunc.Table,
			Columns: rolefunc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolefunc.FieldID,
			},
		},
		From:   rfq.sql,
		Unique: true,
	}
	if unique := rfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolefunc.FieldID)
		for i := range fields {
			if fields[i] != rolefunc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rfq *RoleFuncQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rfq.driver.Dialect())
	t1 := builder.Table(rolefunc.Table)
	columns := rfq.fields
	if len(columns) == 0 {
		columns = rolefunc.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rfq.sql != nil {
		selector = rfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rfq.unique != nil && *rfq.unique {
		selector.Distinct()
	}
	for _, p := range rfq.predicates {
		p(selector)
	}
	for _, p := range rfq.order {
		p(selector)
	}
	if offset := rfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoleFuncGroupBy is the group-by builder for RoleFunc entities.
type RoleFuncGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rfgb *RoleFuncGroupBy) Aggregate(fns ...AggregateFunc) *RoleFuncGroupBy {
	rfgb.fns = append(rfgb.fns, fns...)
	return rfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rfgb *RoleFuncGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rfgb.path(ctx)
	if err != nil {
		return err
	}
	rfgb.sql = query
	return rfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rfgb.fields) > 1 {
		return nil, errors.New("ent: RoleFuncGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) StringsX(ctx context.Context) []string {
	v, err := rfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) StringX(ctx context.Context) string {
	v, err := rfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rfgb.fields) > 1 {
		return nil, errors.New("ent: RoleFuncGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) IntsX(ctx context.Context) []int {
	v, err := rfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) IntX(ctx context.Context) int {
	v, err := rfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rfgb.fields) > 1 {
		return nil, errors.New("ent: RoleFuncGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rfgb.fields) > 1 {
		return nil, errors.New("ent: RoleFuncGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rfgb *RoleFuncGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rfgb *RoleFuncGroupBy) BoolX(ctx context.Context) bool {
	v, err := rfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rfgb *RoleFuncGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rfgb.fields {
		if !rolefunc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rfgb *RoleFuncGroupBy) sqlQuery() *sql.Selector {
	selector := rfgb.sql.Select()
	aggregation := make([]string, 0, len(rfgb.fns))
	for _, fn := range rfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rfgb.fields)+len(rfgb.fns))
		for _, f := range rfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rfgb.fields...)...)
}

// RoleFuncSelect is the builder for selecting fields of RoleFunc entities.
type RoleFuncSelect struct {
	*RoleFuncQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rfs *RoleFuncSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rfs.prepareQuery(ctx); err != nil {
		return err
	}
	rfs.sql = rfs.RoleFuncQuery.sqlQuery(ctx)
	return rfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rfs *RoleFuncSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rfs.fields) > 1 {
		return nil, errors.New("ent: RoleFuncSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rfs *RoleFuncSelect) StringsX(ctx context.Context) []string {
	v, err := rfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rfs *RoleFuncSelect) StringX(ctx context.Context) string {
	v, err := rfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rfs.fields) > 1 {
		return nil, errors.New("ent: RoleFuncSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rfs *RoleFuncSelect) IntsX(ctx context.Context) []int {
	v, err := rfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rfs *RoleFuncSelect) IntX(ctx context.Context) int {
	v, err := rfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rfs.fields) > 1 {
		return nil, errors.New("ent: RoleFuncSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rfs *RoleFuncSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rfs *RoleFuncSelect) Float64X(ctx context.Context) float64 {
	v, err := rfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rfs.fields) > 1 {
		return nil, errors.New("ent: RoleFuncSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rfs *RoleFuncSelect) BoolsX(ctx context.Context) []bool {
	v, err := rfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rfs *RoleFuncSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolefunc.Label}
	default:
		err = fmt.Errorf("ent: RoleFuncSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rfs *RoleFuncSelect) BoolX(ctx context.Context) bool {
	v, err := rfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rfs *RoleFuncSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rfs.sql.Query()
	if err := rfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
