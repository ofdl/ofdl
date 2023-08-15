package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MessageMedia holds the schema definition for the MessageMedia entity.
type MessageMedia struct {
	ent.Schema
}

// Fields of the MessageMedia.
func (MessageMedia) Fields() []ent.Field {
	return []ent.Field{
		field.Int("message_id"),
		field.String("type"),
		field.String("full"),
		field.Time("downloaded_at").Optional(),
		field.String("stash_id"),
		field.Time("organized_at").Optional(),
	}
}

// Edges of the MessageMedia.
func (MessageMedia) Edges() []ent.Edge {
	return nil
}
