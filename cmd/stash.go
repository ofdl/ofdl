package cmd

import (
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Organize Stash Performers, Images, and Scenes",
	Long: `Organize Stash Performers, Images, and Scenes

After scanning for new media in Stash, this command will organize the media into
performers, images, and scenes. See ofdl stash subs --help and ofdl stash media
--help for more information.
`,
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := stashSubsCmd.RunE(cmd, args); err != nil {
			return err
		}

		if err := stashMediaCmd.RunE(cmd, args); err != nil {
			return err
		}

		return nil
	},
}

var stashSubsCmd = &cobra.Command{
	Use:     "subscriptions",
	Aliases: []string{"subs", "s"},
	Short:   "Organize Subscriptions to Stash Performers",
	Long: `Organize Subscriptions to Stash Performers

This command will organize subscriptions into performers. It assigns their name,
URL, and photo.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ms, err := OFDL.GetUnorganizedSubscriptions(100)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(ms)), "Organizing Subscriptions")

		for _, sub := range ms {
			if err := OFDL.OrganizeSubscription(sub); err != nil {
				return err
			}
			bar.Add(1)
		}

		return nil
	},
}

var stashMediaCmd = &cobra.Command{
	Use:     "media",
	Aliases: []string{"m"},
	Short:   "Organize Media to Stash Images and Scenes",
	Long: `Organize Media to Stash Images and Scenes

This command will organize media into images and scenes. The corresponding
subscription must have a Stash ID (already be organized) for this to work. It
assigns the Performer, Studio, Date, and Text on both Images and Scenes. Stash
will probably organize Images into Galleries per-post.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ms, err := OFDL.GetUnorganizedMedia(1000)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(ms)), "Organizing Media")

		for _, m := range ms {
			if err := OFDL.OrganizeMedia(m); err != nil {
				return err
			}
			bar.Add(1)
		}

		return nil
	},
}

func init() {
	stashCmd.AddCommand(stashSubsCmd)
	stashCmd.AddCommand(stashMediaCmd)
	CLI.AddCommand(stashCmd)

	stashCmd.PersistentFlags().String("address", "", "Stash GraphQL Address")
	viper.BindPFlag("stash.address", stashCmd.PersistentFlags().Lookup("address"))
	stashCmd.PersistentFlags().String("studio-id", "", "Stash Studio ID")
	viper.BindPFlag("stash.studio_id", stashCmd.PersistentFlags().Lookup("studio-id"))
}
