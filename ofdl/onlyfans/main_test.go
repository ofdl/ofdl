package onlyfans

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/viper"
)

// The expected value comes from OF-Scraper's create_login_sign, which signs
// requests that OnlyFans accepts, run on the same inputs.
func TestSignMatchesWorkingImplementation(t *testing.T) {
	got := Sign(
		"abc",
		"123:%x:%x:def",
		[]int{0, 5, 10, 39},
		-5000, // the sum is negative, so the checksum must be made positive
		1700000000000,
		"/api2/v2/users/me?offset=0&limit=10",
		"12345",
	)
	want := "123:3eff58c9c0bd3e5685fab922702404dd310c32fb:1259:def"
	if got != want {
		t.Fatalf("Sign() = %q, want %q", got, want)
	}
}

func TestSignHeaderUsesQueryAsSentAndMilliseconds(t *testing.T) {
	dr := viper.New()
	dr.Set("static_param", "abc")
	dr.Set("format", "123:%x:%x:def")
	dr.Set("checksum_indexes", []int{0, 5, 10, 39})
	dr.Set("checksum_constant", -5000)
	viper.Set("auth.user-id", "12345")
	o := &OnlyFans{dr: dr}

	// Query().Encode() would sort this to limit=10&offset=0, signing a path
	// that differs from the one sent
	r, err := http.NewRequest("GET", "https://onlyfans.com/api2/v2/users/me?offset=0&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	before := time.Now().UnixMilli()
	if err := o.sign(r); err != nil {
		t.Fatal(err)
	}

	ts, err := strconv.ParseInt(r.Header.Get("time"), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	if ts < before || ts > time.Now().UnixMilli() {
		t.Fatalf("time header %d isn't the current time in milliseconds", ts)
	}
	want := Sign("abc", "123:%x:%x:def", []int{0, 5, 10, 39}, -5000, ts, "/api2/v2/users/me?offset=0&limit=10", "12345")
	if got := r.Header.Get("sign"); got != want {
		t.Fatalf("sign header = %q, want %q", got, want)
	}
}
