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
	"github.com/rustic-beans/spotify-viewer/ent/artist"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// ArtistQuery is the builder for querying Artist entities.
type ArtistQuery struct {
	config
	ctx             *QueryContext
	order           []artist.OrderOption
	inters          []Interceptor
	predicates      []predicate.Artist
	withAlbums      *AlbumQuery
	withTracks      *TrackQuery
	loadTotal       []func(context.Context, []*Artist) error
	modifiers       []func(*sql.Selector)
	withNamedAlbums map[string]*AlbumQuery
	withNamedTracks map[string]*TrackQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArtistQuery builder.
func (aq *ArtistQuery) Where(ps ...predicate.Artist) *ArtistQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ArtistQuery) Limit(limit int) *ArtistQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ArtistQuery) Offset(offset int) *ArtistQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArtistQuery) Unique(unique bool) *ArtistQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ArtistQuery) Order(o ...artist.OrderOption) *ArtistQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAlbums chains the current query on the "albums" edge.
func (aq *ArtistQuery) QueryAlbums() *AlbumQuery {
	query := (&AlbumClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, selector),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, artist.AlbumsTable, artist.AlbumsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTracks chains the current query on the "tracks" edge.
func (aq *ArtistQuery) QueryTracks() *TrackQuery {
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
			sqlgraph.From(artist.Table, artist.FieldID, selector),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, artist.TracksTable, artist.TracksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Artist entity from the query.
// Returns a *NotFoundError when no Artist was found.
func (aq *ArtistQuery) First(ctx context.Context) (*Artist, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{artist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArtistQuery) FirstX(ctx context.Context) *Artist {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Artist ID from the query.
// Returns a *NotFoundError when no Artist ID was found.
func (aq *ArtistQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{artist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArtistQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Artist entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Artist entity is found.
// Returns a *NotFoundError when no Artist entities are found.
func (aq *ArtistQuery) Only(ctx context.Context) (*Artist, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{artist.Label}
	default:
		return nil, &NotSingularError{artist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArtistQuery) OnlyX(ctx context.Context) *Artist {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Artist ID in the query.
// Returns a *NotSingularError when more than one Artist ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArtistQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{artist.Label}
	default:
		err = &NotSingularError{artist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArtistQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Artists.
func (aq *ArtistQuery) All(ctx context.Context) ([]*Artist, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Artist, *ArtistQuery]()
	return withInterceptors[[]*Artist](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArtistQuery) AllX(ctx context.Context) []*Artist {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Artist IDs.
func (aq *ArtistQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(artist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArtistQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArtistQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ArtistQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArtistQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArtistQuery) Exist(ctx context.Context) (bool, error) {
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
func (aq *ArtistQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArtistQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArtistQuery) Clone() *ArtistQuery {
	if aq == nil {
		return nil
	}
	return &ArtistQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]artist.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Artist{}, aq.predicates...),
		withAlbums: aq.withAlbums.Clone(),
		withTracks: aq.withTracks.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithAlbums tells the query-builder to eager-load the nodes that are connected to
// the "albums" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtistQuery) WithAlbums(opts ...func(*AlbumQuery)) *ArtistQuery {
	query := (&AlbumClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAlbums = query
	return aq
}

// WithTracks tells the query-builder to eager-load the nodes that are connected to
// the "tracks" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtistQuery) WithTracks(opts ...func(*TrackQuery)) *ArtistQuery {
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
//		ExternalUrls *schema.StringMap `json:"external_urls,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Artist.Query().
//		GroupBy(artist.FieldExternalUrls).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ArtistQuery) GroupBy(field string, fields ...string) *ArtistGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArtistGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = artist.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExternalUrls *schema.StringMap `json:"external_urls,omitempty"`
//	}
//
//	client.Artist.Query().
//		Select(artist.FieldExternalUrls).
//		Scan(ctx, &v)
func (aq *ArtistQuery) Select(fields ...string) *ArtistSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ArtistSelect{ArtistQuery: aq}
	sbuild.label = artist.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArtistSelect configured with the given aggregations.
func (aq *ArtistQuery) Aggregate(fns ...AggregateFunc) *ArtistSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ArtistQuery) prepareQuery(ctx context.Context) error {
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
		if !artist.ValidColumn(f) {
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

func (aq *ArtistQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Artist, error) {
	var (
		nodes       = []*Artist{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withAlbums != nil,
			aq.withTracks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Artist).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Artist{config: aq.config}
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
	if query := aq.withAlbums; query != nil {
		if err := aq.loadAlbums(ctx, query, nodes,
			func(n *Artist) { n.Edges.Albums = []*Album{} },
			func(n *Artist, e *Album) { n.Edges.Albums = append(n.Edges.Albums, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTracks; query != nil {
		if err := aq.loadTracks(ctx, query, nodes,
			func(n *Artist) { n.Edges.Tracks = []*Track{} },
			func(n *Artist, e *Track) { n.Edges.Tracks = append(n.Edges.Tracks, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedAlbums {
		if err := aq.loadAlbums(ctx, query, nodes,
			func(n *Artist) { n.appendNamedAlbums(name) },
			func(n *Artist, e *Album) { n.appendNamedAlbums(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedTracks {
		if err := aq.loadTracks(ctx, query, nodes,
			func(n *Artist) { n.appendNamedTracks(name) },
			func(n *Artist, e *Track) { n.appendNamedTracks(name, e) }); err != nil {
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

func (aq *ArtistQuery) loadAlbums(ctx context.Context, query *AlbumQuery, nodes []*Artist, init func(*Artist), assign func(*Artist, *Album)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Artist)
	nids := make(map[string]map[*Artist]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(artist.AlbumsTable)
		s.Join(joinT).On(s.C(album.FieldID), joinT.C(artist.AlbumsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(artist.AlbumsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(artist.AlbumsPrimaryKey[0]))
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
					nids[inValue] = map[*Artist]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Album](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "albums" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *ArtistQuery) loadTracks(ctx context.Context, query *TrackQuery, nodes []*Artist, init func(*Artist), assign func(*Artist, *Track)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Artist)
	nids := make(map[string]map[*Artist]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(artist.TracksTable)
		s.Join(joinT).On(s.C(track.FieldID), joinT.C(artist.TracksPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(artist.TracksPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(artist.TracksPrimaryKey[0]))
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
					nids[inValue] = map[*Artist]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Track](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tracks" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *ArtistQuery) sqlCount(ctx context.Context) (int, error) {
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

func (aq *ArtistQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for i := range fields {
			if fields[i] != artist.FieldID {
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

func (aq *ArtistQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(artist.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = artist.Columns
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
func (aq *ArtistQuery) ForUpdate(opts ...sql.LockOption) *ArtistQuery {
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
func (aq *ArtistQuery) ForShare(opts ...sql.LockOption) *ArtistQuery {
	if aq.driver.Dialect() == dialect.Postgres {
		aq.Unique(false)
	}
	aq.modifiers = append(aq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aq
}

// WithNamedAlbums tells the query-builder to eager-load the nodes that are connected to the "albums"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtistQuery) WithNamedAlbums(name string, opts ...func(*AlbumQuery)) *ArtistQuery {
	query := (&AlbumClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedAlbums == nil {
		aq.withNamedAlbums = make(map[string]*AlbumQuery)
	}
	aq.withNamedAlbums[name] = query
	return aq
}

// WithNamedTracks tells the query-builder to eager-load the nodes that are connected to the "tracks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtistQuery) WithNamedTracks(name string, opts ...func(*TrackQuery)) *ArtistQuery {
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

// ArtistGroupBy is the group-by builder for Artist entities.
type ArtistGroupBy struct {
	selector
	build *ArtistQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArtistGroupBy) Aggregate(fns ...AggregateFunc) *ArtistGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ArtistGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtistQuery, *ArtistGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ArtistGroupBy) sqlScan(ctx context.Context, root *ArtistQuery, v any) error {
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

// ArtistSelect is the builder for selecting fields of Artist entities.
type ArtistSelect struct {
	*ArtistQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ArtistSelect) Aggregate(fns ...AggregateFunc) *ArtistSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArtistSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtistQuery, *ArtistSelect](ctx, as.ArtistQuery, as, as.inters, v)
}

func (as *ArtistSelect) sqlScan(ctx context.Context, root *ArtistQuery, v any) error {
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
