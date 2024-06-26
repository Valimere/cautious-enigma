package shapes

import (
	"fmt"
	"strings"
)

type Shape [][]int

// Map for storing Tetrimino shapes
var shapeMap = map[string]Shape{
	"Q": {
		{1, 1},
		{1, 1},
	},
	"Z": {
		{1, 1, 0},
		{0, 1, 1},
	},
	"S": {
		{0, 1, 1},
		{1, 1, 0},
	},
	"T": {
		{1, 1, 1},
		{0, 1, 0},
	},
	"I": {
		{1, 1, 1, 1},
	},
	"L": {
		{1, 0},
		{1, 0},
		{1, 1},
	},
	"J": {
		{0, 1},
		{0, 1},
		{1, 1},
	},
}

type Tetromino struct {
	Shape Shape
	X     int // X position on the board
	Y     int // Y position on the board
}

func NewTetromino(tetriminoType string, x, y int) (Tetromino, error) {
	shape, exists := shapeMap[tetriminoType]
	if !exists {
		return Tetromino{}, fmt.Errorf("unknown tetrimino type: %s", tetriminoType)
	}
	return Tetromino{Shape: shape, X: x, Y: y}, nil
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
	offsetX, offsetY := max((4-len(t.Shape[0]))/2, 0), max((4-len(t.Shape))/2, 0)

	// place the tetromino Shape in the grid
	for i, row := range t.Shape {
		for j, val := range row {
			if val != 0 {
				if offsetY+i < 4 && offsetX+j < 4 { // Ensure it doesn't go out of bounds
					grid[offsetY+i][offsetX+j] = "#"
				}
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Tetromino: X=%d, Y=%d\n", t.X, t.Y))
	for _, row := range grid {
		for _, cell := range row {
			sb.WriteString(cell)
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *Tetromino) Drop(grid [][]int, debugFlag bool) {
	yPos := 0 // Start from the top of the grid
	for t.canMoveDown(yPos, grid) {
		yPos++ // Increment yPos until the Tetromino can no longer move down
	}
	//yPos--                        // Subtract one because the loop exits when Tetromino can no longer move down
	t.place(yPos, grid, debugFlag) // Now place the Tetromino at the calculated position
	if debugFlag {
		printGrid(grid)
	}
}

func (t *Tetromino) canMoveDown(yPos int, grid [][]int) bool {
	// Check if moving down would cause collision with bottom or another Tetromino
	for i, row := range t.Shape {
		for j, val := range row {
			if val != 0 { // Check only occupied cells of the Tetromino
				newY := yPos + i + 1
				x := t.X + j
				if newY >= len(grid) || grid[newY][x] != 0 {
					return false // Collision or out of bounds
				}
			}
		}
	}
	return true
}

func (t *Tetromino) place(yPos int, grid [][]int, debugFlag bool) {
	// place the Tetromino in the grid at the specified y position
	for i, row := range t.Shape {
		for j, val := range row {
			newX := t.X + j
			newY := yPos + i
			if val != 0 && newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
				grid[newY][newX] = 1 // Mark the grid cell as occupied
			}
		}
	}
	clearFullRows(grid, debugFlag)
}

func CalculateHeight(grid [][]int, debugFlag bool) int {
	lastFilledRow := -1 // Start with -1 to indicate no rows filled yet.

	// Iterate from bottom to top of the grid.
	for y := len(grid) - 1; y >= 0; y-- {
		rowFilled := false
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != 0 {
				rowFilled = true
				break // Stop checking this row as soon as one filled cell is found.
			}
		}
		if rowFilled {
			lastFilledRow = y // Update the last filled row index.
		} else {
			// Stop if we encounter an empty row after finding a filled row.
			if lastFilledRow != -1 {
				break
			}
		}
	}
	if debugFlag {
		fmt.Println("Grid during height calculation:")
		printGrid(grid)
	}

	// Calculate the height based on the last filled row.
	// This returns the number of filled rows from the bottom to the first empty row encountered from below.
	if lastFilledRow == -1 {
		return 0 // If no rows are filled.
	} else {
		return len(grid) - lastFilledRow
	}
}

func clearFullRows(grid [][]int, debugFlag bool) {
	if debugFlag {
		printGrid(grid)
	}
	rowCount := len(grid)
	colCount := len(grid[0])

	for i := 0; i < rowCount; i++ {
		full := true
		for j := 0; j < colCount; j++ {
			if grid[i][j] == 0 {
				full = false
				break
			}
		}
		if full {
			if debugFlag {
				fmt.Printf("Clearing row: %d \n", rowCount-i) // Updated print statement
			}
			// Move all rows above this row down by one
			for k := i; k > 0; k-- {
				copy(grid[k], grid[k-1])
			}
			// Clear the top row
			for j := 0; j < colCount; j++ {
				grid[0][j] = 0
			}
			i-- // Recheck this row index since it now has new content
		}
	}
}

// printGrid prints the grid to the console where each cell is either filled or empty.
// Empty cells are represented by '.' and filled cells by '#'.
func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		hasFilledCell := false // Track if the current row has any filled cells
		rowOutput := ""        // Build the output for the current row

		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				rowOutput += ". " // Add a dot for empty cells
			} else {
				rowOutput += "# " // Add a hash for filled cells
				hasFilledCell = true
			}
		}

		if hasFilledCell { // Only print the row if it has at least one filled cell
			fmt.Println(rowOutput)
		}
	}
	fmt.Println("---------------------------") // New line separator between grid prints
}
