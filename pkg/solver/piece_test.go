package solver

import (
	"reflect"
	"testing"
)

func TestPieceEquals(t *testing.T) {
	tests := []struct {
		name string
		a, b Piece
		want bool
	}{
		{
			name: "Same piece, different block order.",
			a: Piece{
				blocks: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
					{
						X: 3,
						Y: 0,
					},
				},
			},
			b: Piece{
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
						X: 1,
						Y: 0,
					},
					{
						X: 0,
						Y: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "same piece, but rotated (and therefore not equal)",
			a: Piece{
				blocks: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
					{
						X: 3,
						Y: 0,
					},
				},
			},
			b: Piece{
				blocks: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 0,
						Y: 1,
					},
					{
						X: 0,
						Y: 2,
					},
					{
						X: 0,
						Y: 3,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Equals(&tt.b)
			if tt.want != got {
				t.Errorf("a.Equals(b) is incorrect.\nWant: %v\nGot: %v\n", tt.want, got)
			}
		})
	}
}

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

func TestGetUniqueRotations(t *testing.T) {
	tests := []struct {
		name  string
		piece Piece
		want  int
	}{
		{
			name: "T tetramino has 4",
			piece: Piece{
				Definition: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 1,
						Y: 1,
					},
					{
						X: 2,
						Y: 0,
					},
				},
			},
			want: 4,
		},
		{
			name: "I tetramino has 2",
			piece: Piece{
				Definition: []Location{
					{
						X: 0,
						Y: 0,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 2,
						Y: 0,
					},
					{
						X: 3,
						Y: 0,
					},
				},
			},
			want: 2,
		},
		{
			name: "O tetramino has 1",
			piece: Piece{
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
						X: 1,
						Y: 1,
					},
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.piece.GetUniqueRotations()

			if tt.want != got {
				t.Errorf("Incorrect number of unique rotations recorded.\nWant: %d\nGot: %d\n", tt.want, got)
			}
		})
	}
}
