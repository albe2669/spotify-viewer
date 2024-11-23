// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/artist"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
	"github.com/rustic-beans/spotify-viewer/ent/schema"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// TrackUpdate is the builder for updating Track entities.
type TrackUpdate struct {
	config
	hooks    []Hook
	mutation *TrackMutation
}

// Where appends a list predicates to the TrackUpdate builder.
func (tu *TrackUpdate) Where(ps ...predicate.Track) *TrackUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TrackUpdate) SetUpdatedAt(t time.Time) *TrackUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetAlbumID sets the "album_id" field.
func (tu *TrackUpdate) SetAlbumID(s string) *TrackUpdate {
	tu.mutation.SetAlbumID(s)
	return tu
}

// SetNillableAlbumID sets the "album_id" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableAlbumID(s *string) *TrackUpdate {
	if s != nil {
		tu.SetAlbumID(*s)
	}
	return tu
}

// SetAvailableMarkets sets the "available_markets" field.
func (tu *TrackUpdate) SetAvailableMarkets(s []string) *TrackUpdate {
	tu.mutation.SetAvailableMarkets(s)
	return tu
}

// AppendAvailableMarkets appends s to the "available_markets" field.
func (tu *TrackUpdate) AppendAvailableMarkets(s []string) *TrackUpdate {
	tu.mutation.AppendAvailableMarkets(s)
	return tu
}

// SetDiscNumber sets the "disc_number" field.
func (tu *TrackUpdate) SetDiscNumber(i int) *TrackUpdate {
	tu.mutation.ResetDiscNumber()
	tu.mutation.SetDiscNumber(i)
	return tu
}

// SetNillableDiscNumber sets the "disc_number" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableDiscNumber(i *int) *TrackUpdate {
	if i != nil {
		tu.SetDiscNumber(*i)
	}
	return tu
}

// AddDiscNumber adds i to the "disc_number" field.
func (tu *TrackUpdate) AddDiscNumber(i int) *TrackUpdate {
	tu.mutation.AddDiscNumber(i)
	return tu
}

// ClearDiscNumber clears the value of the "disc_number" field.
func (tu *TrackUpdate) ClearDiscNumber() *TrackUpdate {
	tu.mutation.ClearDiscNumber()
	return tu
}

// SetDurationMs sets the "duration_ms" field.
func (tu *TrackUpdate) SetDurationMs(i int) *TrackUpdate {
	tu.mutation.ResetDurationMs()
	tu.mutation.SetDurationMs(i)
	return tu
}

// SetNillableDurationMs sets the "duration_ms" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableDurationMs(i *int) *TrackUpdate {
	if i != nil {
		tu.SetDurationMs(*i)
	}
	return tu
}

// AddDurationMs adds i to the "duration_ms" field.
func (tu *TrackUpdate) AddDurationMs(i int) *TrackUpdate {
	tu.mutation.AddDurationMs(i)
	return tu
}

// SetExplicit sets the "explicit" field.
func (tu *TrackUpdate) SetExplicit(b bool) *TrackUpdate {
	tu.mutation.SetExplicit(b)
	return tu
}

// SetNillableExplicit sets the "explicit" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableExplicit(b *bool) *TrackUpdate {
	if b != nil {
		tu.SetExplicit(*b)
	}
	return tu
}

// SetExternalUrls sets the "external_urls" field.
func (tu *TrackUpdate) SetExternalUrls(sm *schema.StringMap) *TrackUpdate {
	tu.mutation.SetExternalUrls(sm)
	return tu
}

// SetHref sets the "href" field.
func (tu *TrackUpdate) SetHref(s string) *TrackUpdate {
	tu.mutation.SetHref(s)
	return tu
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableHref(s *string) *TrackUpdate {
	if s != nil {
		tu.SetHref(*s)
	}
	return tu
}

// SetIsPlayable sets the "is_playable" field.
func (tu *TrackUpdate) SetIsPlayable(b bool) *TrackUpdate {
	tu.mutation.SetIsPlayable(b)
	return tu
}

// SetNillableIsPlayable sets the "is_playable" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableIsPlayable(b *bool) *TrackUpdate {
	if b != nil {
		tu.SetIsPlayable(*b)
	}
	return tu
}

// SetName sets the "name" field.
func (tu *TrackUpdate) SetName(s string) *TrackUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableName(s *string) *TrackUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetPopularity sets the "popularity" field.
func (tu *TrackUpdate) SetPopularity(i int) *TrackUpdate {
	tu.mutation.ResetPopularity()
	tu.mutation.SetPopularity(i)
	return tu
}

// SetNillablePopularity sets the "popularity" field if the given value is not nil.
func (tu *TrackUpdate) SetNillablePopularity(i *int) *TrackUpdate {
	if i != nil {
		tu.SetPopularity(*i)
	}
	return tu
}

// AddPopularity adds i to the "popularity" field.
func (tu *TrackUpdate) AddPopularity(i int) *TrackUpdate {
	tu.mutation.AddPopularity(i)
	return tu
}

// SetPreviewURL sets the "preview_url" field.
func (tu *TrackUpdate) SetPreviewURL(s string) *TrackUpdate {
	tu.mutation.SetPreviewURL(s)
	return tu
}

// SetNillablePreviewURL sets the "preview_url" field if the given value is not nil.
func (tu *TrackUpdate) SetNillablePreviewURL(s *string) *TrackUpdate {
	if s != nil {
		tu.SetPreviewURL(*s)
	}
	return tu
}

// ClearPreviewURL clears the value of the "preview_url" field.
func (tu *TrackUpdate) ClearPreviewURL() *TrackUpdate {
	tu.mutation.ClearPreviewURL()
	return tu
}

// SetTrackNumber sets the "track_number" field.
func (tu *TrackUpdate) SetTrackNumber(i int) *TrackUpdate {
	tu.mutation.ResetTrackNumber()
	tu.mutation.SetTrackNumber(i)
	return tu
}

// SetNillableTrackNumber sets the "track_number" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableTrackNumber(i *int) *TrackUpdate {
	if i != nil {
		tu.SetTrackNumber(*i)
	}
	return tu
}

// AddTrackNumber adds i to the "track_number" field.
func (tu *TrackUpdate) AddTrackNumber(i int) *TrackUpdate {
	tu.mutation.AddTrackNumber(i)
	return tu
}

// SetURI sets the "uri" field.
func (tu *TrackUpdate) SetURI(s string) *TrackUpdate {
	tu.mutation.SetURI(s)
	return tu
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (tu *TrackUpdate) SetNillableURI(s *string) *TrackUpdate {
	if s != nil {
		tu.SetURI(*s)
	}
	return tu
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (tu *TrackUpdate) AddArtistIDs(ids ...string) *TrackUpdate {
	tu.mutation.AddArtistIDs(ids...)
	return tu
}

// AddArtists adds the "artists" edges to the Artist entity.
func (tu *TrackUpdate) AddArtists(a ...*Artist) *TrackUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.AddArtistIDs(ids...)
}

// SetAlbum sets the "album" edge to the Album entity.
func (tu *TrackUpdate) SetAlbum(a *Album) *TrackUpdate {
	return tu.SetAlbumID(a.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tu *TrackUpdate) Mutation() *TrackMutation {
	return tu.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (tu *TrackUpdate) ClearArtists() *TrackUpdate {
	tu.mutation.ClearArtists()
	return tu
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (tu *TrackUpdate) RemoveArtistIDs(ids ...string) *TrackUpdate {
	tu.mutation.RemoveArtistIDs(ids...)
	return tu
}

// RemoveArtists removes "artists" edges to Artist entities.
func (tu *TrackUpdate) RemoveArtists(a ...*Artist) *TrackUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.RemoveArtistIDs(ids...)
}

// ClearAlbum clears the "album" edge to the Album entity.
func (tu *TrackUpdate) ClearAlbum() *TrackUpdate {
	tu.mutation.ClearAlbum()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TrackUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TrackUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TrackUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TrackUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TrackUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := track.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TrackUpdate) check() error {
	if v, ok := tu.mutation.AlbumID(); ok {
		if err := track.AlbumIDValidator(v); err != nil {
			return &ValidationError{Name: "album_id", err: fmt.Errorf(`ent: validator failed for field "Track.album_id": %w`, err)}
		}
	}
	if v, ok := tu.mutation.DurationMs(); ok {
		if err := track.DurationMsValidator(v); err != nil {
			return &ValidationError{Name: "duration_ms", err: fmt.Errorf(`ent: validator failed for field "Track.duration_ms": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Href(); ok {
		if err := track.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Track.href": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := track.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Track.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.URI(); ok {
		if err := track.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Track.uri": %w`, err)}
		}
	}
	if tu.mutation.AlbumCleared() && len(tu.mutation.AlbumIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Track.album"`)
	}
	return nil
}

func (tu *TrackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(track.Table, track.Columns, sqlgraph.NewFieldSpec(track.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(track.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.AvailableMarkets(); ok {
		_spec.SetField(track.FieldAvailableMarkets, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedAvailableMarkets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, track.FieldAvailableMarkets, value)
		})
	}
	if value, ok := tu.mutation.DiscNumber(); ok {
		_spec.SetField(track.FieldDiscNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedDiscNumber(); ok {
		_spec.AddField(track.FieldDiscNumber, field.TypeInt, value)
	}
	if tu.mutation.DiscNumberCleared() {
		_spec.ClearField(track.FieldDiscNumber, field.TypeInt)
	}
	if value, ok := tu.mutation.DurationMs(); ok {
		_spec.SetField(track.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedDurationMs(); ok {
		_spec.AddField(track.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Explicit(); ok {
		_spec.SetField(track.FieldExplicit, field.TypeBool, value)
	}
	if value, ok := tu.mutation.ExternalUrls(); ok {
		_spec.SetField(track.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.Href(); ok {
		_spec.SetField(track.FieldHref, field.TypeString, value)
	}
	if value, ok := tu.mutation.IsPlayable(); ok {
		_spec.SetField(track.FieldIsPlayable, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(track.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Popularity(); ok {
		_spec.SetField(track.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedPopularity(); ok {
		_spec.AddField(track.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := tu.mutation.PreviewURL(); ok {
		_spec.SetField(track.FieldPreviewURL, field.TypeString, value)
	}
	if tu.mutation.PreviewURLCleared() {
		_spec.ClearField(track.FieldPreviewURL, field.TypeString)
	}
	if value, ok := tu.mutation.TrackNumber(); ok {
		_spec.SetField(track.FieldTrackNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedTrackNumber(); ok {
		_spec.AddField(track.FieldTrackNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.URI(); ok {
		_spec.SetField(track.FieldURI, field.TypeString, value)
	}
	if tu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !tu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{track.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TrackUpdateOne is the builder for updating a single Track entity.
type TrackUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrackMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TrackUpdateOne) SetUpdatedAt(t time.Time) *TrackUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetAlbumID sets the "album_id" field.
func (tuo *TrackUpdateOne) SetAlbumID(s string) *TrackUpdateOne {
	tuo.mutation.SetAlbumID(s)
	return tuo
}

// SetNillableAlbumID sets the "album_id" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableAlbumID(s *string) *TrackUpdateOne {
	if s != nil {
		tuo.SetAlbumID(*s)
	}
	return tuo
}

// SetAvailableMarkets sets the "available_markets" field.
func (tuo *TrackUpdateOne) SetAvailableMarkets(s []string) *TrackUpdateOne {
	tuo.mutation.SetAvailableMarkets(s)
	return tuo
}

// AppendAvailableMarkets appends s to the "available_markets" field.
func (tuo *TrackUpdateOne) AppendAvailableMarkets(s []string) *TrackUpdateOne {
	tuo.mutation.AppendAvailableMarkets(s)
	return tuo
}

// SetDiscNumber sets the "disc_number" field.
func (tuo *TrackUpdateOne) SetDiscNumber(i int) *TrackUpdateOne {
	tuo.mutation.ResetDiscNumber()
	tuo.mutation.SetDiscNumber(i)
	return tuo
}

// SetNillableDiscNumber sets the "disc_number" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableDiscNumber(i *int) *TrackUpdateOne {
	if i != nil {
		tuo.SetDiscNumber(*i)
	}
	return tuo
}

// AddDiscNumber adds i to the "disc_number" field.
func (tuo *TrackUpdateOne) AddDiscNumber(i int) *TrackUpdateOne {
	tuo.mutation.AddDiscNumber(i)
	return tuo
}

// ClearDiscNumber clears the value of the "disc_number" field.
func (tuo *TrackUpdateOne) ClearDiscNumber() *TrackUpdateOne {
	tuo.mutation.ClearDiscNumber()
	return tuo
}

// SetDurationMs sets the "duration_ms" field.
func (tuo *TrackUpdateOne) SetDurationMs(i int) *TrackUpdateOne {
	tuo.mutation.ResetDurationMs()
	tuo.mutation.SetDurationMs(i)
	return tuo
}

// SetNillableDurationMs sets the "duration_ms" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableDurationMs(i *int) *TrackUpdateOne {
	if i != nil {
		tuo.SetDurationMs(*i)
	}
	return tuo
}

// AddDurationMs adds i to the "duration_ms" field.
func (tuo *TrackUpdateOne) AddDurationMs(i int) *TrackUpdateOne {
	tuo.mutation.AddDurationMs(i)
	return tuo
}

// SetExplicit sets the "explicit" field.
func (tuo *TrackUpdateOne) SetExplicit(b bool) *TrackUpdateOne {
	tuo.mutation.SetExplicit(b)
	return tuo
}

// SetNillableExplicit sets the "explicit" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableExplicit(b *bool) *TrackUpdateOne {
	if b != nil {
		tuo.SetExplicit(*b)
	}
	return tuo
}

// SetExternalUrls sets the "external_urls" field.
func (tuo *TrackUpdateOne) SetExternalUrls(sm *schema.StringMap) *TrackUpdateOne {
	tuo.mutation.SetExternalUrls(sm)
	return tuo
}

// SetHref sets the "href" field.
func (tuo *TrackUpdateOne) SetHref(s string) *TrackUpdateOne {
	tuo.mutation.SetHref(s)
	return tuo
}

// SetNillableHref sets the "href" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableHref(s *string) *TrackUpdateOne {
	if s != nil {
		tuo.SetHref(*s)
	}
	return tuo
}

// SetIsPlayable sets the "is_playable" field.
func (tuo *TrackUpdateOne) SetIsPlayable(b bool) *TrackUpdateOne {
	tuo.mutation.SetIsPlayable(b)
	return tuo
}

// SetNillableIsPlayable sets the "is_playable" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableIsPlayable(b *bool) *TrackUpdateOne {
	if b != nil {
		tuo.SetIsPlayable(*b)
	}
	return tuo
}

// SetName sets the "name" field.
func (tuo *TrackUpdateOne) SetName(s string) *TrackUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableName(s *string) *TrackUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetPopularity sets the "popularity" field.
func (tuo *TrackUpdateOne) SetPopularity(i int) *TrackUpdateOne {
	tuo.mutation.ResetPopularity()
	tuo.mutation.SetPopularity(i)
	return tuo
}

// SetNillablePopularity sets the "popularity" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillablePopularity(i *int) *TrackUpdateOne {
	if i != nil {
		tuo.SetPopularity(*i)
	}
	return tuo
}

// AddPopularity adds i to the "popularity" field.
func (tuo *TrackUpdateOne) AddPopularity(i int) *TrackUpdateOne {
	tuo.mutation.AddPopularity(i)
	return tuo
}

// SetPreviewURL sets the "preview_url" field.
func (tuo *TrackUpdateOne) SetPreviewURL(s string) *TrackUpdateOne {
	tuo.mutation.SetPreviewURL(s)
	return tuo
}

// SetNillablePreviewURL sets the "preview_url" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillablePreviewURL(s *string) *TrackUpdateOne {
	if s != nil {
		tuo.SetPreviewURL(*s)
	}
	return tuo
}

// ClearPreviewURL clears the value of the "preview_url" field.
func (tuo *TrackUpdateOne) ClearPreviewURL() *TrackUpdateOne {
	tuo.mutation.ClearPreviewURL()
	return tuo
}

// SetTrackNumber sets the "track_number" field.
func (tuo *TrackUpdateOne) SetTrackNumber(i int) *TrackUpdateOne {
	tuo.mutation.ResetTrackNumber()
	tuo.mutation.SetTrackNumber(i)
	return tuo
}

// SetNillableTrackNumber sets the "track_number" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableTrackNumber(i *int) *TrackUpdateOne {
	if i != nil {
		tuo.SetTrackNumber(*i)
	}
	return tuo
}

// AddTrackNumber adds i to the "track_number" field.
func (tuo *TrackUpdateOne) AddTrackNumber(i int) *TrackUpdateOne {
	tuo.mutation.AddTrackNumber(i)
	return tuo
}

// SetURI sets the "uri" field.
func (tuo *TrackUpdateOne) SetURI(s string) *TrackUpdateOne {
	tuo.mutation.SetURI(s)
	return tuo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableURI(s *string) *TrackUpdateOne {
	if s != nil {
		tuo.SetURI(*s)
	}
	return tuo
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (tuo *TrackUpdateOne) AddArtistIDs(ids ...string) *TrackUpdateOne {
	tuo.mutation.AddArtistIDs(ids...)
	return tuo
}

// AddArtists adds the "artists" edges to the Artist entity.
func (tuo *TrackUpdateOne) AddArtists(a ...*Artist) *TrackUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.AddArtistIDs(ids...)
}

// SetAlbum sets the "album" edge to the Album entity.
func (tuo *TrackUpdateOne) SetAlbum(a *Album) *TrackUpdateOne {
	return tuo.SetAlbumID(a.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tuo *TrackUpdateOne) Mutation() *TrackMutation {
	return tuo.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (tuo *TrackUpdateOne) ClearArtists() *TrackUpdateOne {
	tuo.mutation.ClearArtists()
	return tuo
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (tuo *TrackUpdateOne) RemoveArtistIDs(ids ...string) *TrackUpdateOne {
	tuo.mutation.RemoveArtistIDs(ids...)
	return tuo
}

// RemoveArtists removes "artists" edges to Artist entities.
func (tuo *TrackUpdateOne) RemoveArtists(a ...*Artist) *TrackUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.RemoveArtistIDs(ids...)
}

// ClearAlbum clears the "album" edge to the Album entity.
func (tuo *TrackUpdateOne) ClearAlbum() *TrackUpdateOne {
	tuo.mutation.ClearAlbum()
	return tuo
}

// Where appends a list predicates to the TrackUpdate builder.
func (tuo *TrackUpdateOne) Where(ps ...predicate.Track) *TrackUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TrackUpdateOne) Select(field string, fields ...string) *TrackUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Track entity.
func (tuo *TrackUpdateOne) Save(ctx context.Context) (*Track, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TrackUpdateOne) SaveX(ctx context.Context) *Track {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TrackUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TrackUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TrackUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := track.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TrackUpdateOne) check() error {
	if v, ok := tuo.mutation.AlbumID(); ok {
		if err := track.AlbumIDValidator(v); err != nil {
			return &ValidationError{Name: "album_id", err: fmt.Errorf(`ent: validator failed for field "Track.album_id": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.DurationMs(); ok {
		if err := track.DurationMsValidator(v); err != nil {
			return &ValidationError{Name: "duration_ms", err: fmt.Errorf(`ent: validator failed for field "Track.duration_ms": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Href(); ok {
		if err := track.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Track.href": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := track.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Track.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.URI(); ok {
		if err := track.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Track.uri": %w`, err)}
		}
	}
	if tuo.mutation.AlbumCleared() && len(tuo.mutation.AlbumIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Track.album"`)
	}
	return nil
}

func (tuo *TrackUpdateOne) sqlSave(ctx context.Context) (_node *Track, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(track.Table, track.Columns, sqlgraph.NewFieldSpec(track.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Track.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, track.FieldID)
		for _, f := range fields {
			if !track.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != track.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(track.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.AvailableMarkets(); ok {
		_spec.SetField(track.FieldAvailableMarkets, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedAvailableMarkets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, track.FieldAvailableMarkets, value)
		})
	}
	if value, ok := tuo.mutation.DiscNumber(); ok {
		_spec.SetField(track.FieldDiscNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedDiscNumber(); ok {
		_spec.AddField(track.FieldDiscNumber, field.TypeInt, value)
	}
	if tuo.mutation.DiscNumberCleared() {
		_spec.ClearField(track.FieldDiscNumber, field.TypeInt)
	}
	if value, ok := tuo.mutation.DurationMs(); ok {
		_spec.SetField(track.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedDurationMs(); ok {
		_spec.AddField(track.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Explicit(); ok {
		_spec.SetField(track.FieldExplicit, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.ExternalUrls(); ok {
		_spec.SetField(track.FieldExternalUrls, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.Href(); ok {
		_spec.SetField(track.FieldHref, field.TypeString, value)
	}
	if value, ok := tuo.mutation.IsPlayable(); ok {
		_spec.SetField(track.FieldIsPlayable, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(track.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Popularity(); ok {
		_spec.SetField(track.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedPopularity(); ok {
		_spec.AddField(track.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.PreviewURL(); ok {
		_spec.SetField(track.FieldPreviewURL, field.TypeString, value)
	}
	if tuo.mutation.PreviewURLCleared() {
		_spec.ClearField(track.FieldPreviewURL, field.TypeString)
	}
	if value, ok := tuo.mutation.TrackNumber(); ok {
		_spec.SetField(track.FieldTrackNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedTrackNumber(); ok {
		_spec.AddField(track.FieldTrackNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.URI(); ok {
		_spec.SetField(track.FieldURI, field.TypeString, value)
	}
	if tuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !tuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Track{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{track.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
