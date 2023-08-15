package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar"),
		field.String("header"),
		field.String("name"),
		field.String("username"),
		field.String("head_marker"),
		field.String("stash_id"),
		field.Time("organized_at").Optional(),
		field.Bool("enabled").Default(true),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).
			StorageKey(edge.Column("subscription_id")),
	}
}
