package main

import (
	"fmt"

	"github.com/soultoyo/connect4/board"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Todo
// Document code
// Upload onto github - public

var renderer *sdl.Renderer
var winWidth, winHeight int32 = 800, 600
var running = false

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Connect 4", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		panic(err)
	}
	renderer.Clear()
	defer renderer.Destroy()

	b := board.New()

	running = true
	for running {
		if b.Win {
			result, err := winMessageBox()
			if result && err == nil {
				b = board.New()
			} else if result && err == nil {
				running = false
			} else {
				fmt.Println(err)
				running = false
			}
		}
		input(b)
		draw(renderer, b)
	}
}

func draw(renderer *sdl.Renderer, b *board.Board) {
	renderer.SetDrawColor(100, 100, 100, 255)
	renderer.Clear()

	if b.Turn == "yellow" {
		gfx.FilledCircleColor(renderer, 25, 25, 15, sdl.Color{245, 229, 27, 255})
	} else {
		gfx.FilledCircleColor(renderer, 25, 25, 15, sdl.Color{255, 0, 0, 255})
	}

	b.DrawBoard(renderer)
	renderer.Present()
}

func input(b *board.Board) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
			break
		case *sdl.MouseButtonEvent:
			if b.Win {
				break
			}

			// checking if left mouse button has been clicked
			if t.State == 1 && t.Button == 1 {
				b.ColumnCheck(t.X, t.Y)
			}
			break
		}
	}
}

func winMessageBox() (bool, error) {
	buttons := []sdl.MessageBoxButtonData{
		{sdl.MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, 0, "No"},
		{sdl.MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "Yes"},
	}

	// background, text, button border, button background, button selected
	colorScheme := sdl.MessageBoxColorScheme{
		Colors: [5]sdl.MessageBoxColor{
			sdl.MessageBoxColor{44, 62, 80},
			sdl.MessageBoxColor{189, 195, 199},
			sdl.MessageBoxColor{41, 128, 185},
			sdl.MessageBoxColor{149, 165, 166},
			sdl.MessageBoxColor{127, 140, 141},
		},
	}

	messageboxdata := sdl.MessageBoxData{
		sdl.MESSAGEBOX_INFORMATION,
		nil,
		"Play again",
		"Do you want to play again?",
		buttons,
		&colorScheme,
	}

	var buttonid int32
	var err error
	if buttonid, err = sdl.ShowMessageBox(&messageboxdata); err != nil {
		fmt.Println("error displaying message box")
		return false, err
	}

	if buttonid == -1 {
		fmt.Println("no selection")
	} else if buttonid == 1 {
		return true, nil
	} else {
		return false, nil
	}
	return false, nil
}
