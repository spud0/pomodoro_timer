package view

import (
	"fmt"
	"strings"
	"time"
	
	"github.com/spud0/pomodoro_timer/internal/model"
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

// The Update function, 
func (a app) Update (msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:	
		switch msg.String() {
			case "ctrl+c", "q": 
				return a, tea.Quit
			
			// case : return a, nil
		} 

	case tea.TickMsg:
		if a.Pomodoro.Active  {
			return a, tea.Tick(time.Second, 
				func(time.Time) tea.Msg {
					return tea.TickMsg{}	
			})
		} 

	}
}

	return a, nil
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
