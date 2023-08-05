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
	GetSubscriptions() ([]Subscription, error)
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
	// path= fullpath with query
	path := r.URL.Path
	if q := r.URL.Query().Encode(); q != "" {
		path += "?" + q
	}

	// unix timestamp
	unixtime := time.Now().Unix()

	// msg [static_param, unixtime, path, user-id].join(\n)
	msg := []string{
		o.dr.GetString("static_param"),
		fmt.Sprintf("%d", unixtime),
		path,
		viper.GetString("auth.user-id"),
	}

	// sha1(msg)
	sha := sha1.New()
	sha.Write([]byte(strings.Join(msg, "\n")))
	sum := sha.Sum(nil)
	xum := fmt.Sprintf("%x", sum)

	checksum := o.dr.GetInt("checksum_constant")
	for _, i := range o.dr.GetIntSlice("checksum_indexes") {
		checksum += int(xum[i])
	}

	r.Header.Add("sign", fmt.Sprintf(
		o.dr.GetString("format"),
		sum,
		checksum,
	))
	r.Header.Add("time", fmt.Sprintf("%d", unixtime))

	return nil
}
