package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type MessageBox struct {
	Messages []*Message
}

// Draw redraws the chat box.
func (b *MessageBox) Draw() {
	width, height := termbox.Size()
	boxHeight := height - 3

	// Draw a border around the box
	drawBox(0, 0, width-1, boxHeight)

	// Subtract 2 from height to account for top and bottom borders
	visibleMessages := b.getVisibleMessages(boxHeight - 2)

	for i, message := range visibleMessages {
		drawText(2, 1+i, fmt.Sprintf("%s: %s", message.Author, message.Data))
	}
}

func (b *MessageBox) PushMessage(message *Message) {
	b.Messages = append(b.Messages, message)
}

func (b *MessageBox) getVisibleMessages(height int) []*Message {
	messageCount := len(b.Messages)
	if messageCount > height {
		return b.Messages[messageCount-height:]
	} else {
		return b.Messages
	}
}
