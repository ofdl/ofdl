package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Int("post_id"),
		field.String("type"),
		field.String("full"),
		field.Time("downloaded_at").Optional(),
		field.String("stash_id"),
		field.Time("organized_at").Optional(),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return nil
}
