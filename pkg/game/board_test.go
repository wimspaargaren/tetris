package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoardTestSuite struct {
	suite.Suite
}

func (s *BoardTestSuite) TestCheckCompletedLines() {
	tests := []struct {
		Name           string
		CreateBoard    func() *Board
		ExpectedBoard  func() Board
		CompletedLines int
	}{
		{
			Name: "Empty board",
			CreateBoard: func() *Board {
				b := NewBoard()
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				return *b
			},
			CompletedLines: 0,
		},
		{
			Name: "Bottom row",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[21] = filledRow()
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				return *b
			},
			CompletedLines: 1,
		},
		{
			Name: "Second row",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[20] = filledRow()
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return *b
			},
			CompletedLines: 1,
		},
		{
			Name: "Triple rows nothing on top",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[18] = filledRow()
				b[19] = filledRow()
				b[20] = filledRow()
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return *b
			},
			CompletedLines: 3,
		},
		{
			Name: "Tetris",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[16] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				b[17] = filledRow()
				b[18] = filledRow()
				b[19] = filledRow()
				b[20] = filledRow()
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				b[20] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				b[21] = [10]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}
				return *b
			},
			CompletedLines: 4,
		},
		{
			Name: "Triple rows multi on top",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[18] = [10]int{1, 0, 1, 1, 0, 0, 0, 0, 1, 0}
				b[19] = [10]int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
				b[20] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				b[21] = filledRow()
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				b[19] = [10]int{1, 0, 1, 1, 0, 0, 0, 0, 1, 0}
				b[20] = [10]int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
				b[21] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				return *b
			},
			CompletedLines: 1,
		},
		{
			Name: "Triple rows multi on top",
			CreateBoard: func() *Board {
				b := NewBoard()
				b[15] = [10]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
				b[16] = [10]int{1, 0, 1, 1, 0, 0, 0, 0, 1, 0}
				b[17] = [10]int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
				b[18] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				b[19] = filledRow()
				b[20] = filledRow()
				b[21] = filledRow()
				return b
			},
			ExpectedBoard: func() Board {
				b := NewBoard()
				b[18] = [10]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
				b[19] = [10]int{1, 0, 1, 1, 0, 0, 0, 0, 1, 0}
				b[20] = [10]int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
				b[21] = [10]int{1, 0, 1, 1, 0, 1, 1, 0, 1, 0}
				return *b
			},
			CompletedLines: 3,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			board := test.CreateBoard()
			completedLines := board.checkCompletedLines()

			s.Require().Equal(test.CompletedLines, completedLines)
			s.Equal(test.ExpectedBoard(), *board)
		})
	}
}

func debugBoard(board *Board) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}

func filledRow() [10]int {
	return [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
}

func TestBoardTestSuite(t *testing.T) {
	suite.Run(t, new(BoardTestSuite))
}
