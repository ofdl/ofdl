package main

import (
	_ "embed"

	"github.com/ofdl/ofdl/cmd"
)

//go:embed ofdl.example.yaml
var ExampleConfig []byte

func init() {
	cmd.ExampleConfig = ExampleConfig
}
