package onlyfans

import (
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func NewDynamicRules() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("json")

	r, err := http.Get("https://github.com/DIGITALCRIMINALS/dynamic-rules/raw/main/onlyfans.json")
	if err != nil {
		return nil, err
	}

	if err := v.ReadConfig(r.Body); err != nil {
		return nil, err
	}

	f := v.GetString("format")
	f = strings.ReplaceAll(f, "{}", `%x`)
	f = strings.ReplaceAll(f, "{:x}", `%x`)
	v.Set("format", f)

	return v, nil
}
