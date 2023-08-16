package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.Int("subscription_id"),
		field.String("text"),
		field.String("posted_at"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("media", MessageMedia.Type),
		edge.From("subscription", Subscription.Type).Ref("messages").Unique().Field("subscription_id").Required(),
	}
}
