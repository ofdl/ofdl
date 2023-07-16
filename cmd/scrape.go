package cmd

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape OnlyFans API",
	Long: `Scrape OnlyFans API

This command will scrape the OnlyFans API for subscriptions and media posts. See
ofdl scrape subs --help and ofdl scrape media-posts --help for more information.
`,
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := scrapeSubsCmd.RunE(cmd, args); err != nil {
			return err
		}

		if err := scrapeMediaPostsCmd.RunE(cmd, args); err != nil {
			return err
		}

		return nil
	},
}

var scrapeSubsCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Scrape OnlyFans Subscriptions",
	Long: `Scrape OnlyFans Subscriptions

This command will scrape the OnlyFans API for subscriptions and save them to the
database.
`,
	Aliases: []string{"subs", "s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		subs, err := OFDL.OnlyFans.GetSubscriptions()
		if err != nil {
			return err
		}

		for _, sub := range subs {
			if err := OFDL.Data().SaveSubscription(sub); err != nil {
				return err
			}
		}

		fmt.Printf("Saved %d Subscriptions\n", len(subs))

		return nil
	},
}

var scrapeMediaPostsCmd = &cobra.Command{
	Use:   "media-posts",
	Short: "Scrape OnlyFans Media Posts",
	Long: `Scrape OnlyFans Media Posts

This command will scrape the OnlyFans API for media posts and save them to the
database. This command will also update the head marker for each subscription
that is scraped. This allows for incremental scraping of media posts.
`,
	Aliases: []string{"media", "mp", "m"},
	RunE: func(cmd *cobra.Command, args []string) error {
		subs, err := OFDL.Data().GetEnabledSubscriptions()
		if err != nil {
			return err
		}

		for _, sub := range subs {
			total := 0
			hasMore := true
			var headMarker *string
			var tailMarker *string

			bar := progressbar.Default(1, fmt.Sprintf("%s", sub.Name))

			// Sync Media Posts
			for hasMore {
				// Get a page
				page, err := OFDL.OnlyFans.GetMediaPosts(int(sub.ID), tailMarker)
				if err != nil {
					return err
				}

				// Handle Pagination
				hasMore = page.HasMore
				tailMarker = &page.TailMarker

				if total == 0 {
					total = page.Counters.MediasCount
					headMarker = &page.HeadMarker
					bar.ChangeMax(total)
				}

				// Save Media Posts
				for _, m := range page.List {
					if err := OFDL.Data().SaveMediaPost(m); err != nil {
						return err
					}
					bar.Add(len(m.Media))
				}

				// Consider breaking due to head marker
				if hasMore && page.HeadMarker < sub.HeadMarker {
					bar.Close()
					fmt.Println("Skipping already scraped pages")
					hasMore = false
				}

			}

			if headMarker != nil {
				// Update Subscription HeadMarker
				sub.HeadMarker = *headMarker
				if err := OFDL.DB.Save(sub).Error; err != nil {
					return err
				}
			}

			bar.Close()

		}

		return nil
	},
}

func init() {
	scrapeCmd.AddCommand(scrapeSubsCmd)
	scrapeCmd.AddCommand(scrapeMediaPostsCmd)
	CLI.AddCommand(scrapeCmd)
}
