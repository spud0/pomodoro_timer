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
    if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.BgColor = gocui.ColorBlue // Changing background color
        v.FgColor = gocui.ColorWhite // Changing font color
        v.SelBgColor = gocui.ColorGreen // Changing selected background color
        v.SelFgColor = gocui.ColorBlack // Changing selected font color
        v.Wrap = true // Enabling text wrapping
        v.Frame = true // Enable a frame around the view
        v.Title = "Hello View" // Setting a title for the view
        v.Autoscroll = true // Automatically scroll to the bottom when new content is added
        v.Highlight = true // Enable highlighting
		
		// Parse CLI
		timerAmount, _ := internal.ParseInputs()
        fmt.Fprintln(v, timerAmount.Format())
    }
    return nil
}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
