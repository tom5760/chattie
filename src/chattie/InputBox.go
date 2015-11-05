package main

import (
	"github.com/nsf/termbox-go"
)

// Leave 4 cells around the actual text of the input box.
const inputPadding = 4

// InputBox represents the box where the user inputs chat messages.
type InputBox struct {
	Input string
}

// Draw redraws the input box.
func (b *InputBox) Draw() {
	width, height := termbox.Size()

	// Draw a border around the screen.
	drawBox(height-3, 0, width-1, 3)

	// Draw our current input.
	visibleInput := b.getVisibleInput(width)
	drawText(2, height-2, visibleInput)

	// Set the cursor to the end of the input.
	termbox.SetCursor(2 + len(visibleInput), height-2)
}

// PushRune appends the rune r to the end of the input string.
func (b *InputBox) PushRune(r rune) {
	b.Input += string(r)
}

// PopRune removes the last rune from the input string.
func (b *InputBox) PopRune() {
	if len(b.Input) > 0 {
		b.Input = b.Input[:len(b.Input)-1]
	}
}

// Clear resets the input string to the empty string.
func (b *InputBox) Clear() {
	b.Input = ""
}

func (b *InputBox) getVisibleInput(width int) string {
	inputWidth := width - inputPadding
	inputLength := len(b.Input)
	if inputLength > inputWidth {
		return b.Input[inputLength - inputWidth:]
	} else {
		return b.Input
	}
}
