package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.Int("post_id"),
		field.String("type"),
		field.String("full"),
		field.Time("downloaded_at").Optional(),
		field.String("stash_id").Optional(),
		field.Time("organized_at").Optional(),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("medias").Unique().Field("post_id").Required(),
	}
}
