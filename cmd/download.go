package cmd

import (
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"dl"},
	Short:   "Download media",
	Long: `Download media

OFDL uses Aria2 to manage downloads. After you've configured Aria2, this command
will queue up to 1,000 undownloaded media.
`,
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		missing, err := OFDL.GetMissingMedia(1000)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(missing)), "Queueing downloads")
		for _, m := range missing {
			_, err := OFDL.DownloadMedia(m)
			if err != nil {
				return err
			}
			bar.Add(1)
		}
		return nil
	},
}

func init() {
	CLI.AddCommand(downloadCmd)

	downloadCmd.Flags().String("address", "", "Aria2 WebSocket RPC Address")
	viper.BindPFlag("aria2.address", downloadCmd.Flags().Lookup("address"))
	downloadCmd.Flags().String("secret", "", "Aria2 RPC Secret")
	viper.BindPFlag("aria2.secret", downloadCmd.Flags().Lookup("secret"))
	downloadCmd.Flags().String("root", "", "Root directory for downloads")
	viper.BindPFlag("aria2.root", downloadCmd.Flags().Lookup("root"))
}
