// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/image"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// AlbumQuery is the builder for querying Album entities.
type AlbumQuery struct {
	config
	ctx             *QueryContext
	order           []album.OrderOption
	inters          []Interceptor
	predicates      []predicate.Album
	withImages      *ImageQuery
	withTracks      *TrackQuery
	loadTotal       []func(context.Context, []*Album) error
	modifiers       []func(*sql.Selector)
	withNamedImages map[string]*ImageQuery
	withNamedTracks map[string]*TrackQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AlbumQuery builder.
func (aq *AlbumQuery) Where(ps ...predicate.Album) *AlbumQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AlbumQuery) Limit(limit int) *AlbumQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AlbumQuery) Offset(offset int) *AlbumQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AlbumQuery) Unique(unique bool) *AlbumQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AlbumQuery) Order(o ...album.OrderOption) *AlbumQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryImages chains the current query on the "images" edge.
func (aq *AlbumQuery) QueryImages() *ImageQuery {
	query := (&ImageClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, album.ImagesTable, album.ImagesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTracks chains the current query on the "tracks" edge.
func (aq *AlbumQuery) QueryTracks() *TrackQuery {
	query := (&TrackClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, selector),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, album.TracksTable, album.TracksColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Album entity from the query.
// Returns a *NotFoundError when no Album was found.
func (aq *AlbumQuery) First(ctx context.Context) (*Album, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{album.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AlbumQuery) FirstX(ctx context.Context) *Album {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Album ID from the query.
// Returns a *NotFoundError when no Album ID was found.
func (aq *AlbumQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{album.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AlbumQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Album entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Album entity is found.
// Returns a *NotFoundError when no Album entities are found.
func (aq *AlbumQuery) Only(ctx context.Context) (*Album, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{album.Label}
	default:
		return nil, &NotSingularError{album.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AlbumQuery) OnlyX(ctx context.Context) *Album {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Album ID in the query.
// Returns a *NotSingularError when more than one Album ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AlbumQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = &NotSingularError{album.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AlbumQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Albums.
func (aq *AlbumQuery) All(ctx context.Context) ([]*Album, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Album, *AlbumQuery]()
	return withInterceptors[[]*Album](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AlbumQuery) AllX(ctx context.Context) []*Album {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Album IDs.
func (aq *AlbumQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(album.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AlbumQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AlbumQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AlbumQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AlbumQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AlbumQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AlbumQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AlbumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AlbumQuery) Clone() *AlbumQuery {
	if aq == nil {
		return nil
	}
	return &AlbumQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]album.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Album{}, aq.predicates...),
		withImages: aq.withImages.Clone(),
		withTracks: aq.withTracks.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithImages tells the query-builder to eager-load the nodes that are connected to
// the "images" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AlbumQuery) WithImages(opts ...func(*ImageQuery)) *AlbumQuery {
	query := (&ImageClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withImages = query
	return aq
}

// WithTracks tells the query-builder to eager-load the nodes that are connected to
// the "tracks" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AlbumQuery) WithTracks(opts ...func(*TrackQuery)) *AlbumQuery {
	query := (&TrackClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withTracks = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AlbumType album.AlbumType `json:"album_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Album.Query().
//		GroupBy(album.FieldAlbumType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AlbumQuery) GroupBy(field string, fields ...string) *AlbumGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AlbumGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = album.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AlbumType album.AlbumType `json:"album_type,omitempty"`
//	}
//
//	client.Album.Query().
//		Select(album.FieldAlbumType).
//		Scan(ctx, &v)
func (aq *AlbumQuery) Select(fields ...string) *AlbumSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AlbumSelect{AlbumQuery: aq}
	sbuild.label = album.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AlbumSelect configured with the given aggregations.
func (aq *AlbumQuery) Aggregate(fns ...AggregateFunc) *AlbumSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AlbumQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !album.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AlbumQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Album, error) {
	var (
		nodes       = []*Album{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withImages != nil,
			aq.withTracks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Album).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Album{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withImages; query != nil {
		if err := aq.loadImages(ctx, query, nodes,
			func(n *Album) { n.Edges.Images = []*Image{} },
			func(n *Album, e *Image) { n.Edges.Images = append(n.Edges.Images, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTracks; query != nil {
		if err := aq.loadTracks(ctx, query, nodes,
			func(n *Album) { n.Edges.Tracks = []*Track{} },
			func(n *Album, e *Track) { n.Edges.Tracks = append(n.Edges.Tracks, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedImages {
		if err := aq.loadImages(ctx, query, nodes,
			func(n *Album) { n.appendNamedImages(name) },
			func(n *Album, e *Image) { n.appendNamedImages(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedTracks {
		if err := aq.loadTracks(ctx, query, nodes,
			func(n *Album) { n.appendNamedTracks(name) },
			func(n *Album, e *Track) { n.appendNamedTracks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AlbumQuery) loadImages(ctx context.Context, query *ImageQuery, nodes []*Album, init func(*Album), assign func(*Album, *Image)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Album)
	nids := make(map[string]map[*Album]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(album.ImagesTable)
		s.Join(joinT).On(s.C(image.FieldID), joinT.C(album.ImagesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(album.ImagesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(album.ImagesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Album]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Image](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "images" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AlbumQuery) loadTracks(ctx context.Context, query *TrackQuery, nodes []*Album, init func(*Album), assign func(*Album, *Track)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Album)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Track(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(album.TracksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.album_tracks
		if fk == nil {
			return fmt.Errorf(`foreign-key "album_tracks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "album_tracks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AlbumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AlbumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(album.Table, album.Columns, sqlgraph.NewFieldSpec(album.FieldID, field.TypeString))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, album.FieldID)
		for i := range fields {
			if fields[i] != album.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AlbumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(album.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = album.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aq *AlbumQuery) ForUpdate(opts ...sql.LockOption) *AlbumQuery {
	if aq.driver.Dialect() == dialect.Postgres {
		aq.Unique(false)
	}
	aq.modifiers = append(aq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aq *AlbumQuery) ForShare(opts ...sql.LockOption) *AlbumQuery {
	if aq.driver.Dialect() == dialect.Postgres {
		aq.Unique(false)
	}
	aq.modifiers = append(aq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aq
}

// WithNamedImages tells the query-builder to eager-load the nodes that are connected to the "images"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AlbumQuery) WithNamedImages(name string, opts ...func(*ImageQuery)) *AlbumQuery {
	query := (&ImageClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedImages == nil {
		aq.withNamedImages = make(map[string]*ImageQuery)
	}
	aq.withNamedImages[name] = query
	return aq
}

// WithNamedTracks tells the query-builder to eager-load the nodes that are connected to the "tracks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AlbumQuery) WithNamedTracks(name string, opts ...func(*TrackQuery)) *AlbumQuery {
	query := (&TrackClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedTracks == nil {
		aq.withNamedTracks = make(map[string]*TrackQuery)
	}
	aq.withNamedTracks[name] = query
	return aq
}

// AlbumGroupBy is the group-by builder for Album entities.
type AlbumGroupBy struct {
	selector
	build *AlbumQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AlbumGroupBy) Aggregate(fns ...AggregateFunc) *AlbumGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AlbumGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AlbumQuery, *AlbumGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AlbumGroupBy) sqlScan(ctx context.Context, root *AlbumQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AlbumSelect is the builder for selecting fields of Album entities.
type AlbumSelect struct {
	*AlbumQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AlbumSelect) Aggregate(fns ...AggregateFunc) *AlbumSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AlbumSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AlbumQuery, *AlbumSelect](ctx, as.AlbumQuery, as, as.inters, v)
}

func (as *AlbumSelect) sqlScan(ctx context.Context, root *AlbumQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
