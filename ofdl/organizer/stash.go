package organizer

import (
	"context"
	"fmt"
	"time"

	"github.com/ofdl/ofdl/ent"
	"github.com/shurcooL/graphql"
	"github.com/spf13/viper"
)

type Organizer interface {
	OrganizeSubscription(sub *ent.Subscription) error
	OrganizeMedia(m *ent.Media) error
	OrganizeMessageMedia(m *ent.MessageMedia) error
}

type Stash struct {
	*graphql.Client
	ctx context.Context
}

func NewOrganizer(ctx context.Context) Organizer {
	return &Stash{
		Client: graphql.NewClient(viper.GetString("stash.address"), nil),
		ctx:    ctx,
	}
}

func (s *Stash) OrganizeSubscription(sub *ent.Subscription) error {
	lu := &PerformerLookup{}
	if err := s.Query(s.ctx, lu, map[string]interface{}{
		"name": graphql.String(sub.Username),
	}); err != nil {
		return err
	}

	// If missing, create it
	if lu.FindPerformers.Count == 0 {
		pc := &PerformerCreate{}
		if err := s.Mutate(s.ctx, pc, map[string]interface{}{
			"performer": PerformerCreateInput{
				"name":  sub.Username,
				"url":   fmt.Sprintf("https://onlyfans.com/%s", sub.Username),
				"image": sub.Avatar,
			},
		}); err != nil {
			return err
		}

		return sub.Update().
			SetStashID(pc.PerformerCreate.ID).
			SetOrganizedAt(time.Now()).
			Exec(s.ctx)
	}

	// If found, update it
	id := lu.FindPerformers.Performers[0].ID
	pc := &PerformerUpdate{}
	if err := s.Mutate(s.ctx, pc, map[string]interface{}{
		"performer": PerformerUpdateInput{
			"id":    id,
			"name":  sub.Username,
			"url":   fmt.Sprintf("https://onlyfans.com/%s", sub.Username),
			"image": sub.Avatar,
		},
	}); err != nil {
		return err
	}

	return sub.Update().
		SetStashID(pc.PerformerUpdate.ID).
		SetOrganizedAt(time.Now()).
		Exec(s.ctx)
}

func (s *Stash) OrganizeMedia(m *ent.Media) error {
	if m.GetFull() == "" {
		return m.MarkOrganized(s.ctx)
	}

	date, err := time.Parse(time.RFC3339, m.Edges.Post.GetPostedAt())
	if err != nil {
		return err
	}
	sm, err := s.organize(
		m.Type,
		m.Filename(),
		m.Edges.Post.Text,
		m.Edges.Post.Edges.Subscription.StashID,
		date,
	)
	if err != nil {
		return err
	}

	// update the ent
	return m.Organize(s.ctx, sm.ID)
}

func (s *Stash) OrganizeMessageMedia(m *ent.MessageMedia) error {
	if m.GetFull() == "" {
		return m.MarkOrganized(s.ctx)
	}

	date, err := time.Parse(time.RFC3339, m.Edges.Message.GetPostedAt())
	if err != nil {
		return err
	}
	sm, err := s.organize(
		m.Type,
		m.Filename(),
		m.Edges.Message.Text,
		m.Edges.Message.Edges.Subscription.StashID,
		date,
	)
	if err != nil {
		return err
	}

	// update the ent
	return m.Organize(s.ctx, sm.ID)
}

func (o *Stash) organize(kind, filename, title, performerID string, date time.Time) (*StashMedia, error) {
	var sm StashLookup
	var sr interface{}

	switch kind {
	case "photo":
		sm = &ImageLookup{}
		sr = &ImageUpdate{}
	case "video", "gif":
		sm = &SceneLookup{}
		sr = &SceneUpdate{}
	default:
		return nil, nil
	}

	if err := o.Query(o.ctx, sm, map[string]interface{}{
		"path": graphql.String(filename),
	}); err != nil {
		return nil, err
	}

	if sm.Count() != 1 {
		return nil, nil
	}
	m := sm.Medias()[0]

	vars := map[string]interface{}{
		"id":          m.ID,
		"title":       graphql.String(title),
		"performerId": graphql.ID(performerID),
		"studioId":    graphql.ID(viper.GetString("stash.studio_id")),
		"date":        graphql.String(date.Format("2006-01-02")),
	}
	if err := o.Mutate(o.ctx, sr, vars); err != nil {
		return &m, err
	}

	return &m, nil
}
