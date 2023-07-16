package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ExampleConfig []byte

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long: `OFDL Configuration

The configuration file is stored in YAML format. The default location is
$PWD/ofdl.yaml. The following is a comprehensive example configuration file:

  # OFDL Config
  ## Location of the database file
  database: ofdl.sqlite
  
  ## Website App Token
  app-token: 33d57ade8c02dbc5a333db99ff9ae26a
  
  # Chromium Config
  chromium:
  	# Path to Chromium executable
  	exec: /usr/bin/chromium
  	# Path to Chromium profile directory
  	profile: ~/.config/chromium/Default
  
  # Authentication Config
  auth:
  	cookie:
  	user-agent:
  	user-id:
  	x-bc:
  
  # Aria2 Config
  aria2:
  	# Address of Aria2 WebSocket RPC server
  	addr: ws://localhost:6800/jsonrpc
  	# Aria2 RPC secret token
  	secret: secret
  	# Root directory for Aria downloads
  	root: /ofdl
`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := os.Stat("ofdl.yaml")
		if err == nil {
			fmt.Println("Config file already exists")
			return nil
		}

		v := viper.New()
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewReader(ExampleConfig))

		return v.WriteConfigAs("ofdl.yaml")
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set {key} {value}",
	Short: "Set config value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set(args[0], args[1])

		return viper.WriteConfig()
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get {key?}",
	Short: "Get config value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for _, k := range viper.AllKeys() {
				fmt.Printf("%s: %s\n", k, viper.Get(k))
			}
		} else {
			fmt.Printf("%s: %s\n", args[0], viper.Get(args[0]))
		}
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	CLI.AddCommand(configCmd)
}
