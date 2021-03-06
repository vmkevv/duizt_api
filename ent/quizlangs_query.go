// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/i18n"
	"github.com/vmkevv/duiztapi/ent/predicate"
	"github.com/vmkevv/duiztapi/ent/quiz"
	"github.com/vmkevv/duiztapi/ent/quizlangs"
)

// QuizLangsQuery is the builder for querying QuizLangs entities.
type QuizLangsQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.QuizLangs
	// eager-loading edges.
	withQuiz *QuizQuery
	withI18n *I18nQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QuizLangsQuery builder.
func (qlq *QuizLangsQuery) Where(ps ...predicate.QuizLangs) *QuizLangsQuery {
	qlq.predicates = append(qlq.predicates, ps...)
	return qlq
}

// Limit adds a limit step to the query.
func (qlq *QuizLangsQuery) Limit(limit int) *QuizLangsQuery {
	qlq.limit = &limit
	return qlq
}

// Offset adds an offset step to the query.
func (qlq *QuizLangsQuery) Offset(offset int) *QuizLangsQuery {
	qlq.offset = &offset
	return qlq
}

// Order adds an order step to the query.
func (qlq *QuizLangsQuery) Order(o ...OrderFunc) *QuizLangsQuery {
	qlq.order = append(qlq.order, o...)
	return qlq
}

// QueryQuiz chains the current query on the "quiz" edge.
func (qlq *QuizLangsQuery) QueryQuiz() *QuizQuery {
	query := &QuizQuery{config: qlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qlq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(quizlangs.Table, quizlangs.FieldID, selector),
			sqlgraph.To(quiz.Table, quiz.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, quizlangs.QuizTable, quizlangs.QuizColumn),
		)
		fromU = sqlgraph.SetNeighbors(qlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryI18n chains the current query on the "i18n" edge.
func (qlq *QuizLangsQuery) QueryI18n() *I18nQuery {
	query := &I18nQuery{config: qlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qlq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(quizlangs.Table, quizlangs.FieldID, selector),
			sqlgraph.To(i18n.Table, i18n.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, quizlangs.I18nTable, quizlangs.I18nColumn),
		)
		fromU = sqlgraph.SetNeighbors(qlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first QuizLangs entity from the query.
// Returns a *NotFoundError when no QuizLangs was found.
func (qlq *QuizLangsQuery) First(ctx context.Context) (*QuizLangs, error) {
	nodes, err := qlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{quizlangs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qlq *QuizLangsQuery) FirstX(ctx context.Context) *QuizLangs {
	node, err := qlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first QuizLangs ID from the query.
// Returns a *NotFoundError when no QuizLangs ID was found.
func (qlq *QuizLangsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{quizlangs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qlq *QuizLangsQuery) FirstIDX(ctx context.Context) int {
	id, err := qlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single QuizLangs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one QuizLangs entity is not found.
// Returns a *NotFoundError when no QuizLangs entities are found.
func (qlq *QuizLangsQuery) Only(ctx context.Context) (*QuizLangs, error) {
	nodes, err := qlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{quizlangs.Label}
	default:
		return nil, &NotSingularError{quizlangs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qlq *QuizLangsQuery) OnlyX(ctx context.Context) *QuizLangs {
	node, err := qlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only QuizLangs ID in the query.
// Returns a *NotSingularError when exactly one QuizLangs ID is not found.
// Returns a *NotFoundError when no entities are found.
func (qlq *QuizLangsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = &NotSingularError{quizlangs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qlq *QuizLangsQuery) OnlyIDX(ctx context.Context) int {
	id, err := qlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QuizLangsSlice.
func (qlq *QuizLangsQuery) All(ctx context.Context) ([]*QuizLangs, error) {
	if err := qlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qlq *QuizLangsQuery) AllX(ctx context.Context) []*QuizLangs {
	nodes, err := qlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of QuizLangs IDs.
func (qlq *QuizLangsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := qlq.Select(quizlangs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qlq *QuizLangsQuery) IDsX(ctx context.Context) []int {
	ids, err := qlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qlq *QuizLangsQuery) Count(ctx context.Context) (int, error) {
	if err := qlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qlq *QuizLangsQuery) CountX(ctx context.Context) int {
	count, err := qlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qlq *QuizLangsQuery) Exist(ctx context.Context) (bool, error) {
	if err := qlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qlq *QuizLangsQuery) ExistX(ctx context.Context) bool {
	exist, err := qlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QuizLangsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qlq *QuizLangsQuery) Clone() *QuizLangsQuery {
	if qlq == nil {
		return nil
	}
	return &QuizLangsQuery{
		config:     qlq.config,
		limit:      qlq.limit,
		offset:     qlq.offset,
		order:      append([]OrderFunc{}, qlq.order...),
		predicates: append([]predicate.QuizLangs{}, qlq.predicates...),
		withQuiz:   qlq.withQuiz.Clone(),
		withI18n:   qlq.withI18n.Clone(),
		// clone intermediate query.
		sql:  qlq.sql.Clone(),
		path: qlq.path,
	}
}

// WithQuiz tells the query-builder to eager-load the nodes that are connected to
// the "quiz" edge. The optional arguments are used to configure the query builder of the edge.
func (qlq *QuizLangsQuery) WithQuiz(opts ...func(*QuizQuery)) *QuizLangsQuery {
	query := &QuizQuery{config: qlq.config}
	for _, opt := range opts {
		opt(query)
	}
	qlq.withQuiz = query
	return qlq
}

// WithI18n tells the query-builder to eager-load the nodes that are connected to
// the "i18n" edge. The optional arguments are used to configure the query builder of the edge.
func (qlq *QuizLangsQuery) WithI18n(opts ...func(*I18nQuery)) *QuizLangsQuery {
	query := &I18nQuery{config: qlq.config}
	for _, opt := range opts {
		opt(query)
	}
	qlq.withI18n = query
	return qlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QuizLangs.Query().
//		GroupBy(quizlangs.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (qlq *QuizLangsQuery) GroupBy(field string, fields ...string) *QuizLangsGroupBy {
	group := &QuizLangsGroupBy{config: qlq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qlq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.QuizLangs.Query().
//		Select(quizlangs.FieldName).
//		Scan(ctx, &v)
//
func (qlq *QuizLangsQuery) Select(field string, fields ...string) *QuizLangsSelect {
	qlq.fields = append([]string{field}, fields...)
	return &QuizLangsSelect{QuizLangsQuery: qlq}
}

func (qlq *QuizLangsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range qlq.fields {
		if !quizlangs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qlq.path != nil {
		prev, err := qlq.path(ctx)
		if err != nil {
			return err
		}
		qlq.sql = prev
	}
	return nil
}

func (qlq *QuizLangsQuery) sqlAll(ctx context.Context) ([]*QuizLangs, error) {
	var (
		nodes       = []*QuizLangs{}
		withFKs     = qlq.withFKs
		_spec       = qlq.querySpec()
		loadedTypes = [2]bool{
			qlq.withQuiz != nil,
			qlq.withI18n != nil,
		}
	)
	if qlq.withQuiz != nil || qlq.withI18n != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, quizlangs.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &QuizLangs{config: qlq.config}
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
	if err := sqlgraph.QueryNodes(ctx, qlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := qlq.withQuiz; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*QuizLangs)
		for i := range nodes {
			if fk := nodes[i].quiz_langs; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(quiz.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "quiz_langs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Quiz = n
			}
		}
	}

	if query := qlq.withI18n; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*QuizLangs)
		for i := range nodes {
			if fk := nodes[i].i18n_quiz_langs; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(i18n.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "i18n_quiz_langs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.I18n = n
			}
		}
	}

	return nodes, nil
}

func (qlq *QuizLangsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qlq.querySpec()
	return sqlgraph.CountNodes(ctx, qlq.driver, _spec)
}

func (qlq *QuizLangsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := qlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (qlq *QuizLangsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   quizlangs.Table,
			Columns: quizlangs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: quizlangs.FieldID,
			},
		},
		From:   qlq.sql,
		Unique: true,
	}
	if fields := qlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, quizlangs.FieldID)
		for i := range fields {
			if fields[i] != quizlangs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, quizlangs.ValidColumn)
			}
		}
	}
	return _spec
}

func (qlq *QuizLangsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(qlq.driver.Dialect())
	t1 := builder.Table(quizlangs.Table)
	selector := builder.Select(t1.Columns(quizlangs.Columns...)...).From(t1)
	if qlq.sql != nil {
		selector = qlq.sql
		selector.Select(selector.Columns(quizlangs.Columns...)...)
	}
	for _, p := range qlq.predicates {
		p(selector)
	}
	for _, p := range qlq.order {
		p(selector, quizlangs.ValidColumn)
	}
	if offset := qlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QuizLangsGroupBy is the group-by builder for QuizLangs entities.
type QuizLangsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qlgb *QuizLangsGroupBy) Aggregate(fns ...AggregateFunc) *QuizLangsGroupBy {
	qlgb.fns = append(qlgb.fns, fns...)
	return qlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (qlgb *QuizLangsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := qlgb.path(ctx)
	if err != nil {
		return err
	}
	qlgb.sql = query
	return qlgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := qlgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(qlgb.fields) > 1 {
		return nil, errors.New("ent: QuizLangsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := qlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) StringsX(ctx context.Context) []string {
	v, err := qlgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qlgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) StringX(ctx context.Context) string {
	v, err := qlgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(qlgb.fields) > 1 {
		return nil, errors.New("ent: QuizLangsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := qlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) IntsX(ctx context.Context) []int {
	v, err := qlgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qlgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) IntX(ctx context.Context) int {
	v, err := qlgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(qlgb.fields) > 1 {
		return nil, errors.New("ent: QuizLangsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := qlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := qlgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qlgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := qlgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(qlgb.fields) > 1 {
		return nil, errors.New("ent: QuizLangsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := qlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := qlgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (qlgb *QuizLangsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qlgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qlgb *QuizLangsGroupBy) BoolX(ctx context.Context) bool {
	v, err := qlgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qlgb *QuizLangsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range qlgb.fields {
		if !quizlangs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qlgb *QuizLangsGroupBy) sqlQuery() *sql.Selector {
	selector := qlgb.sql
	columns := make([]string, 0, len(qlgb.fields)+len(qlgb.fns))
	columns = append(columns, qlgb.fields...)
	for _, fn := range qlgb.fns {
		columns = append(columns, fn(selector, quizlangs.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(qlgb.fields...)
}

// QuizLangsSelect is the builder for selecting fields of QuizLangs entities.
type QuizLangsSelect struct {
	*QuizLangsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (qls *QuizLangsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := qls.prepareQuery(ctx); err != nil {
		return err
	}
	qls.sql = qls.QuizLangsQuery.sqlQuery()
	return qls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qls *QuizLangsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := qls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(qls.fields) > 1 {
		return nil, errors.New("ent: QuizLangsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := qls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qls *QuizLangsSelect) StringsX(ctx context.Context) []string {
	v, err := qls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qls *QuizLangsSelect) StringX(ctx context.Context) string {
	v, err := qls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(qls.fields) > 1 {
		return nil, errors.New("ent: QuizLangsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := qls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qls *QuizLangsSelect) IntsX(ctx context.Context) []int {
	v, err := qls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qls *QuizLangsSelect) IntX(ctx context.Context) int {
	v, err := qls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(qls.fields) > 1 {
		return nil, errors.New("ent: QuizLangsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := qls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qls *QuizLangsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := qls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qls *QuizLangsSelect) Float64X(ctx context.Context) float64 {
	v, err := qls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(qls.fields) > 1 {
		return nil, errors.New("ent: QuizLangsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := qls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qls *QuizLangsSelect) BoolsX(ctx context.Context) []bool {
	v, err := qls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (qls *QuizLangsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{quizlangs.Label}
	default:
		err = fmt.Errorf("ent: QuizLangsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qls *QuizLangsSelect) BoolX(ctx context.Context) bool {
	v, err := qls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qls *QuizLangsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qls.sqlQuery().Query()
	if err := qls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qls *QuizLangsSelect) sqlQuery() sql.Querier {
	selector := qls.sql
	selector.Select(selector.Columns(qls.fields...)...)
	return selector
}
