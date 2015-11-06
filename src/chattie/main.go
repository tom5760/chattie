package main

import (
	"log"
	"strings"

	"github.com/nsf/termbox-go"
)

// App is a simple structure used to hold all of the UI structures, and the
// main loop.
type App struct {
	isRunning bool
	channel   Channel

	inputBox   InputBox
	messageBox MessageBox
}

func main() {
	var err error
	var app App

	app.channel, err = NewMulticastChannel("224.0.0.1:6000")
	if err != nil {
		log.Fatal("Failed to create multicast channel:", err)
		return
	}
	defer app.channel.Close()

	if err := termbox.Init(); err != nil {
		log.Fatal("Failed to initialize termbox:", err)
		return
	}
	defer termbox.Close()

	go app.recv()

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

func (a *App) recv() {
	for a.isRunning {
		message, err := a.channel.Receive()
		if err != nil {
			log.Fatal("Failed to read from channel:", err)
			return
		}

		a.addMessage(message)
		a.draw()
	}
}

func (a *App) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	a.inputBox.Draw()
	a.messageBox.Draw()
	termbox.Flush()
}

func (a *App) onKeyEvent(event termbox.Event) {
	switch event.Key {
	case termbox.KeyCtrlC:
		a.isRunning = false

	case termbox.KeyEnter:
		data := strings.TrimSpace(a.inputBox.Input)
		if len(data) > 0 {
			a.channel.Send(a.addString(a.inputBox.Input))
			a.inputBox.Clear()
		}

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

func (a *App) addMessage(message *Message) {
	a.messageBox.PushMessage(message)
}

func (a *App) addString(data string) *Message {
	message := &Message{
		Author: "chappie",
		Data:   data,
	}
	a.addMessage(message)
	return message
}
