package main

import (
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
