package gui

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ofdl/ofdl/ent"
)

type SubsGUI struct {
	ctx    context.Context
	ent    *ent.Client
	subs   []*ent.Subscription
	cursor int
	msg    string
}

func NewSubsGUI(ctx context.Context, ent *ent.Client) (*SubsGUI, error) {
	subs, err := ent.Subscription.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return &SubsGUI{
		ctx:    ctx,
		ent:    ent,
		subs:   subs,
		cursor: 0,
	}, nil
}

func (SubsGUI) Init() tea.Cmd {
	return nil
}

func (g SubsGUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "up", "k":
			if g.cursor > 0 {
				g.cursor--
			}
		case "down", "j":
			if g.cursor < len(g.subs)-1 {
				g.cursor++
			}
		case "enter", " ":
			sub := g.subs[g.cursor]
			sub, err := sub.Update().SetEnabled(!sub.Enabled).Save(g.ctx)
			if err != nil {
				return g, tea.Quit
			}
			g.subs[g.cursor] = sub

			verb := "enabled"
			if !sub.Enabled {
				verb = "disabled"
			}
			g.msg = fmt.Sprintf("%s (%s) %s", sub.Name, sub.Username, verb)
		}
	}

	return g, nil
}

func (g SubsGUI) View() string {
	out := "Select which subscriptions you'd like to scrape:\n\n"

	for i, choice := range g.subs {
		cursor := " "
		if g.cursor == i {
			cursor = ">"
		}

		checked := " "
		if choice.Enabled {
			checked = "X"
		}

		out += fmt.Sprintf("%s [%s] %s (%s)\n", cursor, checked, choice.Name, choice.Username)
	}

	out += fmt.Sprintf("\n%s\nPress q to quit.\n", g.msg)
	return out
}
