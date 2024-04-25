package shapes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTetromino(t *testing.T) {
	tests := []struct {
		name          string
		tetriminoType string
		x             int
		y             int
		want          Tetromino
		wantErr       bool
	}{
		{
			name:          "Q type",
			tetriminoType: "Q",
			x:             1,
			y:             1,
			want: Tetromino{
				Shape: [][]int{
					{1, 1},
					{1, 1},
				},
				X: 1,
				Y: 1,
			},
			wantErr: false,
		},
		{
			name:          "S type",
			tetriminoType: "S",
			x:             2,
			y:             2,
			want: Tetromino{
				Shape: [][]int{
					{0, 1, 1},
					{1, 1, 0},
				},
				X: 2,
				Y: 2,
			},
			wantErr: false,
		},
		{
			name:          "Z type",
			tetriminoType: "Z",
			x:             3,
			y:             3,
			want: Tetromino{
				Shape: [][]int{
					{1, 1, 0},
					{0, 1, 1},
				},
				X: 3,
				Y: 3,
			},
			wantErr: false,
		},
		{
			name:          "T type",
			tetriminoType: "T",
			x:             4,
			y:             4,
			want: Tetromino{
				Shape: [][]int{
					{1, 1, 1},
					{0, 1, 0},
				},
				X: 4,
				Y: 4,
			},
			wantErr: false,
		},
		{
			name:          "I type",
			tetriminoType: "I",
			x:             5,
			y:             5,
			want: Tetromino{
				Shape: [][]int{
					{1, 1, 1, 1},
				},
				X: 5,
				Y: 5,
			},
			wantErr: false,
		},
		{
			name:          "L type",
			tetriminoType: "L",
			x:             6,
			y:             6,
			want: Tetromino{
				Shape: [][]int{
					{1, 0},
					{1, 0},
					{1, 1},
				},
				X: 6,
				Y: 6,
			},
			wantErr: false,
		},
		{
			name:          "J type",
			tetriminoType: "J",
			x:             7,
			y:             7,
			want: Tetromino{
				Shape: [][]int{
					{0, 1},
					{0, 1},
					{1, 1},
				},
				X: 7,
				Y: 7,
			},
			wantErr: false,
		},
		{
			name:          "Unknown type",
			tetriminoType: "Unknown",
			x:             0,
			y:             0,
			want:          Tetromino{},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTetromino(tt.tetriminoType, tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTetromino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Add unit tests for clearing full rows
// Add unit tests for checking Tetromino placement
func TestPlaceTetromino(t *testing.T) {
	tests := []struct {
		name          string
		tetriminoType string
		x             int
		y             int
		startY        int
		grid          [][]int
		expected      [][]int
		wantErr       bool
	}{
		{
			name:          "place Q at top left",
			tetriminoType: "Q",
			x:             0,
			y:             0,
			startY:        0,
			grid:          makeEmptyGrid(4, 4),
			expected: [][]int{
				{1, 1, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			wantErr: false,
		},
		{
			name:          "place Q at top x=1",
			tetriminoType: "Q",
			x:             1,
			y:             0,
			startY:        0,
			grid:          makeEmptyGrid(4, 4),
			expected: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			wantErr: false,
		},
		{
			name:          "place Q at bottom right",
			tetriminoType: "Q",
			x:             2,
			y:             2,
			startY:        2,
			grid:          makeEmptyGrid(4, 4),
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 1, 1},
			},
			wantErr: false,
		},
		{
			name:          "Invalid tetromino type",
			tetriminoType: "Unknown",
			x:             0,
			y:             0,
			startY:        0,
			grid:          makeEmptyGrid(4, 4),
			expected:      makeEmptyGrid(4, 4),
			wantErr:       true,
		},
		// More tests for other scenarios and tetromino types
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetromino, err := NewTetromino(tt.tetriminoType, tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTetromino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return // Do not proceed if an error is expected and received
			}
			tetromino.place(tt.startY, tt.grid, false)
			if !reflect.DeepEqual(tt.grid, tt.expected) {
				t.Errorf("placeTetromino() for %s, got grid = %v, want %v", tt.name, tt.grid, tt.expected)
			}
		})
	}
}

func TestCalculateHeight(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Single Q block at bottom left",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: 2, // The 'Q' block fills the two bottommost rows
		},
		{
			name: "two Q's",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 0, 1, 1, 0, 0, 0, 0},
				{0, 1, 1, 0, 1, 1, 0, 0, 0, 0},
			},
			expected: 2, // The 'Q' blocks are still only 2 high
		},
		{
			name: "L on top of Q's",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
				{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
			},
			expected: 5, // The 'Q' blocks are still only 2 high
		},
		// Additional test cases can be added here.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateHeight(tt.grid, false)
			if result != tt.expected {
				t.Errorf("TestCalculateHeight failed for %s: expected height %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

// Add unit tests for clearing full rows
func TestClearFullRows(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected [][]int
	}{
		{
			name: "Clear single full row",
			grid: [][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 1}, // This row should be cleared
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Clear 2 full rows",
			grid: [][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 1}, // This row should be cleared
				{1, 1, 1, 1},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 0},
			},
		},
		// More tests for multiple full rows, edge cases, etc.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearFullRows(tt.grid, false)
			if !reflect.DeepEqual(tt.grid, tt.expected) {
				t.Errorf("clearFullRows() for %s, got grid = %v, want %v", tt.name, tt.grid, tt.expected)
			}
		})
	}
}

func TestDropTetromino(t *testing.T) {
	width, height := 10, 10

	for x := 0; x <= width-2; x++ { // Ensure the 2x2 'Q' block fits within the grid's width
		t.Run(fmt.Sprintf("Drop Q Tetromino from x=%d", x), func(t *testing.T) {
			grid := makeEmptyGrid(height, width)    // Create a grid of height 10 and width 10
			tetromino, _ := NewTetromino("Q", x, 0) // Create a 'Q' Tetromino starting at different x positions

			// Expected final grid after dropping the Tetromino
			expectedGrid := makeEmptyGrid(height, width)
			expectedGrid[height-2][x] = 1
			expectedGrid[height-2][x+1] = 1
			expectedGrid[height-1][x] = 1
			expectedGrid[height-1][x+1] = 1

			// Drop the Tetromino
			tetromino.Drop(grid, false)

			if !reflect.DeepEqual(grid, expectedGrid) {
				t.Errorf("DropTetromino failed: expected grid to have the Tetromino at x=%d, got different placement", x)
			}
		})
	}
}

// Utility function to create an empty grid
func makeEmptyGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}
