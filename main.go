//go:generate go run ./gen
package main

import "github.com/ofdl/ofdl/cmd"

func main() {
	cmd.CLI.Execute()
}
