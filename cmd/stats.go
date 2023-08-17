package cmd

import (
	"context"
	"fmt"

	"github.com/ofdl/ofdl/ent"
	"github.com/ofdl/ofdl/ent/media"
	"github.com/ofdl/ofdl/ent/messagemedia"
	"github.com/ofdl/ofdl/ent/subscription"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Print database statistics",
	RunE: Inject(func(ctx context.Context, Ent *ent.Client) error {
		subCount, err := Ent.Subscription.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count: %d\n", subCount)

		uosCount, err := Ent.Subscription.Query().Where(subscription.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count (Unorganized): %d\n", uosCount)

		mpCount, err := Ent.Post.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Post Count: %d\n", mpCount)

		mCount, err := Ent.Media.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count: %d\n", mCount)

		udmCount, err := Ent.Media.Query().Where(media.DownloadedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Undownloaded): %d\n", udmCount)

		uomCount, err := Ent.Media.Query().Where(media.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Unorganized): %d\n", uomCount)

		mmCount, err := Ent.MessageMedia.Query().Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count: %d\n", mmCount)

		udmmCount, err := Ent.MessageMedia.Query().Where(messagemedia.DownloadedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count (Undownloaded): %d\n", udmmCount)

		uommCount, err := Ent.MessageMedia.Query().Where(messagemedia.OrganizedAtIsNil()).Count(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count (Unorganized): %d\n", uommCount)

		return nil
	}),
}

func init() {
	CLI.AddCommand(statsCmd)
}
