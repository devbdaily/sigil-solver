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
