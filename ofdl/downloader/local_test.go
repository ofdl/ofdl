package downloader

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"gorm.io/gorm"
)

type fakeMedia struct {
	url    string
	marked bool
}

func (f *fakeMedia) Directory() string { return "sub" }
func (f *fakeMedia) Filename() string  { return "file.jpg" }
func (f *fakeMedia) URL() string       { return f.url }
func (f *fakeMedia) MarkDownloaded(*gorm.DB) error {
	f.marked = true
	return nil
}

func download(t *testing.T, status int) (*fakeMedia, string, error) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte("body"))
	}))
	defer srv.Close()

	root := t.TempDir()
	d := &LocalDownloader{root: root}
	m := &fakeMedia{url: srv.URL}
	progress, done := d.DownloadOne(m)
	go func() {
		for range progress {
		}
	}()
	return m, filepath.Join(root, "sub", "file.jpg"), <-done
}

func TestDownloadRefusedLinkIsNotMarkedDownloaded(t *testing.T) {
	m, path, err := download(t, http.StatusForbidden)
	if err == nil {
		t.Fatal("expected an error for a 403 response")
	}
	if m.marked {
		t.Fatal("media was marked downloaded")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("error page was saved as %s", path)
	}
}

func TestDownloadSavesAndMarks(t *testing.T) {
	m, path, err := download(t, http.StatusOK)
	if err != nil {
		t.Fatal(err)
	}
	if !m.marked {
		t.Fatal("media wasn't marked downloaded")
	}
	if b, err := os.ReadFile(path); err != nil || string(b) != "body" {
		t.Fatalf("file = %q, %v", b, err)
	}
}
