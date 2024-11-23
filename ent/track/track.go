// Code generated by ent, DO NOT EDIT.

package track

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the track type in the database.
	Label = "track"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAlbumID holds the string denoting the album_id field in the database.
	FieldAlbumID = "album_id"
	// FieldAvailableMarkets holds the string denoting the available_markets field in the database.
	FieldAvailableMarkets = "available_markets"
	// FieldDiscNumber holds the string denoting the disc_number field in the database.
	FieldDiscNumber = "disc_number"
	// FieldDurationMs holds the string denoting the duration_ms field in the database.
	FieldDurationMs = "duration_ms"
	// FieldExplicit holds the string denoting the explicit field in the database.
	FieldExplicit = "explicit"
	// FieldExternalUrls holds the string denoting the external_urls field in the database.
	FieldExternalUrls = "external_urls"
	// FieldHref holds the string denoting the href field in the database.
	FieldHref = "href"
	// FieldIsPlayable holds the string denoting the is_playable field in the database.
	FieldIsPlayable = "is_playable"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPopularity holds the string denoting the popularity field in the database.
	FieldPopularity = "popularity"
	// FieldPreviewURL holds the string denoting the preview_url field in the database.
	FieldPreviewURL = "preview_url"
	// FieldTrackNumber holds the string denoting the track_number field in the database.
	FieldTrackNumber = "track_number"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// EdgeArtists holds the string denoting the artists edge name in mutations.
	EdgeArtists = "artists"
	// EdgeAlbum holds the string denoting the album edge name in mutations.
	EdgeAlbum = "album"
	// Table holds the table name of the track in the database.
	Table = "tracks"
	// ArtistsTable is the table that holds the artists relation/edge. The primary key declared below.
	ArtistsTable = "artist_tracks"
	// ArtistsInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	ArtistsInverseTable = "artists"
	// AlbumTable is the table that holds the album relation/edge.
	AlbumTable = "tracks"
	// AlbumInverseTable is the table name for the Album entity.
	// It exists in this package in order to avoid circular dependency with the "album" package.
	AlbumInverseTable = "albums"
	// AlbumColumn is the table column denoting the album relation/edge.
	AlbumColumn = "album_id"
)

// Columns holds all SQL columns for track fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAlbumID,
	FieldAvailableMarkets,
	FieldDiscNumber,
	FieldDurationMs,
	FieldExplicit,
	FieldExternalUrls,
	FieldHref,
	FieldIsPlayable,
	FieldName,
	FieldPopularity,
	FieldPreviewURL,
	FieldTrackNumber,
	FieldURI,
}

var (
	// ArtistsPrimaryKey and ArtistsColumn2 are the table columns denoting the
	// primary key for the artists relation (M2M).
	ArtistsPrimaryKey = []string{"artist_id", "track_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// AlbumIDValidator is a validator for the "album_id" field. It is called by the builders before save.
	AlbumIDValidator func(string) error
	// DurationMsValidator is a validator for the "duration_ms" field. It is called by the builders before save.
	DurationMsValidator func(int) error
	// DefaultExplicit holds the default value on creation for the "explicit" field.
	DefaultExplicit bool
	// HrefValidator is a validator for the "href" field. It is called by the builders before save.
	HrefValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// URIValidator is a validator for the "uri" field. It is called by the builders before save.
	URIValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Track queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAlbumID orders the results by the album_id field.
func ByAlbumID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlbumID, opts...).ToFunc()
}

// ByDiscNumber orders the results by the disc_number field.
func ByDiscNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscNumber, opts...).ToFunc()
}

// ByDurationMs orders the results by the duration_ms field.
func ByDurationMs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDurationMs, opts...).ToFunc()
}

// ByExplicit orders the results by the explicit field.
func ByExplicit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExplicit, opts...).ToFunc()
}

// ByHref orders the results by the href field.
func ByHref(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHref, opts...).ToFunc()
}

// ByIsPlayable orders the results by the is_playable field.
func ByIsPlayable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPlayable, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPopularity orders the results by the popularity field.
func ByPopularity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPopularity, opts...).ToFunc()
}

// ByPreviewURL orders the results by the preview_url field.
func ByPreviewURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPreviewURL, opts...).ToFunc()
}

// ByTrackNumber orders the results by the track_number field.
func ByTrackNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrackNumber, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// ByArtistsCount orders the results by artists count.
func ByArtistsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArtistsStep(), opts...)
	}
}

// ByArtists orders the results by artists terms.
func ByArtists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtistsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAlbumField orders the results by album field.
func ByAlbumField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAlbumStep(), sql.OrderByField(field, opts...))
	}
}
func newArtistsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtistsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ArtistsTable, ArtistsPrimaryKey...),
	)
}
func newAlbumStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AlbumInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AlbumTable, AlbumColumn),
	)
}
