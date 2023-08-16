package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MessageMedia holds the schema definition for the MessageMedia entity.
type MessageMedia struct {
	ent.Schema
}

// Fields of the MessageMedia.
func (MessageMedia) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.Int("message_id"),
		field.String("type"),
		field.String("src").Optional(),

		field.Time("downloaded_at").Optional(),
		field.String("stash_id"),
		field.Time("organized_at").Optional(),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MessageMedia.
func (MessageMedia) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).Ref("media").Unique().Field("message_id").Required(),
	}
}
