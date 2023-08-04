package cmd

import (
	"github.com/ofdl/ofdl/ofdl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var OFDL *ofdl.OFDL

func UseOFDL(cmd *cobra.Command, args []string) error {
	o, err := ofdl.NewOFDL()
	if err != nil {
		return err
	}

	OFDL = o
	return nil
}

var CLI = &cobra.Command{
	Use:   "ofdl",
	Short: "OnlyFans Downloader",
	Long: `OnlyFans Downloader

ofdl is a command-line tool for downloading media from OnlyFans. It uses the
OnlyFans API to scrape subscriptions and media posts, and optionally Aria2 to
anage downloads.

ofdl is not affiliated with OnlyFans in any way. It is a third-party tool that
uses the OnlyFans API to download media. It is not endorsed by OnlyFans, and
should be used at your own risk.

A typical ofdl setup looks like this:

  ofdl config init
  ofdl config set chromium.exec /path/to/chromium
  ofdl config set chromium.profile /path/to/chromium/profile
  ofdl config set downloads.downloader aria2
  ofdl config set downloads.aria2.address ws://localhost:6800/jsonrpc
  ofdl config set downloads.aria2.secret my-secret
  ofdl config set downloads.aria2.root /path/to/downloads

A typical ofdl workflow looks like this:

  # Exit your browser
  ofdl auth
  # Login, press enter
  ofdl scrape
  ofdl download
  # Wait for downloads to finish
  # Scan for new media in Stash
  ofdl stash
`,
	Version: Version,
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("ofdl")
	viper.ReadInConfig()

	viper.SetDefault("database", "ofdl.sqlite")
	viper.SetDefault("app-token", "33d57ade8c02dbc5a333db99ff9ae26a")

	CLI.PersistentFlags().String("database", "", "Path to database file")
	viper.BindPFlag("database", CLI.PersistentFlags().Lookup("database"))
}
