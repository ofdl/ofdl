// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ofdl/ofdl/ent/messagemedia"
)

// MessageMediaCreate is the builder for creating a MessageMedia entity.
type MessageMediaCreate struct {
	config
	mutation *MessageMediaMutation
	hooks    []Hook
}

// SetMessageID sets the "message_id" field.
func (mmc *MessageMediaCreate) SetMessageID(i int) *MessageMediaCreate {
	mmc.mutation.SetMessageID(i)
	return mmc
}

// SetType sets the "type" field.
func (mmc *MessageMediaCreate) SetType(s string) *MessageMediaCreate {
	mmc.mutation.SetType(s)
	return mmc
}

// SetFull sets the "full" field.
func (mmc *MessageMediaCreate) SetFull(s string) *MessageMediaCreate {
	mmc.mutation.SetFull(s)
	return mmc
}

// SetDownloadedAt sets the "downloaded_at" field.
func (mmc *MessageMediaCreate) SetDownloadedAt(t time.Time) *MessageMediaCreate {
	mmc.mutation.SetDownloadedAt(t)
	return mmc
}

// SetNillableDownloadedAt sets the "downloaded_at" field if the given value is not nil.
func (mmc *MessageMediaCreate) SetNillableDownloadedAt(t *time.Time) *MessageMediaCreate {
	if t != nil {
		mmc.SetDownloadedAt(*t)
	}
	return mmc
}

// SetStashID sets the "stash_id" field.
func (mmc *MessageMediaCreate) SetStashID(s string) *MessageMediaCreate {
	mmc.mutation.SetStashID(s)
	return mmc
}

// SetOrganizedAt sets the "organized_at" field.
func (mmc *MessageMediaCreate) SetOrganizedAt(t time.Time) *MessageMediaCreate {
	mmc.mutation.SetOrganizedAt(t)
	return mmc
}

// SetNillableOrganizedAt sets the "organized_at" field if the given value is not nil.
func (mmc *MessageMediaCreate) SetNillableOrganizedAt(t *time.Time) *MessageMediaCreate {
	if t != nil {
		mmc.SetOrganizedAt(*t)
	}
	return mmc
}

// Mutation returns the MessageMediaMutation object of the builder.
func (mmc *MessageMediaCreate) Mutation() *MessageMediaMutation {
	return mmc.mutation
}

// Save creates the MessageMedia in the database.
func (mmc *MessageMediaCreate) Save(ctx context.Context) (*MessageMedia, error) {
	return withHooks(ctx, mmc.sqlSave, mmc.mutation, mmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mmc *MessageMediaCreate) SaveX(ctx context.Context) *MessageMedia {
	v, err := mmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmc *MessageMediaCreate) Exec(ctx context.Context) error {
	_, err := mmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmc *MessageMediaCreate) ExecX(ctx context.Context) {
	if err := mmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mmc *MessageMediaCreate) check() error {
	if _, ok := mmc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message_id", err: errors.New(`ent: missing required field "MessageMedia.message_id"`)}
	}
	if _, ok := mmc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "MessageMedia.type"`)}
	}
	if _, ok := mmc.mutation.Full(); !ok {
		return &ValidationError{Name: "full", err: errors.New(`ent: missing required field "MessageMedia.full"`)}
	}
	if _, ok := mmc.mutation.StashID(); !ok {
		return &ValidationError{Name: "stash_id", err: errors.New(`ent: missing required field "MessageMedia.stash_id"`)}
	}
	return nil
}

func (mmc *MessageMediaCreate) sqlSave(ctx context.Context) (*MessageMedia, error) {
	if err := mmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mmc.mutation.id = &_node.ID
	mmc.mutation.done = true
	return _node, nil
}

func (mmc *MessageMediaCreate) createSpec() (*MessageMedia, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageMedia{config: mmc.config}
		_spec = sqlgraph.NewCreateSpec(messagemedia.Table, sqlgraph.NewFieldSpec(messagemedia.FieldID, field.TypeInt))
	)
	if value, ok := mmc.mutation.MessageID(); ok {
		_spec.SetField(messagemedia.FieldMessageID, field.TypeInt, value)
		_node.MessageID = value
	}
	if value, ok := mmc.mutation.GetType(); ok {
		_spec.SetField(messagemedia.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mmc.mutation.Full(); ok {
		_spec.SetField(messagemedia.FieldFull, field.TypeString, value)
		_node.Full = value
	}
	if value, ok := mmc.mutation.DownloadedAt(); ok {
		_spec.SetField(messagemedia.FieldDownloadedAt, field.TypeTime, value)
		_node.DownloadedAt = value
	}
	if value, ok := mmc.mutation.StashID(); ok {
		_spec.SetField(messagemedia.FieldStashID, field.TypeString, value)
		_node.StashID = value
	}
	if value, ok := mmc.mutation.OrganizedAt(); ok {
		_spec.SetField(messagemedia.FieldOrganizedAt, field.TypeTime, value)
		_node.OrganizedAt = value
	}
	return _node, _spec
}

// MessageMediaCreateBulk is the builder for creating many MessageMedia entities in bulk.
type MessageMediaCreateBulk struct {
	config
	builders []*MessageMediaCreate
}

// Save creates the MessageMedia entities in the database.
func (mmcb *MessageMediaCreateBulk) Save(ctx context.Context) ([]*MessageMedia, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mmcb.builders))
	nodes := make([]*MessageMedia, len(mmcb.builders))
	mutators := make([]Mutator, len(mmcb.builders))
	for i := range mmcb.builders {
		func(i int, root context.Context) {
			builder := mmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMediaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mmcb *MessageMediaCreateBulk) SaveX(ctx context.Context) []*MessageMedia {
	v, err := mmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmcb *MessageMediaCreateBulk) Exec(ctx context.Context) error {
	_, err := mmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmcb *MessageMediaCreateBulk) ExecX(ctx context.Context) {
	if err := mmcb.Exec(ctx); err != nil {
		panic(err)
	}
}