package solver

import (
	"reflect"
	"testing"
)

func TestPieceTranslate(t *testing.T) {
	piece := Piece{
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
	}

	piece.Translate(3, 1)

	want := []Location{
		{
			X: 3,
			Y: 1,
		},
		{
			X: 3,
			Y: 2,
		},
		{
			X: 4,
			Y: 1,
		},
		{
			X: 5,
			Y: 1,
		},
	}

	if !reflect.DeepEqual(piece.Blocks(), want) {
		t.Errorf("Piece does not show correct block locations after translation.\nWant: %v\nGot: %v\n", want, piece.Blocks())
	}
}

func TestPieceRotate90CW(t *testing.T) {
	piece := Piece{
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
	}

	piece.Rotate90CW()

	want := []Location{
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
	}

	if !reflect.DeepEqual(piece.Blocks(), want) {
		t.Errorf("Piece does not show correct block locations after rotation.\nWant: %v\nGot: %v\n", want, piece.Blocks())
	}
}

func TestPieceReset(t *testing.T) {
	piece := Piece{
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
	}

	piece.Rotate90CW()

	piece.Reset()

	want := []Location{
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
	}

	if !reflect.DeepEqual(piece.Blocks(), want) {
		t.Errorf("Piece does not show correct block locations after reset.\nWant: %v\nGot: %v\n", want, piece.Blocks())
	}
}
