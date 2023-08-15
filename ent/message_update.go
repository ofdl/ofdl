// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ofdl/ofdl/ent/message"
	"github.com/ofdl/ofdl/ent/messagemedia"
	"github.com/ofdl/ofdl/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetSubscriptionID sets the "subscription_id" field.
func (mu *MessageUpdate) SetSubscriptionID(i int) *MessageUpdate {
	mu.mutation.ResetSubscriptionID()
	mu.mutation.SetSubscriptionID(i)
	return mu
}

// AddSubscriptionID adds i to the "subscription_id" field.
func (mu *MessageUpdate) AddSubscriptionID(i int) *MessageUpdate {
	mu.mutation.AddSubscriptionID(i)
	return mu
}

// SetText sets the "text" field.
func (mu *MessageUpdate) SetText(s string) *MessageUpdate {
	mu.mutation.SetText(s)
	return mu
}

// SetPostedAt sets the "posted_at" field.
func (mu *MessageUpdate) SetPostedAt(s string) *MessageUpdate {
	mu.mutation.SetPostedAt(s)
	return mu
}

// AddMessageMediumIDs adds the "message_media" edge to the MessageMedia entity by IDs.
func (mu *MessageUpdate) AddMessageMediumIDs(ids ...int) *MessageUpdate {
	mu.mutation.AddMessageMediumIDs(ids...)
	return mu
}

// AddMessageMedia adds the "message_media" edges to the MessageMedia entity.
func (mu *MessageUpdate) AddMessageMedia(m ...*MessageMedia) *MessageUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMessageMediumIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearMessageMedia clears all "message_media" edges to the MessageMedia entity.
func (mu *MessageUpdate) ClearMessageMedia() *MessageUpdate {
	mu.mutation.ClearMessageMedia()
	return mu
}

// RemoveMessageMediumIDs removes the "message_media" edge to MessageMedia entities by IDs.
func (mu *MessageUpdate) RemoveMessageMediumIDs(ids ...int) *MessageUpdate {
	mu.mutation.RemoveMessageMediumIDs(ids...)
	return mu
}

// RemoveMessageMedia removes "message_media" edges to MessageMedia entities.
func (mu *MessageUpdate) RemoveMessageMedia(m ...*MessageMedia) *MessageUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMessageMediumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.SubscriptionID(); ok {
		_spec.SetField(message.FieldSubscriptionID, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedSubscriptionID(); ok {
		_spec.AddField(message.FieldSubscriptionID, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Text(); ok {
		_spec.SetField(message.FieldText, field.TypeString, value)
	}
	if value, ok := mu.mutation.PostedAt(); ok {
		_spec.SetField(message.FieldPostedAt, field.TypeString, value)
	}
	if mu.mutation.MessageMediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMessageMediaIDs(); len(nodes) > 0 && !mu.mutation.MessageMediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MessageMediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetSubscriptionID sets the "subscription_id" field.
func (muo *MessageUpdateOne) SetSubscriptionID(i int) *MessageUpdateOne {
	muo.mutation.ResetSubscriptionID()
	muo.mutation.SetSubscriptionID(i)
	return muo
}

// AddSubscriptionID adds i to the "subscription_id" field.
func (muo *MessageUpdateOne) AddSubscriptionID(i int) *MessageUpdateOne {
	muo.mutation.AddSubscriptionID(i)
	return muo
}

// SetText sets the "text" field.
func (muo *MessageUpdateOne) SetText(s string) *MessageUpdateOne {
	muo.mutation.SetText(s)
	return muo
}

// SetPostedAt sets the "posted_at" field.
func (muo *MessageUpdateOne) SetPostedAt(s string) *MessageUpdateOne {
	muo.mutation.SetPostedAt(s)
	return muo
}

// AddMessageMediumIDs adds the "message_media" edge to the MessageMedia entity by IDs.
func (muo *MessageUpdateOne) AddMessageMediumIDs(ids ...int) *MessageUpdateOne {
	muo.mutation.AddMessageMediumIDs(ids...)
	return muo
}

// AddMessageMedia adds the "message_media" edges to the MessageMedia entity.
func (muo *MessageUpdateOne) AddMessageMedia(m ...*MessageMedia) *MessageUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMessageMediumIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearMessageMedia clears all "message_media" edges to the MessageMedia entity.
func (muo *MessageUpdateOne) ClearMessageMedia() *MessageUpdateOne {
	muo.mutation.ClearMessageMedia()
	return muo
}

// RemoveMessageMediumIDs removes the "message_media" edge to MessageMedia entities by IDs.
func (muo *MessageUpdateOne) RemoveMessageMediumIDs(ids ...int) *MessageUpdateOne {
	muo.mutation.RemoveMessageMediumIDs(ids...)
	return muo
}

// RemoveMessageMedia removes "message_media" edges to MessageMedia entities.
func (muo *MessageUpdateOne) RemoveMessageMedia(m ...*MessageMedia) *MessageUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMessageMediumIDs(ids...)
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.SubscriptionID(); ok {
		_spec.SetField(message.FieldSubscriptionID, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedSubscriptionID(); ok {
		_spec.AddField(message.FieldSubscriptionID, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Text(); ok {
		_spec.SetField(message.FieldText, field.TypeString, value)
	}
	if value, ok := muo.mutation.PostedAt(); ok {
		_spec.SetField(message.FieldPostedAt, field.TypeString, value)
	}
	if muo.mutation.MessageMediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMessageMediaIDs(); len(nodes) > 0 && !muo.mutation.MessageMediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MessageMediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.MessageMediaTable,
			Columns: []string{message.MessageMediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}