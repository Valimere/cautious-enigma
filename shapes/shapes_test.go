package shapes

import (
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
	}{
		{
			name:          "Q type",
			tetriminoType: "Q",
			x:             1,
			y:             1,
			want: Tetromino{
				shape: [][]int{
					{1, 1},
					{1, 1},
				},
				x: 1,
				y: 1,
			},
		},
		{
			name:          "S type",
			tetriminoType: "S",
			x:             2,
			y:             2,
			want: Tetromino{
				shape: [][]int{
					{0, 1, 1},
					{1, 1, 0},
				},
				x: 2,
				y: 2,
			},
		},
		// include tests for other types and edge cases accordingly
		{
			name:          "Unkown type",
			tetriminoType: "Unknown",
			x:             0,
			y:             0,
			want:          Tetromino{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTetromino(tt.tetriminoType, tt.x, tt.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}
