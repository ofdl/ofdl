package onlyfans

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func NewDynamicRules() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("json")

	url := viper.GetString("dynamic-rules-url")

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching dynamic rules from %s: %s", url, r.Status)
	}

	if err := v.ReadConfig(r.Body); err != nil {
		return nil, err
	}

	// Some rule sets give prefix and suffix instead of a format
	f := v.GetString("format")
	if f == "" && v.IsSet("prefix") && v.IsSet("suffix") {
		f = fmt.Sprintf("%s:{}:{:x}:%s", v.GetString("prefix"), v.GetString("suffix"))
	}
	f = strings.ReplaceAll(f, "{}", `%x`)
	f = strings.ReplaceAll(f, "{:x}", `%x`)
	v.Set("format", f)

	return v, nil
}
