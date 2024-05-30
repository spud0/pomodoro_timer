package view

import (
	"fmt"
	"strings"
	"time"
	
	"github.com/spud0/pomodoro_timer/internal"
	tea "github.com/charmbracelet/bubbletea"
)

// Access the model package, for the Pomodoro Object
type model struct {
	model *model.Pomodoro
}

// The application structure
type app struct {
	model model
}

// Creates a new application being given a Pomodoro Object
func NewApp(pom *model.Pomodoro) tea.Model {

	return &app {
		model : model{
			Pomodoro : pom,	
		},  	
	} 
}

// Initializes the app, no cmd function given
func (a app) Init () tea.Cmd {
	return nil
}

// Ticks the timer
func (a app) tick() tea.Cmd {

	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t	
	} 
}

// The Update function, 
func (a app) Update (msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:	
		switch msg.String() {
			case "ctrl+c", "q": 
				return a, tea.Quit
		
			// Start the timer 'r' -> resume
			case "r": 
				a.Pomodoro.Start()
		} 

	case time.Time:
		if a.Pomodoro.Active && a.Pomodoro.RemainingTime() <= 0 {
			a.Pomodoro.Stop()
			fmt.Println("timers done!")
		} 

	}
}

	return a, a.tick
}

func (a app) View () string {
	
	var result strings.Builder

	if a.Pomodoro.Active {
		remTime = a.Pomodoro.RemainingTime()
		mins = int(rem.Minutes()) 
		secs = int(rem.Seconds()) % 60

		fmt.Fprintf(&result, "pom timer: %02d:%02d\n", mins, secs)

		result.WriteString(" [p] : pause, [q] / [esc]: quit")
	} else {

		result.WriteString("pom isn't active...")
		result.WriteString(" [p] : pause, [q] / [esc] : quit")
	}

	return result.String()
}
