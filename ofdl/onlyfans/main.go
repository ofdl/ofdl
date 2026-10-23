package onlyfans

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/sling"
	"github.com/spf13/viper"
)

type OnlyFansAPI interface {
	Following() (*Following, error)
	GetSubscriptions(limit, offset int) ([]Subscription, error)
	GetMediaPosts(uid int, beforePublishTime *string) (*PaginatedList[MediaPost], error)
	GetMessages(uid int, id *int) (*PagableList[Message], error)
}

type OnlyFans struct {
	*sling.Sling

	dr *viper.Viper
}

var _ OnlyFansAPI = &OnlyFans{}

func NewOnlyFans() (OnlyFansAPI, error) {
	dr, err := NewDynamicRules()
	if err != nil {
		return nil, err
	}

	s := sling.New().
		Base("https://onlyfans.com/api2/v2/").
		Add("Accept", "application/json, text/plain, */*").
		Add("User-Agent", viper.GetString("auth.user-agent")).
		Add("User-Id", viper.GetString("auth.user-id")).
		Add("X-BC", viper.GetString("auth.x-bc")).
		Add("Cookie", viper.GetString("auth.cookie")).
		Add("App-Token", viper.GetString("app-token"))

	return &OnlyFans{
		Sling: s,
		dr:    dr,
	}, nil
}

func (o *OnlyFans) sign(r *http.Request) error {
	// path= fullpath with query, exactly as sent
	path := r.URL.Path
	if q := r.URL.RawQuery; q != "" {
		path += "?" + q
	}

	// unix timestamp in milliseconds
	unixtime := time.Now().UnixMilli()

	r.Header.Add("sign", Sign(
		o.dr.GetString("static_param"),
		o.dr.GetString("format"),
		o.dr.GetIntSlice("checksum_indexes"),
		o.dr.GetInt("checksum_constant"),
		unixtime,
		path,
		viper.GetString("auth.user-id"),
	))
	r.Header.Add("time", fmt.Sprintf("%d", unixtime))

	return nil
}

// Sign builds the sign header for a request path. format has %x placeholders
// for the SHA-1 hex digest and the checksum, as NewDynamicRules leaves it.
func Sign(staticParam, format string, checksumIndexes []int, checksumConstant int, unixtime int64, path, userID string) string {
	// msg [static_param, unixtime, path, user-id].join(\n)
	msg := []string{
		staticParam,
		fmt.Sprintf("%d", unixtime),
		path,
		userID,
	}

	// sha1(msg)
	sha := sha1.New()
	sha.Write([]byte(strings.Join(msg, "\n")))
	sum := sha.Sum(nil)
	xum := fmt.Sprintf("%x", sum)

	checksum := checksumConstant
	for _, i := range checksumIndexes {
		checksum += int(xum[i])
	}
	if checksum < 0 {
		checksum = -checksum
	}

	return fmt.Sprintf(format, sum, checksum)
}
