package solver

// Location indicates the location of a block or piece.
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Piece represents a puzzle piece used to cover the board.
type Piece struct {
	// Definition is used to define the piece. This does not get changed.
	Definition []Location
	// blocks are the current location of each of the blocks of the piece.
	blocks []Location
}

// Blocks returns the current placement of the piece's blocks.
func (piece *Piece) Blocks() []Location {
	if len(piece.blocks) == 0 {
		piece.blocks = append(piece.blocks, piece.Definition...)
	}

	return piece.blocks
}

// Translate moves a piece around the x/y plane
func (piece *Piece) Translate(x, y int) {
	if len(piece.blocks) == 0 {
		piece.blocks = append(piece.blocks, piece.Definition...)
	}

	for k, block := range piece.blocks {
		piece.blocks[k].X = block.X + x
		piece.blocks[k].Y = block.Y + y
	}
}

// Rotate90CW rotates a piece 90 degrees clockwise n times.
func (piece *Piece) Rotate90CW() {
	if len(piece.blocks) == 0 {
		piece.blocks = append(piece.blocks, piece.Definition...)
	}

	maxX := 0
	for k, block := range piece.blocks {
		if block.X > maxX {
			maxX = block.X
		}

		piece.blocks[k].X = block.Y
		piece.blocks[k].Y = -block.X
	}

	piece.Translate(0, maxX)
}

// Reset sets a piece's blocks back to the original definition.
func (piece *Piece) Reset() {
	piece.blocks = append([]Location{}, piece.Definition...)
}

// Equals determines the equality of piece a to piece b. If they have the same
// block structure then this returns true.
func (piece *Piece) Equals(b *Piece) bool {
	for _, ablock := range piece.Blocks() {
		found := false
		for _, bblock := range b.Blocks() {
			if ablock.X == bblock.X && ablock.Y == bblock.Y {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}

	return true
}

// GetUniqueRotations determines the number of unique rotation results for the piece.
func (piece *Piece) GetUniqueRotations() int {
	rotations := []*Piece{}
	for i := 0; i < 4; i++ {
		match := false
		for _, p := range rotations {
			if piece.Equals(p) {
				match = true
				break
			}
		}

		if !match {
			new := &Piece{
				Definition: piece.Definition,
				blocks:     append([]Location{}, piece.blocks...),
			}
			rotations = append(rotations, new)
		}

		piece.Rotate90CW()
	}

	return len(rotations)
}
