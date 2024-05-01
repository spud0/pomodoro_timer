package view

import (
	"fmt"
	"time"
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
	
	switch msg.(type) {
		case tea.KeyMsg: 


	} 

	return a, nil

	} 

}
