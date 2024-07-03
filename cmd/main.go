package main

import (
    "time"

    "github.com/spud0/pomodoro_timer/internal/model"
    "github.com/spud0/pomodoro_timer/internal/view"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    pom := model.NewPom(25*time.Minute, 5*time.Minute)
    app := view.NewApp(pom)
    p := tea.NewProgram(app)
    if err := p.Start(); err != nil {
        panic(err)
    }
}

