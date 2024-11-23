// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTrack = "Track"
)

// TrackMutation represents an operation that mutates the Track nodes in the graph.
type TrackMutation struct {
	config
	op                   Op
	typ                  string
	id                   *string
	created_at           *time.Time
	updated_at           *time.Time
	name                 *string
	artists              *[]string
	appendartists        []string
	artists_genres       *[]string
	appendartists_genres []string
	album_name           *string
	album_image_uri      *string
	duration_ms          *int
	addduration_ms       *int
	uri                  *string
	clearedFields        map[string]struct{}
	done                 bool
	oldValue             func(context.Context) (*Track, error)
	predicates           []predicate.Track
}

var _ ent.Mutation = (*TrackMutation)(nil)

// trackOption allows management of the mutation configuration using functional options.
type trackOption func(*TrackMutation)

// newTrackMutation creates new mutation for the Track entity.
func newTrackMutation(c config, op Op, opts ...trackOption) *TrackMutation {
	m := &TrackMutation{
		config:        c,
		op:            op,
		typ:           TypeTrack,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTrackID sets the ID field of the mutation.
func withTrackID(id string) trackOption {
	return func(m *TrackMutation) {
		var (
			err   error
			once  sync.Once
			value *Track
		)
		m.oldValue = func(ctx context.Context) (*Track, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Track.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTrack sets the old Track of the mutation.
func withTrack(node *Track) trackOption {
	return func(m *TrackMutation) {
		m.oldValue = func(context.Context) (*Track, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TrackMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TrackMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Track entities.
func (m *TrackMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TrackMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TrackMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Track.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TrackMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TrackMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TrackMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TrackMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TrackMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TrackMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *TrackMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TrackMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TrackMutation) ResetName() {
	m.name = nil
}

// SetArtists sets the "artists" field.
func (m *TrackMutation) SetArtists(s []string) {
	m.artists = &s
	m.appendartists = nil
}

// Artists returns the value of the "artists" field in the mutation.
func (m *TrackMutation) Artists() (r []string, exists bool) {
	v := m.artists
	if v == nil {
		return
	}
	return *v, true
}

// OldArtists returns the old "artists" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldArtists(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArtists is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArtists requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArtists: %w", err)
	}
	return oldValue.Artists, nil
}

// AppendArtists adds s to the "artists" field.
func (m *TrackMutation) AppendArtists(s []string) {
	m.appendartists = append(m.appendartists, s...)
}

// AppendedArtists returns the list of values that were appended to the "artists" field in this mutation.
func (m *TrackMutation) AppendedArtists() ([]string, bool) {
	if len(m.appendartists) == 0 {
		return nil, false
	}
	return m.appendartists, true
}

// ResetArtists resets all changes to the "artists" field.
func (m *TrackMutation) ResetArtists() {
	m.artists = nil
	m.appendartists = nil
}

// SetArtistsGenres sets the "artists_genres" field.
func (m *TrackMutation) SetArtistsGenres(s []string) {
	m.artists_genres = &s
	m.appendartists_genres = nil
}

// ArtistsGenres returns the value of the "artists_genres" field in the mutation.
func (m *TrackMutation) ArtistsGenres() (r []string, exists bool) {
	v := m.artists_genres
	if v == nil {
		return
	}
	return *v, true
}

// OldArtistsGenres returns the old "artists_genres" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldArtistsGenres(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArtistsGenres is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArtistsGenres requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArtistsGenres: %w", err)
	}
	return oldValue.ArtistsGenres, nil
}

// AppendArtistsGenres adds s to the "artists_genres" field.
func (m *TrackMutation) AppendArtistsGenres(s []string) {
	m.appendartists_genres = append(m.appendartists_genres, s...)
}

// AppendedArtistsGenres returns the list of values that were appended to the "artists_genres" field in this mutation.
func (m *TrackMutation) AppendedArtistsGenres() ([]string, bool) {
	if len(m.appendartists_genres) == 0 {
		return nil, false
	}
	return m.appendartists_genres, true
}

// ResetArtistsGenres resets all changes to the "artists_genres" field.
func (m *TrackMutation) ResetArtistsGenres() {
	m.artists_genres = nil
	m.appendartists_genres = nil
}

// SetAlbumName sets the "album_name" field.
func (m *TrackMutation) SetAlbumName(s string) {
	m.album_name = &s
}

// AlbumName returns the value of the "album_name" field in the mutation.
func (m *TrackMutation) AlbumName() (r string, exists bool) {
	v := m.album_name
	if v == nil {
		return
	}
	return *v, true
}

// OldAlbumName returns the old "album_name" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldAlbumName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAlbumName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAlbumName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAlbumName: %w", err)
	}
	return oldValue.AlbumName, nil
}

// ResetAlbumName resets all changes to the "album_name" field.
func (m *TrackMutation) ResetAlbumName() {
	m.album_name = nil
}

// SetAlbumImageURI sets the "album_image_uri" field.
func (m *TrackMutation) SetAlbumImageURI(s string) {
	m.album_image_uri = &s
}

// AlbumImageURI returns the value of the "album_image_uri" field in the mutation.
func (m *TrackMutation) AlbumImageURI() (r string, exists bool) {
	v := m.album_image_uri
	if v == nil {
		return
	}
	return *v, true
}

// OldAlbumImageURI returns the old "album_image_uri" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldAlbumImageURI(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAlbumImageURI is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAlbumImageURI requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAlbumImageURI: %w", err)
	}
	return oldValue.AlbumImageURI, nil
}

// ResetAlbumImageURI resets all changes to the "album_image_uri" field.
func (m *TrackMutation) ResetAlbumImageURI() {
	m.album_image_uri = nil
}

// SetDurationMs sets the "duration_ms" field.
func (m *TrackMutation) SetDurationMs(i int) {
	m.duration_ms = &i
	m.addduration_ms = nil
}

// DurationMs returns the value of the "duration_ms" field in the mutation.
func (m *TrackMutation) DurationMs() (r int, exists bool) {
	v := m.duration_ms
	if v == nil {
		return
	}
	return *v, true
}

// OldDurationMs returns the old "duration_ms" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldDurationMs(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDurationMs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDurationMs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDurationMs: %w", err)
	}
	return oldValue.DurationMs, nil
}

// AddDurationMs adds i to the "duration_ms" field.
func (m *TrackMutation) AddDurationMs(i int) {
	if m.addduration_ms != nil {
		*m.addduration_ms += i
	} else {
		m.addduration_ms = &i
	}
}

// AddedDurationMs returns the value that was added to the "duration_ms" field in this mutation.
func (m *TrackMutation) AddedDurationMs() (r int, exists bool) {
	v := m.addduration_ms
	if v == nil {
		return
	}
	return *v, true
}

// ResetDurationMs resets all changes to the "duration_ms" field.
func (m *TrackMutation) ResetDurationMs() {
	m.duration_ms = nil
	m.addduration_ms = nil
}

// SetURI sets the "uri" field.
func (m *TrackMutation) SetURI(s string) {
	m.uri = &s
}

// URI returns the value of the "uri" field in the mutation.
func (m *TrackMutation) URI() (r string, exists bool) {
	v := m.uri
	if v == nil {
		return
	}
	return *v, true
}

// OldURI returns the old "uri" field's value of the Track entity.
// If the Track object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TrackMutation) OldURI(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURI is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURI requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURI: %w", err)
	}
	return oldValue.URI, nil
}

// ResetURI resets all changes to the "uri" field.
func (m *TrackMutation) ResetURI() {
	m.uri = nil
}

// Where appends a list predicates to the TrackMutation builder.
func (m *TrackMutation) Where(ps ...predicate.Track) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TrackMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TrackMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Track, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TrackMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TrackMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Track).
func (m *TrackMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TrackMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.created_at != nil {
		fields = append(fields, track.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, track.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, track.FieldName)
	}
	if m.artists != nil {
		fields = append(fields, track.FieldArtists)
	}
	if m.artists_genres != nil {
		fields = append(fields, track.FieldArtistsGenres)
	}
	if m.album_name != nil {
		fields = append(fields, track.FieldAlbumName)
	}
	if m.album_image_uri != nil {
		fields = append(fields, track.FieldAlbumImageURI)
	}
	if m.duration_ms != nil {
		fields = append(fields, track.FieldDurationMs)
	}
	if m.uri != nil {
		fields = append(fields, track.FieldURI)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TrackMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case track.FieldCreatedAt:
		return m.CreatedAt()
	case track.FieldUpdatedAt:
		return m.UpdatedAt()
	case track.FieldName:
		return m.Name()
	case track.FieldArtists:
		return m.Artists()
	case track.FieldArtistsGenres:
		return m.ArtistsGenres()
	case track.FieldAlbumName:
		return m.AlbumName()
	case track.FieldAlbumImageURI:
		return m.AlbumImageURI()
	case track.FieldDurationMs:
		return m.DurationMs()
	case track.FieldURI:
		return m.URI()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TrackMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case track.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case track.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case track.FieldName:
		return m.OldName(ctx)
	case track.FieldArtists:
		return m.OldArtists(ctx)
	case track.FieldArtistsGenres:
		return m.OldArtistsGenres(ctx)
	case track.FieldAlbumName:
		return m.OldAlbumName(ctx)
	case track.FieldAlbumImageURI:
		return m.OldAlbumImageURI(ctx)
	case track.FieldDurationMs:
		return m.OldDurationMs(ctx)
	case track.FieldURI:
		return m.OldURI(ctx)
	}
	return nil, fmt.Errorf("unknown Track field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TrackMutation) SetField(name string, value ent.Value) error {
	switch name {
	case track.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case track.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case track.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case track.FieldArtists:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArtists(v)
		return nil
	case track.FieldArtistsGenres:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArtistsGenres(v)
		return nil
	case track.FieldAlbumName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAlbumName(v)
		return nil
	case track.FieldAlbumImageURI:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAlbumImageURI(v)
		return nil
	case track.FieldDurationMs:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDurationMs(v)
		return nil
	case track.FieldURI:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURI(v)
		return nil
	}
	return fmt.Errorf("unknown Track field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TrackMutation) AddedFields() []string {
	var fields []string
	if m.addduration_ms != nil {
		fields = append(fields, track.FieldDurationMs)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TrackMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case track.FieldDurationMs:
		return m.AddedDurationMs()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TrackMutation) AddField(name string, value ent.Value) error {
	switch name {
	case track.FieldDurationMs:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDurationMs(v)
		return nil
	}
	return fmt.Errorf("unknown Track numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TrackMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TrackMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TrackMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Track nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TrackMutation) ResetField(name string) error {
	switch name {
	case track.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case track.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case track.FieldName:
		m.ResetName()
		return nil
	case track.FieldArtists:
		m.ResetArtists()
		return nil
	case track.FieldArtistsGenres:
		m.ResetArtistsGenres()
		return nil
	case track.FieldAlbumName:
		m.ResetAlbumName()
		return nil
	case track.FieldAlbumImageURI:
		m.ResetAlbumImageURI()
		return nil
	case track.FieldDurationMs:
		m.ResetDurationMs()
		return nil
	case track.FieldURI:
		m.ResetURI()
		return nil
	}
	return fmt.Errorf("unknown Track field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TrackMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TrackMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TrackMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TrackMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TrackMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TrackMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TrackMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Track unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TrackMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Track edge %s", name)
}
