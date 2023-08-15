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

func NewSubsGui(ctx context.Context, ent *ent.Client) (*SubsGUI, error) {
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

func (s SubsGUI) Init() tea.Cmd {
	return nil
}

func (s SubsGUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < len(s.subs)-1 {
				s.cursor++
			}
		case "enter", " ":
			sub := s.subs[s.cursor]
			sub, err := sub.Update().SetEnabled(!sub.Enabled).Save(s.ctx)
			if err != nil {
				return s, tea.Quit
			}
			s.subs[s.cursor] = sub

			verb := "enabled"
			if !sub.Enabled {
				verb = "disabled"
			}
			s.msg = fmt.Sprintf("%s (%s) %s", sub.Name, sub.Username, verb)
		}
	}

	return s, nil
}

func (s SubsGUI) View() string {
	out := "Select which subscriptions you'd like to scrape:\n\n"

	for i, choice := range s.subs {
		cursor := " "
		if s.cursor == i {
			cursor = ">"
		}

		checked := " "
		if choice.Enabled {
			checked = "X"
		}

		out += fmt.Sprintf("%s [%s] %s (%s)\n", cursor, checked, choice.Name, choice.Username)
	}

	out += fmt.Sprintf("\n%s\nPress q to quit.\n", s.msg)
	return out
}
