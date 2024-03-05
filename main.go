package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"pomodoro_timer/internal"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)

	err = g.SetKeybinding("", gocui.KeyCtrlC,  gocui.ModNone, quit); 

	if err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}


func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    if v, err := g.SetView("hello", (maxX/2) - 15 , (maxY/2) - 5 , (maxX/2) + 15,  (maxY/2) + 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

		timerAmount, _ := internal.ParseInputs()
		timeStr := timerAmount.Format("00:00:00")
        fmt.Fprintln(v, timeStr) 

        v.BgColor = gocui.ColorBlack // Changing background color
        v.FgColor = gocui.ColorWhite // Changing font color
        v.SelBgColor = gocui.ColorDefault // Changing selected background color
        v.SelFgColor = gocui.ColorWhite // Changing selected font color
        v.Wrap = true // Enabling text wrapping
        v.Frame = true // Enable a frame around the view
        v.Title = "timer"// Setting a title for the view
        v.Autoscroll = true // Automatically scroll to the bottom when new content is added
        v.Highlight = true // Enable highlighting
		
    }
    return nil
}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
