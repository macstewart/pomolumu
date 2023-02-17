package ui

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/macstewart/pomolumu/pkg"
)

var (
	workMinutes = 25
	breaks      = []int{5, 5, 5, 15}
)

type model struct {
	keymap     keymap
	timer      *pkg.Timer
	breakIndex int
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tickCmd())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.timer.IsTimedOut() {
		m = m.resetTimer()
	}
	m.keymap.start.SetEnabled(!m.timer.IsRunning())
	m.keymap.stop.SetEnabled(m.timer.IsRunning())
	switch msg := msg.(type) {
	case time.Time:
		return m, tickCmd()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.reset):
			m.timer.Reset()
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			m.timer.Toggle()
		}
	}
	return m, nil
}

func (m model) resetTimer() model {
	if m.timer.TimerType == pkg.Work {
		m.timer = pkg.BreakTimer(breaks[m.breakIndex])
		m.timer.Start()
		m.breakIndex++
		if m.breakIndex >= len(breaks) {
			m.breakIndex = 0
		}
	} else {
		m.timer = pkg.WorkTimer(workMinutes)
	}
	return m
}

func (m model) View() string {
	content := m.timer.Render()
	return content
}

func Run(args []string) {
	if len(args) > 0 {
		if argMinutes, err := strconv.Atoi(args[0]); err == nil {
			workMinutes = argMinutes
		}
	}
	m := model{
		timer: pkg.WorkTimer(workMinutes),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys(" "),
				// key.WithHelp("<space>", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys(" "),
				// key.WithHelp("<space>", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				// key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				// key.WithHelp("q", "quit"),
			),
		},
	}
	m.timer.Start()
	m.keymap.start.SetEnabled(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return time.Time(t)
	})
}
