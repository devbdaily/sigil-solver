package solver

import (
	"errors"
)

// Board represents the puzzle board's size.
type Board struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Placed []Piece
}

// Solver is the mechanism used to solve the puzzle.
type Solver interface {
	// Solve solves the puzzle.
	Solve() (result *Board, err error)
}

type solver struct {
	board  *Board
	pieces []Piece
}

// NewSolver creates a new Solver
func NewSolver(board *Board, pieces []Piece) Solver {
	return &solver{
		board:  board,
		pieces: pieces,
	}
}

func (s *solver) Solve() (result *Board, err error) {
	err = s.checkBlockCount()
	if err != nil {
		return
	}

	_, err = s.board.placePieces(s.pieces)
	if err != nil {
		return
	}

	result = s.board

	return
}

func (s *solver) checkBlockCount() error {
	blockCount := 0
	for _, piece := range s.pieces {
		blockCount += len(piece.Definition)
	}

	if blockCount != (s.board.Width * s.board.Height) {
		return errors.New("Board size does not match number of blocks within pieces")
	}

	return nil
}

func (b *Board) placePieces(pieces []Piece) (found bool, err error) {
	if len(pieces) == 0 {
		return
	}
	piece := pieces[0]
	remaining := pieces[1:]

	for i := 0; i < piece.GetUniqueRotations(); i++ {
		for j := 0; j < b.Height; j++ {
			for l := 0; l < b.Width; l++ {
				piece.Reset()
				for r := 0; r < i; r++ {
					piece.Rotate90CW()
				}
				piece.Translate(j, l)

				if b.checkPiece(piece) {
					b.Placed = append(b.Placed, piece)
					if len(remaining) == 0 {
						found = true
						return
					}
					found, err = b.placePieces(remaining)

					if found {
						return
					}

					// if this particular placement didn't work out, remove it.
					if err != nil {
						b.Placed = b.Placed[:len(b.Placed)-1]
						err = nil
					}
				}
			}
		}
	}

	err = errors.New("Puzzle unsolvable")

	return
}

func (b Board) checkPiece(piece Piece) bool {
	for _, block := range piece.Blocks() {
		if !b.inBoardRange(block) || b.blockExists(block) {
			return false
		}
	}

	return true
}

func (b Board) inBoardRange(loc Location) bool {
	return loc.X < b.Width && loc.Y < b.Height
}

func (b Board) blockExists(loc Location) bool {
	for _, piece := range b.Placed {
		for _, block := range piece.Blocks() {
			if loc.X == block.X && loc.Y == block.Y {
				return true
			}
		}
	}

	return false
}
