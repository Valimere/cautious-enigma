package main

import (
	"flag"
	"fmt"
	"github.com/Valimere/cautious-enigma/shapes"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("i", "input.txt", "Input File")
	flag.Parse()

	// Read the input file
	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
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
			tetromino := shapes.NewTetromino(tType, xPos, 0) // Assuming y is always 0 as input does not specify y
			fmt.Println(&tetromino)
		}
		fmt.Println()
	}
}
