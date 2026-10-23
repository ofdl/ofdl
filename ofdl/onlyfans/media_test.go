package onlyfans

import (
	"encoding/json"
	"testing"
)

func TestMediaURL(t *testing.T) {
	cases := []struct {
		name string
		json string
		want string
	}{
		{"files.full.url", `{"id":1,"type":"photo","files":{"full":{"url":"https://cdn/full.jpg"},"thumb":{"url":"https://cdn/t.jpg"}}}`, "https://cdn/full.jpg"},
		{"source.source", `{"id":1,"type":"video","source":{"source":"https://cdn/v.mp4"}}`, "https://cdn/v.mp4"},
		{"old top-level full", `{"id":1,"type":"photo","full":"https://cdn/old.jpg"}`, "https://cdn/old.jpg"},
		{"DRM video has no URL", `{"id":1,"type":"video","files":{"full":{"url":null},"drm":{"manifest":{"dash":"https://cdn/x.mpd"}}}}`, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var m Media
			if err := json.Unmarshal([]byte(c.json), &m); err != nil {
				t.Fatal(err)
			}
			if m.Full != c.want {
				t.Fatalf("Media.Full = %q, want %q", m.Full, c.want)
			}

			var mm MessageMedia
			if err := json.Unmarshal([]byte(c.json), &mm); err != nil {
				t.Fatal(err)
			}
			got := ""
			if mm.Src != nil {
				got = *mm.Src
			}
			if c.name == "old top-level full" {
				return // messages used src, not full
			}
			if got != c.want {
				t.Fatalf("MessageMedia.Src = %q, want %q", got, c.want)
			}
		})
	}
}

func TestMessageMediaKeepsSrc(t *testing.T) {
	var mm MessageMedia
	if err := json.Unmarshal([]byte(`{"id":1,"type":"photo","canView":true,"src":"https://cdn/src.jpg","files":{"full":{"url":"https://cdn/full.jpg"}}}`), &mm); err != nil {
		t.Fatal(err)
	}
	if mm.Src == nil || *mm.Src != "https://cdn/src.jpg" || !mm.CanView {
		t.Fatalf("got %+v", mm)
	}
}
