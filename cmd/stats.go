package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:               "stats",
	Short:             "Print database statistics",
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := OFDL.Query.Subscription
		subCount, err := s.Count()
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count: %d\n", subCount)

		uosCount, err := s.Where(s.OrganizedAt.IsNull()).Count()
		if err != nil {
			return err
		}
		fmt.Printf("Subscription Count (Unorganized): %d\n", uosCount)

		mpCount, err := OFDL.Query.Post.Count()
		if err != nil {
			return err
		}
		fmt.Printf("Post Count: %d\n", mpCount)

		m := OFDL.Query.Media
		mCount, err := m.Count()
		if err != nil {
			return err
		}
		fmt.Printf("Media Count: %d\n", mCount)

		udmCount, err := m.Where(m.DownloadedAt.IsNull()).Count()
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Undownloaded): %d\n", udmCount)

		uomCount, err := m.Where(m.OrganizedAt.IsNull()).Count()
		if err != nil {
			return err
		}
		fmt.Printf("Media Count (Unorganized): %d\n", uomCount)

		mm := OFDL.Query.MessageMedia
		mmCount, err := mm.Count()
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count: %d\n", mmCount)

		udmmCount, err := mm.Where(mm.DownloadedAt.IsNull()).Count()
		if err != nil {
			return err
		}
		fmt.Printf("Message Media Count (Undownloaded): %d\n", udmmCount)

		uommCount, err := mm.Where(mm.OrganizedAt.IsNull()).Count()
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
