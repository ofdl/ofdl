package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ofdl/ofdl/cmd/gui"
	"github.com/spf13/cobra"
)

var subsCmd = &cobra.Command{
	Use:     "subscriptions",
	Aliases: []string{"subs"},
	Short:   "Manage Tracked Subscriptions",
	Long: `Manage Tracked Subscriptions

Subscriptions expired? Disenchanted with a creator? Manage which subscriptions
you're tracking here.
	`,
	PersistentPreRunE: UseOFDL,
	RunE: func(cmd *cobra.Command, args []string) error {
		var g *gui.SubsGUI
		if err := App.Resolve(&g); err != nil {
			return err
		}

		p := tea.NewProgram(g)
		_, err := p.Run()
		return err
	},
}

func init() {
	CLI.AddCommand(subsCmd)
}
