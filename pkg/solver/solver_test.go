package solver

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	// Example using first puzzle
	board := Board{
		Width:  4,
		Height: 3,
	}

	pieces := []Piece{
		{
			Definition: []Location{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
			},
		},
		{
			Definition: []Location{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
			},
		},
		{
			Definition: []Location{
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
			},
		},
	}

	solver := NewSolver(&board, pieces)

	want := Board{
		Width:  4,
		Height: 3,
		Placed: []Piece{
			{
				Definition: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 0,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
				},
				blocks: []Location{
					{
						X: 0,
						Y: 2,
					},
					{
						X: 1,
						Y: 2,
					},
					{
						X: 0,
						Y: 1,
					},
					{
						X: 0,
						Y: 0,
					},
				},
			},
			{
				Definition: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 0,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
				},
				blocks: []Location{
					{
						X: 3,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
					{
						X: 3,
						Y: 1,
					},
					{
						X: 3,
						Y: 2,
					},
				},
			},
			{
				Definition: []Location{
					{
						X: 0,
						Y: 1,
					},
					{
						X: 1,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
				},
				blocks: []Location{
					{
						X: 2,
						Y: 2,
					},
					{
						X: 2,
						Y: 1,
					},
					{
						X: 1,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
				},
			},
		},
	}

	got, err := solver.Solve()

	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Did not receive correct result.\nWant: %v\nGot:  %v\n", want, got)
	}
}
