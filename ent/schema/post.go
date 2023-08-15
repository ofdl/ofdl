package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("subscription_id"),
		field.String("text"),
		field.String("posted_at"),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("medias", Media.Type).
			StorageKey(edge.Column("post_id")),
	}
}
