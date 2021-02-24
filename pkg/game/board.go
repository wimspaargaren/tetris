package game

// Board holding the tetris blocks
type Board [22][10]int

// NewBoard creates an empty board
func NewBoard() *Board {
	return &Board{}
}

func (b *Board) unSetPiece(p *Piece) {
	for i := 0; i < 4; i++ {
		b.setGridVal(p.Shape[i].row, p.Shape[i].col, 0)
	}
}

func (b *Board) setPiece(p *Piece) {
	for i := 0; i < 4; i++ {
		b.setGridVal(p.Shape[i].row, p.Shape[i].col, 1)
	}
}

func (b *Board) setGridVal(x, y, val int) {
	b[x][y] = val
}

func (b *Board) checkCompletedLines() int {
	completedLines := 0
	for i := len(b) - 1; i > 0; i-- {
		if b[i] != [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1} {
			continue
		}
		completedLines++
		for i-1 >= 0 {
			b[i] = b[i-1]
			i--
		}
		i = len(b) - 1
	}
	return completedLines
}

func (b *Board) canPlace(piece *Piece) bool {
	for i := 0; i < 4; i++ {
		row := piece.Shape[i].row
		col := piece.Shape[i].col
		if col < 0 || col > len(b[0])-1 {
			return false
		}
		if row < 0 || row > len(b)-1 {
			return false
		}
		if b[row][col] != 0 {
			return false
		}
	}
	return true
}

func (b *Board) canMoveDown(piece *Piece) bool {
	for i := 0; i < 4; i++ {
		row := piece.Shape[i].row
		if row >= len(b)-1 {
			return false
		}

		if b[row+1][piece.Shape[i].col] != 0 {
			return false
		}
	}
	return true
}

// func (b *Board) canMoveRight(piece *Piece) bool {
// 	for i := 0; i < 4; i++ {
// 		col := piece.Shape[i].col
// 		if col >= len(b[0])-1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (b *Board) canMoveLeft(piece *Piece) bool {
// 	for i := 0; i < 4; i++ {
// 		col := piece.Shape[i].col
// 		if col <= 0 {
// 			return false
// 		}
// 		if b[piece.Shape[i].row][col] != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func (b *Board) moveRight(piece *Piece) {
	b.unSetPiece(piece)
	copiedPiece := piece.copy()
	copiedPiece.right()
	if b.canPlace(copiedPiece) {
		*piece = *copiedPiece
	}
	b.setPiece(piece)
}

func (b *Board) moveLeft(piece *Piece) {
	b.unSetPiece(piece)
	copiedPiece := piece.copy()
	copiedPiece.left()
	if b.canPlace(copiedPiece) {
		*piece = *copiedPiece
	}
	b.setPiece(piece)
}

func (b *Board) rotate(piece *Piece) {
	b.unSetPiece(piece)
	copiedPiece := piece.copy()
	copiedPiece.rotate()
	if b.canPlace(copiedPiece) {
		*piece = *copiedPiece
	}
	b.setPiece(piece)
}
