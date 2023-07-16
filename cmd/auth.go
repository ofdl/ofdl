package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Obtain and populate Auth Config with Chromium",
	Long: `Obtain and populate Auth Config with Chromium

Once you've configured Chromium, quit your browser and run this command. It will
open a new browser window, and prompt you to login to OnlyFans. Once you've
logged in, press enter to continue. The command will then populate your
configuration with the necessary values.

It was originally developed against Arc, a newer flavor of Chromium which
handles Chrome Profiles differently. If the profile is configured incorrectly,
the new Chromium window will not have your extensions or sessions present.

If you can't get the auth helper to work properly, you can always manually
configure authentication. Follow the instructions at the link below to obtain
the correct values.

https://github.com/DIGITALCRIMINALS/UltimaScraper#running-the-app-locally
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := []chromedp.ExecAllocatorOption{
			chromedp.ExecPath(viper.GetString("chromium.exec")),
			chromedp.UserDataDir(viper.GetString("chromium.profile")),
		}
		allocatorContext, cancel := chromedp.NewExecAllocator(cmd.Context(), opts...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocatorContext)
		defer cancel()

		if err := chromedp.Run(ctx,
			chromedp.Navigate(`https://onlyfans.com/`),
		); err != nil {
			return err
		}

		fmt.Print("Login, then press enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		var userAgent string
		var xBc string
		return chromedp.Run(ctx,
			chromedp.Navigate(`https://onlyfans.com/`),
			chromedp.Evaluate(`navigator.userAgent`, &userAgent),
			chromedp.Evaluate(`localStorage.getItem('bcTokenSha')`, &xBc),

			chromedp.ActionFunc(func(ctx context.Context) error {
				viper.Set("auth.user-agent", userAgent)
				viper.Set("auth.x-bc", xBc)

				c, err := storage.GetCookies().Do(ctx)
				if err != nil {
					return err
				}

				ofc := lo.Filter(c, func(cookie *network.Cookie, i int) bool {
					return cookie.Domain == ".onlyfans.com"
				})

				if ac, ok := lo.Find(ofc, func(cookie *network.Cookie) bool {
					return cookie.Name == "auth_id"
				}); ok {
					viper.Set("auth.user-id", ac.Value)
				}

				cs := strings.Join(lo.Map(ofc, func(cookie *network.Cookie, i int) string {
					return fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
				}), "; ")
				viper.Set("auth.cookie", cs)

				fmt.Println("Writing auth config:")
				spew.Dump(viper.Get("auth"))
				return viper.WriteConfig()
			}),
		)
	},
}

func init() {
	CLI.AddCommand(authCmd)

	authCmd.Flags().String("chromium", "", "Path to chromium executable")
	viper.BindPFlag("chromium.exec", authCmd.Flags().Lookup("chromium"))
	authCmd.Flags().String("profile", "", "Path to chromium userdata")
	viper.BindPFlag("chromium.profile", authCmd.Flags().Lookup("profile"))
}
