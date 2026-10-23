package onlyfans

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dghubble/sling"
	"github.com/spf13/viper"
)

func testAPI(t *testing.T, status int, body string) *OnlyFans {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)

	dr := viper.New()
	dr.Set("static_param", "abc")
	dr.Set("format", "1:%x:%x:2")
	return &OnlyFans{Sling: sling.New().Base(srv.URL + "/"), dr: dr}
}

func TestExpiredSessionIsAnError(t *testing.T) {
	o := testAPI(t, http.StatusUnauthorized, `{"error":{"code":401,"message":"Please refresh the page"}}`)
	_, err := o.Following()
	if err == nil {
		t.Fatal("expected an error for a 401 response")
	}
	if !strings.Contains(err.Error(), "401") || !strings.Contains(err.Error(), "Please refresh the page") {
		t.Fatalf("error doesn't show the status and message: %v", err)
	}
}

func TestSuccessfulResponseDecodes(t *testing.T) {
	o := testAPI(t, http.StatusOK, `{"usersCount":6}`)
	f, err := o.Following()
	if err != nil {
		t.Fatal(err)
	}
	if f.UsersCount != 6 {
		t.Fatalf("UsersCount = %d", f.UsersCount)
	}
}
