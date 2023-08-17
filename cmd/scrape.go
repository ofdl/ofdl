package cmd

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/ofdl/ofdl/ent"
	"github.com/ofdl/ofdl/ent/subscription"
	"github.com/ofdl/ofdl/ofdl/onlyfans"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape OnlyFans API",
	Long: `Scrape OnlyFans API

This command will scrape the OnlyFans API for subscriptions, media posts, and
messages. See ofdl scrape subs --help, ofdl scrape media-posts --help, and
ofdl scrape msg --help for more information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := scrapeSubsCmd.RunE(cmd, args); err != nil {
			return err
		}

		if err := scrapeMediaPostsCmd.RunE(cmd, args); err != nil {
			return err
		}

		if err := scrapeMessagesCmd.RunE(cmd, args); err != nil {
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
	RunE: Inject(func(ctx context.Context, OnlyFans onlyfans.OnlyFansAPI, Ent *ent.Client) error {
		f, err := OnlyFans.Following()
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(f.UsersCount), "Scraping subscriptions")
		defer bar.Close()

		l := 50
		for o := 0; o < f.UsersCount; o += l {
			subs, err := OnlyFans.GetSubscriptions(l, o)
			if err != nil {
				return err
			}

			for _, v := range subs {
				err := Ent.Subscription.Create().
					SetID(v.ID).
					SetAvatar(v.Avatar).
					SetHeader(v.Header).
					SetName(v.Name).
					SetUsername(v.Username).
					OnConflict(
						sql.ResolveWithNewValues(),
						sql.ResolveWith(func(u *sql.UpdateSet) {
							u.SetIgnore(subscription.FieldID).
								SetIgnore(subscription.FieldEnabled)
						}),
					).
					Exec(ctx)
				if err != nil {
					return err
				}

				bar.Add(1)
			}
		}

		return nil
	}),
}

var scrapeMediaPostsCmd = &cobra.Command{
	Use:   "media-posts",
	Short: "Scrape OnlyFans Media Posts",
	Long: `Scrape OnlyFans Media Posts

This command will scrape the OnlyFans API for media posts and save them to the
database. This command will also update the head marker for each subscription
that is scraped. This allows for incremental scraping of media posts.
`,
	Aliases: []string{"media", "mp"},
	RunE: Inject(func(ctx context.Context, OnlyFans onlyfans.OnlyFansAPI, Ent *ent.Client) error {
		subs, err := Ent.Subscription.Query().Where(subscription.EnabledEQ(true)).All(ctx)
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
				page, err := OnlyFans.GetMediaPosts(int(sub.ID), tailMarker)
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
					postId, err := Ent.Post.Create().
						SetID(m.ID).
						SetSubscriptionID(sub.ID).
						SetText(m.Text).
						SetPostedAt(m.PostedAt).
						OnConflict().UpdateNewValues().
						ID(ctx)
					if err != nil {
						return err
					}

					mc := []*ent.MediaCreate{}
					for _, v := range m.Media {
						mc = append(mc, Ent.Media.Create().
							SetID(v.ID).
							SetPostID(postId).
							SetType(v.Type).
							SetFull(v.Full),
						)
					}

					if err := Ent.Media.CreateBulk(mc...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
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
				if err := sub.Update().SetHeadMarker(*headMarker).Exec(ctx); err != nil {
					return err
				}
			}

			bar.Close()
		}

		return nil
	}),
}

var scrapeMessagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "Scrape OnlyFans Messages",
	Long: `Scrape OnlyFans Messages

This command will scrape the OnlyFans API for messages and save them to the
database.
`,
	Aliases: []string{"msg"},
	// RunE: func(cmd *cobra.Command, args []string) error {
	RunE: Inject(func(ctx context.Context, OnlyFans onlyfans.OnlyFansAPI, Ent *ent.Client) error {
		subs, err := Ent.Subscription.Query().Where(subscription.EnabledEQ(true)).All(ctx)
		if err != nil {
			return err
		}

		for _, sub := range subs {
			hasMore := true
			var nextId *int

			bar := progressbar.Default(-1, fmt.Sprintf("%s", sub.Name))

			for hasMore {
				// Get a page
				page, err := OnlyFans.GetMessages(int(sub.ID), nextId)
				if err != nil {
					return err
				}

				// Handle Pagination
				hasMore = page.HasMore

				// Save Messages
				for _, m := range page.List {
					messageId, err := Ent.Message.Create().
						SetID(m.ID).
						SetSubscriptionID(sub.ID).
						SetText(m.Text).
						SetPostedAt(m.CreatedAt).
						OnConflict().UpdateNewValues().
						ID(ctx)
					if err != nil {
						return err
					}

					mc := []*ent.MessageMediaCreate{}
					for _, v := range m.Media {
						if !v.CanView {
							continue
						}

						m := Ent.MessageMedia.Create().
							SetID(v.ID).
							SetMessageID(messageId).
							SetType(v.Type)

						if v.Src != nil {
							m.SetFull(*v.Src)
						}

						mc = append(mc, m)
					}

					if err := Ent.MessageMedia.CreateBulk(mc...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
						return err
					}

					bar.Add(1)
					nextId = &m.ID
				}
			}

			// Consider a head marker or something?

			bar.Close()

		}

		return nil
	}),
}

func init() {
	scrapeCmd.AddCommand(scrapeSubsCmd)
	scrapeCmd.AddCommand(scrapeMediaPostsCmd)
	scrapeCmd.AddCommand(scrapeMessagesCmd)
	CLI.AddCommand(scrapeCmd)
}
