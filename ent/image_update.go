// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/image"
	"github.com/rustic-beans/spotify-viewer/ent/predicate"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetURL sets the "url" field.
func (iu *ImageUpdate) SetURL(s string) *ImageUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableURL(s *string) *ImageUpdate {
	if s != nil {
		iu.SetURL(*s)
	}
	return iu
}

// SetWidth sets the "width" field.
func (iu *ImageUpdate) SetWidth(i int) *ImageUpdate {
	iu.mutation.ResetWidth()
	iu.mutation.SetWidth(i)
	return iu
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableWidth(i *int) *ImageUpdate {
	if i != nil {
		iu.SetWidth(*i)
	}
	return iu
}

// AddWidth adds i to the "width" field.
func (iu *ImageUpdate) AddWidth(i int) *ImageUpdate {
	iu.mutation.AddWidth(i)
	return iu
}

// SetHeight sets the "height" field.
func (iu *ImageUpdate) SetHeight(i int) *ImageUpdate {
	iu.mutation.ResetHeight()
	iu.mutation.SetHeight(i)
	return iu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableHeight(i *int) *ImageUpdate {
	if i != nil {
		iu.SetHeight(*i)
	}
	return iu
}

// AddHeight adds i to the "height" field.
func (iu *ImageUpdate) AddHeight(i int) *ImageUpdate {
	iu.mutation.AddHeight(i)
	return iu
}

// SetText sets the "text" field.
func (iu *ImageUpdate) SetText(s string) *ImageUpdate {
	iu.mutation.SetText(s)
	return iu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableText(s *string) *ImageUpdate {
	if s != nil {
		iu.SetText(*s)
	}
	return iu
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (iu *ImageUpdate) AddAlbumIDs(ids ...string) *ImageUpdate {
	iu.mutation.AddAlbumIDs(ids...)
	return iu
}

// AddAlbums adds the "albums" edges to the Album entity.
func (iu *ImageUpdate) AddAlbums(a ...*Album) *ImageUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.AddAlbumIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (iu *ImageUpdate) ClearAlbums() *ImageUpdate {
	iu.mutation.ClearAlbums()
	return iu
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (iu *ImageUpdate) RemoveAlbumIDs(ids ...string) *ImageUpdate {
	iu.mutation.RemoveAlbumIDs(ids...)
	return iu
}

// RemoveAlbums removes "albums" edges to Album entities.
func (iu *ImageUpdate) RemoveAlbums(a ...*Album) *ImageUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.RemoveAlbumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if v, ok := iu.mutation.URL(); ok {
		if err := image.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Image.url": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Width(); ok {
		if err := image.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "Image.width": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Height(); ok {
		if err := image.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "Image.height": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Text(); ok {
		if err := image.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Image.text": %w`, err)}
		}
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeString))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
	}
	if value, ok := iu.mutation.Width(); ok {
		_spec.SetField(image.FieldWidth, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedWidth(); ok {
		_spec.AddField(image.FieldWidth, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Height(); ok {
		_spec.SetField(image.FieldHeight, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedHeight(); ok {
		_spec.AddField(image.FieldHeight, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Text(); ok {
		_spec.SetField(image.FieldText, field.TypeString, value)
	}
	if iu.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !iu.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetURL sets the "url" field.
func (iuo *ImageUpdateOne) SetURL(s string) *ImageUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableURL(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetURL(*s)
	}
	return iuo
}

// SetWidth sets the "width" field.
func (iuo *ImageUpdateOne) SetWidth(i int) *ImageUpdateOne {
	iuo.mutation.ResetWidth()
	iuo.mutation.SetWidth(i)
	return iuo
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableWidth(i *int) *ImageUpdateOne {
	if i != nil {
		iuo.SetWidth(*i)
	}
	return iuo
}

// AddWidth adds i to the "width" field.
func (iuo *ImageUpdateOne) AddWidth(i int) *ImageUpdateOne {
	iuo.mutation.AddWidth(i)
	return iuo
}

// SetHeight sets the "height" field.
func (iuo *ImageUpdateOne) SetHeight(i int) *ImageUpdateOne {
	iuo.mutation.ResetHeight()
	iuo.mutation.SetHeight(i)
	return iuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableHeight(i *int) *ImageUpdateOne {
	if i != nil {
		iuo.SetHeight(*i)
	}
	return iuo
}

// AddHeight adds i to the "height" field.
func (iuo *ImageUpdateOne) AddHeight(i int) *ImageUpdateOne {
	iuo.mutation.AddHeight(i)
	return iuo
}

// SetText sets the "text" field.
func (iuo *ImageUpdateOne) SetText(s string) *ImageUpdateOne {
	iuo.mutation.SetText(s)
	return iuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableText(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetText(*s)
	}
	return iuo
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (iuo *ImageUpdateOne) AddAlbumIDs(ids ...string) *ImageUpdateOne {
	iuo.mutation.AddAlbumIDs(ids...)
	return iuo
}

// AddAlbums adds the "albums" edges to the Album entity.
func (iuo *ImageUpdateOne) AddAlbums(a ...*Album) *ImageUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.AddAlbumIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (iuo *ImageUpdateOne) ClearAlbums() *ImageUpdateOne {
	iuo.mutation.ClearAlbums()
	return iuo
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (iuo *ImageUpdateOne) RemoveAlbumIDs(ids ...string) *ImageUpdateOne {
	iuo.mutation.RemoveAlbumIDs(ids...)
	return iuo
}

// RemoveAlbums removes "albums" edges to Album entities.
func (iuo *ImageUpdateOne) RemoveAlbums(a ...*Album) *ImageUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.RemoveAlbumIDs(ids...)
}

// Where appends a list predicates to the ImageUpdate builder.
func (iuo *ImageUpdateOne) Where(ps ...predicate.Image) *ImageUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if v, ok := iuo.mutation.URL(); ok {
		if err := image.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Image.url": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Width(); ok {
		if err := image.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "Image.width": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Height(); ok {
		if err := image.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "Image.height": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Text(); ok {
		if err := image.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Image.text": %w`, err)}
		}
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeString))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Width(); ok {
		_spec.SetField(image.FieldWidth, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedWidth(); ok {
		_spec.AddField(image.FieldWidth, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Height(); ok {
		_spec.SetField(image.FieldHeight, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedHeight(); ok {
		_spec.AddField(image.FieldHeight, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Text(); ok {
		_spec.SetField(image.FieldText, field.TypeString, value)
	}
	if iuo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !iuo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.AlbumsTable,
			Columns: image.AlbumsPrimaryKey,
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
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
