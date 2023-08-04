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

OFDL uses your configured Downloader (Local or Aria2) to manage downloads.
After you've configured your downloader, this command will queue up to 1,000
ndownloaded media.
`,
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := downloadPostsCmd.RunE(cmd, args); err != nil {
			return err
		}

		if err := downloadMessagesCmd.RunE(cmd, args); err != nil {
			return err
		}

		return nil
	},
}

var downloadPostsCmd = &cobra.Command{
	Use:     "media-posts",
	Short:   "Download media posts",
	Aliases: []string{"media", "mp"},
	RunE: func(cmd *cobra.Command, args []string) error {
		missing, err := OFDL.GetMissingMedia(1000)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(missing)), "Downloading media")
		progress := OFDL.Downloader.DownloadMany(missing)
		for err := range progress {
			if err != nil {
				return err
			}
			bar.Add(1)
		}

		return nil
	},
}

var downloadMessagesCmd = &cobra.Command{
	Use:     "messages",
	Short:   "Download messages",
	Aliases: []string{"msg"},
	RunE: func(cmd *cobra.Command, args []string) error {
		missing, err := OFDL.GetMissingMessageMedia(1000)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(missing)), "Downloading message media")
		progress := OFDL.Downloader.DownloadMany(missing)
		for err := range progress {
			if err != nil {
				return err
			}
			bar.Add(1)
		}

		return nil
	},
}

func init() {
	downloadCmd.AddCommand(downloadPostsCmd)
	downloadCmd.AddCommand(downloadMessagesCmd)
	CLI.AddCommand(downloadCmd)

	downloadCmd.PersistentFlags().String("downloader", "", "Download method (local, aria2)")
	viper.BindPFlag("downloads.downloader", downloadCmd.PersistentFlags().Lookup("downloader"))

	downloadCmd.PersistentFlags().String("local-root", "", "Root directory for Local downloads")
	viper.BindPFlag("downloads.local.root", downloadCmd.PersistentFlags().Lookup("local-root"))

	downloadCmd.PersistentFlags().String("aria2-address", "", "Aria2 WebSocket RPC Address")
	viper.BindPFlag("downloads.aria2.address", downloadCmd.PersistentFlags().Lookup("aria2-address"))
	downloadCmd.PersistentFlags().String("aria2-secret", "", "Aria2 RPC Secret")
	viper.BindPFlag("downloads.aria2.secret", downloadCmd.PersistentFlags().Lookup("aria2-secret"))
	downloadCmd.PersistentFlags().String("aria2-root", "", "Root directory for Aria2 downloads")
	viper.BindPFlag("downloads.aria2.root", downloadCmd.PersistentFlags().Lookup("aria2-root"))
}
