package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

// App is a simple structure used to hold all of the UI structures, and the
// main loop.
type App struct {
	isRunning bool
	inputBox InputBox
}

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatal("Failed to initialize termbox:", err)
	}
	defer termbox.Close()

	var app App
	app.loop()
}

func (a *App) loop() {
	a.isRunning = true

	for a.isRunning {
		a.draw()

		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			a.onKeyEvent(event)
		}
	}
}

func (a *App) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	a.inputBox.Draw()
	termbox.Flush()
}

func (a *App) onKeyEvent(event termbox.Event) {
	switch event.Key {
	case termbox.KeyCtrlC:
		a.isRunning = false

	case termbox.KeyEnter:
		a.inputBox.Clear()

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		a.inputBox.PopRune()

	case termbox.KeySpace:
		a.inputBox.PushRune(' ')

	default:
		if event.Ch != 0 {
			a.inputBox.PushRune(event.Ch)
		}
	}
}
