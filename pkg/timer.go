package pkg

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/crypto/ssh/terminal"
)

type TimerType int

const (
	Work TimerType = iota
	Break
)

var (
	breakFrames      = []string{"Z  ", " z ", "  z", "   "}
	workFrames       = []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"}
	width, height, _ = terminal.GetSize(0)
	baseStyle        = lipgloss.NewStyle().
				Width(width-2).
				Height(height-2).
				Bold(true).
				Align(lipgloss.Center, lipgloss.Center)
	workStyle = baseStyle.Copy().
			BorderForeground(lipgloss.Color("2")).
			BorderStyle(lipgloss.ThickBorder()).
			Foreground(lipgloss.Color("2"))
	breakStyle = baseStyle.Copy().
			BorderForeground(lipgloss.Color("3")).
			BorderStyle(lipgloss.ThickBorder()).
			Foreground(lipgloss.Color("3"))
	pauseStyle = baseStyle.Copy().
			BorderStyle(lipgloss.HiddenBorder()).
			Foreground(lipgloss.Color("7"))
)

type Timer struct {
	startTime        time.Time
	duration         time.Duration
	originalDuration time.Duration
	TimerType        TimerType
	frames           []string
	frameIndex       int
	running          bool
	style            lipgloss.Style
}

func WorkTimer(minutes int) *Timer {
	duration := time.Duration(minutes) * time.Minute
	return &Timer{
		startTime:        time.Time{},
		duration:         duration,
		originalDuration: duration,
		TimerType:        Work,
		frameIndex:       0,
		frames:           workFrames,
		running:          false,
		style:            workStyle,
	}
}

func BreakTimer(minutes int) *Timer {
	duration := time.Duration(minutes) * time.Minute
	return &Timer{
		startTime:        time.Time{},
		duration:         duration,
		originalDuration: duration,
		TimerType:        Break,
		frameIndex:       0,
		frames:           breakFrames,
		running:          false,
		style:            breakStyle,
	}
}

func (t *Timer) Start() {
	if !t.running {
		t.running = true
		t.startTime = time.Now()
	}
}

func (t *Timer) Stop() {
	if t.running {
		t.running = false
		t.duration = time.Until(t.startTime.Add(t.duration))
	}
}

func (t *Timer) Toggle() bool {
	if t.running {
		t.Stop()
	} else {
		t.Start()
	}
	return t.running
}

func (t *Timer) Reset() {
	t.startTime = time.Now()
	t.duration = t.originalDuration
	t.frameIndex = 0
}


func (t *Timer) Render() string {
	style := t.style
	if !t.IsRunning() {
		style = pauseStyle
	}
	return style.Render(fmt.Sprintf("%3s %4s", t.getFrame(), t.TimeLeft()))
}

func (t *Timer) TimeLeft() time.Duration {
	var timeleft time.Duration
	if t.running {
		timeleft = time.Until(t.startTime.Add(t.duration))
	} else {
		timeleft = t.duration
	}
	if timeleft < 0 {
		timeleft = 0
	}
	return timeleft.Round(time.Second)
}

func (t *Timer) IsTimedOut() bool {
	return t.TimeLeft() <= 0
}

func (t *Timer) IsRunning() bool {
	return t.running
}


func (t *Timer) getFrame() (frame string) {
	frame = t.frames[t.frameIndex]
	if t.running {
		t.frameIndex++
		if t.frameIndex >= len(t.frames) {
			t.frameIndex = 0
		}
	}
	return
}
