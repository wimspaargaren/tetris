package game

import "math/rand"

// Point point of a shape
type Point struct {
	row int
	col int
}

// Shape shape of a piece consisting of 4 points
type Shape [4]Point

// PieceType type of a Tetris piece
type PieceType int

// Different piece types
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

func (p *Piece) down() {
	for i := 0; i < 4; i++ {
		p.Shape[i].row++
	}
}

func (p *Piece) left() {
	p.moveHorizontal(-1)
}

func (p *Piece) right() {
	p.moveHorizontal(1)
}

func (p *Piece) moveHorizontal(offset int) {
	for i := 0; i < 4; i++ {
		p.Shape[i].col += offset
	}
}

func (p *Piece) rotate() {
	// pivot piece at 1
	pivot := p.Shape[1]
	for i := 0; i < 4; i++ {
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
		panic("unknown piece type")
	}
}

func (p *Piece) offset() int {
	switch p.Type {
	case IType:
		return rand.Intn(7)
	case OType:
		return rand.Intn(9)
	case JType, LType, SType, TType, ZType:
		return rand.Intn(8)
	default:
		return rand.Intn(8)
	}
}
