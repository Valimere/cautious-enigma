package shapes

import (
	"fmt"
	"strings"
)

type Tetromino struct {
	shape [][]int
	x     int // x position on the board
	y     int // y position on the board
}

// NewTetromino creates a new Tetrimono of the specified  type at the given position
func NewTetromino(tetriminoType string, x, y int) Tetromino {
	var shape [][]int
	switch tetriminoType {
	case "Q": // Square
		shape = [][]int{
			{1, 1},
			{1, 1},
		}
	case "Z":
		shape = [][]int{
			{1, 1, 0},
			{0, 1, 1},
		}
	case "S":
		shape = [][]int{
			{0, 1, 1},
			{1, 1, 0},
		}
	case "T":
		shape = [][]int{
			{1, 1, 1},
			{0, 1, 0},
		}
	case "I":
		shape = [][]int{
			{1, 0},
			{1, 0},
			{1, 0},
			{1, 0},
		}
	case "L":
		shape = [][]int{
			{1, 0},
			{1, 0},
			{1, 1},
		}
	case "J":
		shape = [][]int{
			{0, 1},
			{0, 1},
			{1, 1},
		}
	default:
		fmt.Errorf("unknown tetrimino type: %s", tetriminoType)
		return Tetromino{}
	}
	return Tetromino{
		shape: shape,
		x:     x,
		y:     y,
	}
}

func (t *Tetromino) String() string {
	grid := make([][]string, 4)
	for i := range grid {
		grid[i] = make([]string, 4)
		for j := range grid[i] {
			grid[i][j] = "." // Use dot for visualization
		}
	}

	// Adjust for shapes larger than 2x2 but not full 4x4
	offsetX, offsetY := max((4-len(t.shape[0]))/2, 0), max((4-len(t.shape))/2, 0)

	// Place the tetromino shape in the grid
	for i, row := range t.shape {
		for j, val := range row {
			if val != 0 {
				if offsetY+i < 4 && offsetX+j < 4 { // Ensure it doesn't go out of bounds
					grid[offsetY+i][offsetX+j] = "#"
				}
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Tetromino: x=%d, y=%d\n", t.x, t.y))
	for _, row := range grid {
		for _, cell := range row {
			sb.WriteString(cell)
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
