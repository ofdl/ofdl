package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ofdl/ofdl/cmd/gui"
	"github.com/ofdl/ofdl/ofdl"
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
	RunE: ofdl.RunE(func(g *gui.SubsGUI) error {
		p := tea.NewProgram(g)
		_, err := p.Run()
		return err
	}),
}

func init() {
	CLI.AddCommand(subsCmd)
}
