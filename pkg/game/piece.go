package game

// Point point of a shape
type Point struct {
	row int
	col int
}

// Shape shape of a piece consisting of 4 points
type Shape [4]Point

// PieceType is a constant for a shape of piece. There are 7 classic pieces like L, and O
type PieceType int

// Various values that the pieces can be
const (
	IType PieceType = iota
	JType
	LType
	OType
	SType
	TType
	ZType
)

// Piece tetris piece which can be moved
type Piece struct {
	Shape *Shape
	Type  PieceType
}

func (p *Piece) copy() *Piece {
	return &Piece{
		Shape: &Shape{
			p.Shape[0],
			p.Shape[1],
			p.Shape[2],
			p.Shape[3],
		},
		Type: p.Type,
	}
}

func (p *Piece) gravity() {
	for i := 0; i < 4; i++ {
		p.Shape[i].row++
	}
}

func (p *Piece) left() {
	for i := 0; i < 4; i++ {
		p.Shape[i].col--
	}
}

func (p *Piece) right() {
	for i := 0; i < 4; i++ {
		p.Shape[i].col++
	}
}

func (p *Piece) rotate() {
	pivot := p.Shape[1]
	for i := 0; i < 4; i++ {
		// Index 1 is the pivot point
		if i == 1 {
			continue
		}
		dRow := pivot.row - p.Shape[i].row
		dCol := pivot.col - p.Shape[i].col
		p.Shape[i].row = pivot.row + (dCol * -1)
		p.Shape[i].col = pivot.col + (dRow)
	}
}

func newPieceFromType(pieceType PieceType) *Piece {
	switch pieceType {
	case LType:
		return &Piece{
			Type: LType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 1, col: 1},
				Point{row: 1, col: 2},
				Point{row: 0, col: 0},
			},
		}
	case IType:
		return &Piece{
			Type: IType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 1, col: 1},
				Point{row: 1, col: 2},
				Point{row: 1, col: 3},
			},
		}
	case OType:
		return &Piece{
			Type: OType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 1, col: 1},
				Point{row: 0, col: 0},
				Point{row: 0, col: 1},
			},
		}
	case TType:
		return &Piece{
			Type: TType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 1, col: 1},
				Point{row: 1, col: 2},
				Point{row: 0, col: 1},
			},
		}
	case SType:
		return &Piece{
			Type: SType,
			Shape: &Shape{
				Point{row: 0, col: 0},
				Point{row: 0, col: 1},
				Point{row: 1, col: 1},
				Point{row: 1, col: 2},
			},
		}
	case ZType:
		return &Piece{
			Type: ZType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 1, col: 1},
				Point{row: 0, col: 1},
				Point{row: 0, col: 2},
			},
		}
	case JType:
		return &Piece{
			Type: JType,
			Shape: &Shape{
				Point{row: 1, col: 0},
				Point{row: 0, col: 1},
				Point{row: 0, col: 0},
				Point{row: 0, col: 2},
			},
		}
	default:
		panic("getShapeFromPiece(Piece): Invalid piece entered")
	}
}
