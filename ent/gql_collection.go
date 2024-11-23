// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TrackQuery) CollectFields(ctx context.Context, satisfies ...string) (*TrackQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TrackQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(track.Columns))
		selectedFields = []string{track.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[track.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, track.FieldCreatedAt)
				fieldSeen[track.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[track.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, track.FieldUpdatedAt)
				fieldSeen[track.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[track.FieldName]; !ok {
				selectedFields = append(selectedFields, track.FieldName)
				fieldSeen[track.FieldName] = struct{}{}
			}
		case "artists":
			if _, ok := fieldSeen[track.FieldArtists]; !ok {
				selectedFields = append(selectedFields, track.FieldArtists)
				fieldSeen[track.FieldArtists] = struct{}{}
			}
		case "artistsGenres":
			if _, ok := fieldSeen[track.FieldArtistsGenres]; !ok {
				selectedFields = append(selectedFields, track.FieldArtistsGenres)
				fieldSeen[track.FieldArtistsGenres] = struct{}{}
			}
		case "albumName":
			if _, ok := fieldSeen[track.FieldAlbumName]; !ok {
				selectedFields = append(selectedFields, track.FieldAlbumName)
				fieldSeen[track.FieldAlbumName] = struct{}{}
			}
		case "albumImageURI":
			if _, ok := fieldSeen[track.FieldAlbumImageURI]; !ok {
				selectedFields = append(selectedFields, track.FieldAlbumImageURI)
				fieldSeen[track.FieldAlbumImageURI] = struct{}{}
			}
		case "durationMs":
			if _, ok := fieldSeen[track.FieldDurationMs]; !ok {
				selectedFields = append(selectedFields, track.FieldDurationMs)
				fieldSeen[track.FieldDurationMs] = struct{}{}
			}
		case "uri":
			if _, ok := fieldSeen[track.FieldURI]; !ok {
				selectedFields = append(selectedFields, track.FieldURI)
				fieldSeen[track.FieldURI] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type trackPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TrackPaginateOption
}

func newTrackPaginateArgs(rv map[string]any) *trackPaginateArgs {
	args := &trackPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
