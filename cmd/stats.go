package cmd

import (
	"fmt"

	"github.com/ofdl/ofdl/ent/media"
	"github.com/ofdl/ofdl/ent/messagemedia"
	"github.com/ofdl/ofdl/ent/subscription"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:               "stats",
	Short:             "Print database statistics",
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		subCount, err := OFDL.Ent.Subscription.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count: %d\n", subCount)

		uosCount, err := OFDL.Ent.Subscription.Query().Where(subscription.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count (Unorganized): %d\n", uosCount)

		mpCount, err := OFDL.Ent.Post.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Post Count: %d\n", mpCount)

		mCount, err := OFDL.Ent.Media.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count: %d\n", mCount)

		udmCount, err := OFDL.Ent.Media.Query().Where(media.DownloadedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Undownloaded): %d\n", udmCount)

		uomCount, err := OFDL.Ent.Media.Query().Where(media.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Unorganized): %d\n", uomCount)

		mmCount, err := OFDL.Ent.MessageMedia.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count: %d\n", mmCount)

		udmmCount, err := OFDL.Ent.MessageMedia.Query().Where(messagemedia.DownloadedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count (Undownloaded): %d\n", udmmCount)

		uommCount, err := OFDL.Ent.MessageMedia.Query().Where(messagemedia.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count (Unorganized): %d\n", uommCount)

		return nil
	},
}

func init() {
	CLI.AddCommand(statsCmd)
}
