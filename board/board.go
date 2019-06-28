package board

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Counter - Represent piece in connect 4
type Counter struct {
	x      int32
	y      int32
	colour string
	radius int32
}

// Board ... Connect 4 board 7x board (most common type)
type Board struct {
	grid [6][7]Counter
	x    int16
	y    int16
	Turn string
	Win  bool
}

//New ... Initialiser
func New() *Board {
	board := &Board{}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	randNumber := random.Intn(2) + 1

	if randNumber == 1 {
		board.Turn = "red"
	} else {
		board.Turn = "yellow"
	}

	board.setupBoard()
	return board
}

func (b *Board) setupBoard() {
	radius := 35
	startx := 120
	starty := 100
	offsetx := 60
	offsety := 50
	b.x = 50
	b.y = 40
	for row := range b.grid {
		for column := range b.grid[row] {
			c := Counter{}
			c.x = int32(startx + (column * radius) + offsetx*column)
			c.y = int32(starty + (row * radius) + offsety*row)
			c.colour = "black"
			c.radius = int32(radius)
			b.grid[row][column] = c
		}
	}
}

// DrawBoard ... Drawing method for board
func (b *Board) DrawBoard(renderer *sdl.Renderer) {
	filledRectColor(renderer, b.x, b.y, 700, 550, sdl.Color{0, 0, 255, 255})
	for row := range b.grid {
		for _, counter := range b.grid[row] {
			if counter.colour == "red" {
				gfx.FilledCircleColor(renderer, counter.x, counter.y, counter.radius, sdl.Color{255, 0, 0, 255})
			} else if counter.colour == "yellow" {
				gfx.FilledCircleColor(renderer, counter.x, counter.y, counter.radius, sdl.Color{245, 229, 27, 255})
			} else {
				gfx.FilledCircleColor(renderer, counter.x, counter.y, counter.radius, sdl.Color{0, 0, 0, 255})
			}
		}
	}
}

// ColumnCheck ... Collision box for mouse click
func (b *Board) ColumnCheck(x int32, y int32) {
	for row := range b.grid {
		for column, counter := range b.grid[row] {
			if x > counter.x-counter.radius && x < counter.x+(counter.radius) &&
				y > counter.y-counter.radius && y < counter.y+(counter.radius) {
				if b.placeCounter(column, b.Turn) {
					b.Win = b.ruleCheck(b.Turn)
					if b.Win {
						return
					}
					if b.Turn == "red" {
						b.Turn = "yellow"
					} else {
						b.Turn = "red"
					}
				}

			}
		}
	}
}

func (b *Board) ruleCheck(colour string) bool {
	return b.verticalCheck(colour) || b.horizontalCheck(colour) || b.diagonalCheck(colour)
}

func (b *Board) placeCounter(column int, colour string) bool {
	for row := len(b.grid) - 1; row >= 0; row-- {
		if b.grid[row][column].colour == "black" {
			b.grid[row][column].colour = colour
			return true
		}
	}
	return false
}

// DrawCollisionBoxes ... Debug draw collision boxes
func (b *Board) DrawCollisionBoxes(renderer *sdl.Renderer) {
	for row := range b.grid {
		for _, counter := range b.grid[row] {
			gfx.RectangleColor(renderer, counter.x-counter.radius, counter.y-counter.radius, counter.x+(counter.radius), counter.y+(counter.radius), sdl.Color{0, 255, 0, 255})
		}
	}
}

func filledRectColor(renderer *sdl.Renderer, left int16, top int16, width int16, height int16, colour sdl.Color) {
	var vx, vy = make([]int16, 4), make([]int16, 4)
	// top left
	vx[0] = int16(left)
	vy[0] = int16(top)
	//top right
	vx[1] = int16(left + width)
	vy[1] = int16(top)
	//bottom left
	vx[3] = int16(left)
	vy[3] = int16(top + height)
	//bottom right
	vx[2] = int16(left + width)
	vy[2] = int16(top + height)

	gfx.FilledPolygonColor(renderer, vx, vy, colour)
}

func (b *Board) verticalCheck(colour string) bool {
	maxRows := len(b.grid)
	//maxColumns := len(b.grid[0])

	var maxRow int
	//var maxColumn int

	for row := range b.grid {
		for column := range b.grid[row] {
			maxRow = row + 3
			//maxColumn = column + 3
			if maxRow >= maxRows {
				continue
			}

			counter1 := b.grid[row][column]
			counter2 := b.grid[row+1][column]
			counter3 := b.grid[row+2][column]
			counter4 := b.grid[row+3][column]

			if counter1.colour == colour && counter2.colour == colour &&
				counter3.colour == colour && counter4.colour == colour {
				return true
			}
		}
	}

	return false
}

func (b *Board) horizontalCheck(colour string) bool {
	//maxRows := len(b.grid)
	maxColumns := len(b.grid[0])

	//var maxRow int
	var maxColumn int

	for row := range b.grid {
		for column := range b.grid[row] {
			//maxRow = row + 3
			maxColumn = column + 3
			if maxColumn >= maxColumns {
				continue
			}

			counter1 := b.grid[row][column]
			counter2 := b.grid[row][column+1]
			counter3 := b.grid[row][column+2]
			counter4 := b.grid[row][column+3]

			if counter1.colour == colour && counter2.colour == colour &&
				counter3.colour == colour && counter4.colour == colour {
				return true
			}
		}
	}

	return false
}

func (b *Board) diagonalCheck(colour string) bool {
	maxRows := len(b.grid)
	maxColumns := len(b.grid[0])

	var maxRow int
	var maxColumn int
	var minColumn int

	for row := range b.grid {
		for column := range b.grid[row] {
			maxRow = row + 3
			maxColumn = column + 3
			if maxColumn < maxColumns && maxRow < maxRows {
				rdc1 := b.grid[row][column]
				rdc2 := b.grid[row+1][column+1]
				rdc3 := b.grid[row+2][column+2]
				rdc4 := b.grid[row+3][column+3]

				if rdc1.colour == colour && rdc2.colour == colour &&
					rdc3.colour == colour && rdc4.colour == colour {
					return true
				}
			}

			minColumn = column - 3
			if minColumn > 0 && maxRow < maxRows {
				ldc1 := b.grid[row][column]
				ldc2 := b.grid[row+1][column-1]
				ldc3 := b.grid[row+2][column-2]
				ldc4 := b.grid[row+3][column-3]

				if ldc1.colour == colour && ldc2.colour == colour &&
					ldc3.colour == colour && ldc4.colour == colour {
					return true
				}
			}
		}
	}
	return false
}
