package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.Int("subscription_id"),
		field.String("text"),
		field.String("posted_at"),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("medias", Media.Type),
		edge.From("subscription", Subscription.Type).Ref("posts").Unique().Field("subscription_id").Required(),
	}
}
