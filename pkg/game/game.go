// Package game provides a simple Tetris implementation
package game

import (
	"math/rand"
	"time"
)

// don't lint since we want to ensure random blocks being generated
// nolint: gochecknoinits
func init() {
	rand.Seed(time.Now().Unix())
}

// StepResult result of performing a step in the game
type StepResult struct {
	GameOver bool
}

// Action is an action which can be taken in the game
type Action int

// List of possible actions
const (
	ActionTurn Action = iota
	ActionLeft
	ActionRight
	ActionSleep
)

// Game interface for the tetris game
type Game interface {
	Score() int
	Step(Action) *StepResult
	DoAction(Action)
	Board() *Board
}

// NewGame create a new Tetris game
func NewGame() Game {
	t := &Tetris{
		score: 0,
		board: NewBoard(),
	}
	gameOver := t.newPiece()
	if gameOver {
		panic("unable to create new game")
	}
	return t
}

// Tetris the game implementation
type Tetris struct {
	board        *Board
	score        int
	currentPiece *Piece
}

// Score retrieves the current score
func (g *Tetris) Score() int {
	return g.score
}

// Step perform a game step
func (g *Tetris) Step(a Action) *StepResult {
	// g.board.debug()
	if a != ActionSleep {
		g.DoAction(a)
	}
	gameOver := g.movePieceDown()
	return &StepResult{
		GameOver: gameOver,
	}
}

// DoAction perform given action
func (g *Tetris) DoAction(a Action) {
	switch a {
	case ActionLeft:
		g.board.moveLeft(g.currentPiece)
	case ActionRight:
		g.board.moveRight(g.currentPiece)
	case ActionTurn:
		g.board.rotate(g.currentPiece)
	case ActionSleep:
		return
	}
}

// Board retrieves the current board
func (g *Tetris) Board() *Board {
	return g.board
}

func (g *Tetris) newPiece() bool {
	newPiece := newPieceFromType(PieceType(rand.Intn(int(ZType))))
	// Use random offset for piece placement
	pieceOffset := newPiece.offset()
	newPiece.moveHorizontal(pieceOffset)
	// If can't place, it's game over
	if !g.board.canPlace(newPiece) {
		return true
	}

	g.currentPiece = newPiece
	g.board.setPiece(g.currentPiece)
	return false
}

func (g *Tetris) movePieceDown() bool {
	gameOver := false
	g.board.unSetPiece(g.currentPiece)
	collided := !g.board.canMoveDown(g.currentPiece)
	if !collided {
		g.currentPiece.down()
	}
	g.board.setPiece(g.currentPiece)
	if collided {
		g.score += 10
		completedLines := g.board.checkCompletedLines()
		g.score += g.completedLinesScore(completedLines)

		gameOver = g.newPiece()
	}

	return gameOver
}

func (g *Tetris) completedLinesScore(completedLines int) int {
	switch completedLines {
	case 1:
		return 100
	case 2:
		return 300
	case 3:
		return 500
	case 4:
		return 800
	default:
		return 0
	}
}
