package main

import (
	"flag"
	"time"

	"github.com/spud0/pomodoro_timer/internal"
	tea "github.com/charmbracelet/bubbletea/"
)

// Contains the entry point
func main () {
	
	durFlag := flag.Duration("dur", 2 * time.Minute, "duration of the timer")
	breakFlag  := flag.Duration("break", 1 * time.Minute, "break time during timer")

	flag.Parse()
	
	pom := model.NewPomodoro(*durFlag, *breakFlag)
	session := tea.NewProgram(view.NewApp(pom))

	if err := session.Start(); err != nil {
		panic(err)
	} 
}
