package game

import (
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
			Name: "Multi rows",
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
			Name: "Multi rows",
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
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			board := test.CreateBoard()
			completedLines := board.checkCompletedLines()
			s.Equal(test.CompletedLines, completedLines)
			s.Equal(test.ExpectedBoard(), *board)
		})
	}
}

func filledRow() [10]int {
	return [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
}

func TestBoardTestSuite(t *testing.T) {
	suite.Run(t, new(BoardTestSuite))
}
