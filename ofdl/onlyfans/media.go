package onlyfans

import (
	"encoding/json"
	"fmt"
)

func (o *OnlyFans) GetMediaPosts(uid int, beforePublishTime *string) (*PaginatedList[MediaPost], error) {
	page := &PaginatedList[MediaPost]{}
	r, err := o.New().
		Get(fmt.Sprintf("users/%d/posts/medias", uid)).
		QueryStruct(&GetMediaOptions{
			Limit:             25,
			Order:             "publish_date_desc",
			SkipUsers:         "all",
			Format:            "infinite",
			Counters:          1,
			BeforePublishTime: beforePublishTime,
		}).
		Request()
	if err != nil {
		return nil, err
	}

	if err := o.sign(r); err != nil {
		return nil, err
	}

	_, err = o.Do(r, page, nil)
	if err != nil {
		return nil, err
	}

	return page, nil
}

type GetMediaOptions struct {
	Limit             int     `url:"limit,omitempty"`
	Order             string  `url:"order,omitempty"`
	SkipUsers         string  `url:"skip_users,omitempty"`
	Format            string  `url:"format,omitempty"`
	Counters          int     `url:"counters,omitempty"`
	BeforePublishTime *string `url:"beforePublishTime,omitempty"`
}

type MediaPost struct {
	ID     int `json:"id"`
	Author struct {
		ID int `json:"id"`
	} `json:"author"`
	Text            string  `json:"text"`
	PostedAt        string  `json:"postedAt"`
	PostedAtPrecise string  `json:"postedAtPrecise"`
	Media           []Media `json:"media"`
}

type Media struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	Type      string `json:"type"`
	Full      string `json:"full"`
}

// mediaFiles is where the API now puts media URLs; the top-level full and src
// fields are no longer filled in.
type mediaFiles struct {
	Files struct {
		Full struct {
			URL string `json:"url"`
		} `json:"full"`
	} `json:"files"`
	Source struct {
		Source string `json:"source"`
	} `json:"source"`
}

func (f mediaFiles) url() string {
	if f.Files.Full.URL != "" {
		return f.Files.Full.URL
	}
	return f.Source.Source
}

func (m *Media) UnmarshalJSON(b []byte) error {
	type plain Media
	if err := json.Unmarshal(b, (*plain)(m)); err != nil {
		return err
	}
	if m.Full == "" {
		var f mediaFiles
		if err := json.Unmarshal(b, &f); err != nil {
			return err
		}
		m.Full = f.url()
	}
	return nil
}

type PaginatedList[T any] struct {
	HasMore    bool   `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	TailMarker string `json:"tailMarker"`
	List       []T    `json:"list"`
	Counters   struct {
		MediasCount int `json:"mediasCount"`
		PostsCount  int `json:"postsCount"`
	} `json:"counters"`
}
