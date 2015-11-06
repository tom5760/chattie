package main

import (
	"github.com/nsf/termbox-go"
)

func drawBox(top, left, width, height int) {
	drawLineHorizontal(left, left+width, top)
	drawLineHorizontal(left, left+width, top+height-1)

	drawLineVertical(left, top, top+height)
	drawLineVertical(left+width, top, top+height-1)

	// Draw rounded corners
	termbox.SetCell(left, top, '┌', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(left+width, top, '┐', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(left+width, top+height-1, '┘', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(left, top+height-1, '└', termbox.ColorDefault, termbox.ColorDefault)
}

func drawLineHorizontal(x1, x2, y int) {
	for x := x1; x < x2; x++ {
		termbox.SetCell(x, y, '─', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawLineVertical(x, y1, y2 int) {
	for y := y1; y < y2; y++ {
		termbox.SetCell(x, y, '│', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawText(x, y int, s string) {
	for i, c := range s {
		termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}
