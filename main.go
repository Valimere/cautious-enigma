package main

import (
	"flag"
	"fmt"
	"github.com/Valimere/cautious-enigma/shapes"
	"os"
	"strconv"
	"strings"
)

const (
	MaxWidth  = 10
	MaxHeight = 100 // as described in the documentation
)

func main() {
	inputFile := flag.String("i", "input.txt", "Input File")
	debugFlag := flag.Bool("d", false, "Debug Mode")
	flag.Parse()

	// Initialize the Tetris grid
	grid := make([][]int, MaxHeight)
	for i := range grid {
		grid[i] = make([]int, MaxWidth)
	}

	// Read the input file
	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	if *debugFlag {
		fmt.Println("Input File:", *inputFile)
		printInput(lines)
	}
	for i, line := range lines {
		if line == "" {
			continue
		}

		// Reset grid for each new line of input
		resetGrid(grid)

		// Processing logic for each Tetromino described in the line
		processLine(line, grid, *debugFlag)

		// Calculate and print the resultant grid height
		fmt.Printf("Case %d Resulting Height:%v\n", i+1, shapes.CalculateHeight(grid, *debugFlag))
		fmt.Println() // Print a newline for separation between outputs
	}
}
func resetGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
}

func processLine(line string, grid [][]int, debugFlag bool) {
	parts := strings.Split(line, ",")
	for _, part := range parts {
		isValidInput(part, MaxWidth)
		tType := part[:len(part)-1]
		xPos, err := strconv.Atoi(part[len(part)-1:])
		if err != nil {
			fmt.Println("Invalid x position in:", part, "Error:", err)
			continue
		}
		tetromino, err := shapes.NewTetromino(tType, xPos, 0) // Start at the top of the grid
		if err != nil {
			fmt.Errorf("error creating tetromino: %v", err)
		}
		shapes.DropTetromino(&tetromino, grid, debugFlag)
	}
}

func isValidInput(part string, gridWidth int) bool {
	if len(part) < 2 {
		return false
	}
	xPos, err := strconv.Atoi(part[len(part)-1:])
	if err != nil || xPos < 0 || xPos >= gridWidth {
		return false
	}
	return true
}

func printInput(lines []string) {
	fmt.Println("Debug Mode Enabled")
	for i, line := range lines {
		if line == "" {
			continue
		}

		fmt.Printf("Run %d:\n", i+1)
		parts := strings.Split(line, ",")
		for _, part := range parts {
			if len(part) < 2 {
				fmt.Println("Invalid Tetromino data:", part)
				continue
			}
			// Extract tetromino type and x position
			tType := part[:len(part)-1]
			xPos, err := strconv.Atoi(part[len(part)-1:])
			if err != nil {
				fmt.Println("Invalid x position in:", part, "Error:", err)
				continue
			}
			// Create a new tetromino and print its details
			tetromino, err := shapes.NewTetromino(tType, xPos, 0) // Assuming y is always 0 as input does not specify y
			if err != nil {
				fmt.Errorf("error creating tetromino: %v", err)
			}
			fmt.Println(&tetromino)
		}
		fmt.Println()
	}
}
