package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termWidth, termHeight := termbox.Size()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x := 0; x < termWidth; x++ {
		termbox.SetCell(x, 0, ' ', termbox.ColorGreen, termbox.ColorCyan)
		termbox.SetCell(x, termHeight - 1, ' ', termbox.ColorGreen, termbox.ColorCyan)
	}
	for y := 0; y < termHeight; y++ {
		termbox.SetCell(0, y, ' ', termbox.ColorGreen, termbox.ColorCyan)
		termbox.SetCell(termWidth - 1, y, ' ', termbox.ColorGreen, termbox.ColorCyan)
	}

	print_tb((termWidth / 2) - 10, (termHeight / 2) - 5, termbox.ColorWhite, termbox.ColorDefault, "Welcome to the Game!")
	print_tb((termWidth / 2) - 11, (termHeight / 2), termbox.ColorWhite, termbox.ColorDefault, "Press any key to exit.")
	termbox.Flush()
	termbox.PollEvent()
}

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
