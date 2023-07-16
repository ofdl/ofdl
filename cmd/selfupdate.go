package cmd

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "selfupdate",
	Short: "Update to the latest version of ofdl",
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := semver.Parse(Version[1:])
		if err != nil {
			v = semver.MustParse("0.0.0")
		}
		latest, err := selfupdate.UpdateSelf(v, "ofdl/ofdl")
		if err != nil {
			return err
		}

		if latest.Version.Equals(v) {
			// latest version is the same as current version. It means current binary is up to date.
			log.Println("Current binary is the latest version", Version)
		} else {
			log.Println("Successfully updated to version", latest.Version)
			log.Println("Release note:\n", latest.ReleaseNotes)
		}
		return nil
	},
}

func init() {
	CLI.AddCommand(selfUpdateCmd)
}
