package gui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ofdl/ofdl/model"
	"gorm.io/gorm"
)

type SubsGUI struct {
	db     *gorm.DB
	subs   []*model.Subscription
	cursor int
	msg    string
}

func NewSubsGui(db *gorm.DB, subs []*model.Subscription) *SubsGUI {
	return &SubsGUI{
		db:     db,
		subs:   subs,
		cursor: 0,
	}
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
			sub.Enabled = !sub.Enabled
			if err := s.db.Save(sub).Error; err != nil {
				return s, tea.Quit
			}

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
