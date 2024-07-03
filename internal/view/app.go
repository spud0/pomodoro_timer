package view

import (
    "fmt"
    "strings"
    "time"
    
    "github.com/spud0/pomodoro_timer/internal/model"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    Pomodoro *model.Pomodoro
}

func NewApp(pom *model.Pomodoro) tea.Model {
    return Model{
        Pomodoro: pom,
    }
}

func (m Model) Init() tea.Cmd {
    return m.tick()
}

func (m Model) tick() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return t
    })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "esc":
            return m, tea.Quit
        case "r":
            m.Pomodoro.Start()
        case "p":
            m.Pomodoro.Stop()
        }
    case time.Time:
        if m.Pomodoro.Active && m.Pomodoro.RemainingTime() <= 0 {
            m.Pomodoro.Stop()
            fmt.Println("Timer's done!")
        }
    }

    return m, m.tick()
}

func (m Model) View() string {
    var result strings.Builder

    if m.Pomodoro.Active {
        remTime := m.Pomodoro.RemainingTime()
        mins := int(remTime.Minutes())
        secs := int(remTime.Seconds()) % 60

        fmt.Fprintf(&result, "Pomodoro timer: %02d:%02d\n", mins, secs)
        result.WriteString("[p]: pause, [q]/[esc]: quit")
    } else {
        result.WriteString("Pomodoro isn't active...\n")
        result.WriteString("[r]: resume, [q]/[esc]: quit")
    }

    return result.String()
}

