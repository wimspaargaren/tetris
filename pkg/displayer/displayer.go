// Package displayer provides a simple implementation for viewing a Tetris game
package displayer

import (
	"fmt"
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/wimspaargaren/tetris/pkg/game"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	windowWidth   = 765.0
	windowHeight  = 450.0
	scoreTextLocX = windowWidth - 265
	scoreTextLocY = windowHeight - 50
	blockSize     = 20
	gameDisplayY  = windowHeight + blockSize/2
	gameDisplayX  = 100
)

// Displayer displayer of the game
type Displayer interface {
	Render(game.Game) game.Action
	WindowClosed() bool
}

// GameDisplayer implementation of the game displayer
type GameDisplayer struct {
	Window       *pixelgl.Window
	SpriteActive *pixel.Sprite
}

// BlockImg block with given color
type BlockImg struct {
	Color color.Color
}

// NewBlockImg creates a new block for given color
func NewBlockImg(color color.Color) image.Image {
	return &BlockImg{
		Color: color,
	}
}

// ColorModel implementation of image.Image
func (t *BlockImg) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds implementation of image.Image
func (t *BlockImg) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{}, Max: image.Point{blockSize, blockSize}}
}

// At implementation of image.Image
func (t *BlockImg) At(x, y int) color.Color {
	return color.White
}

// Render displays the current game for the given window
func (g *GameDisplayer) Render(tetris game.Game) game.Action {
	g.Window.Clear(colornames.Black)
	a := game.ActionSleep
	if g.Window.JustPressed(pixelgl.KeyLeft) {
		a = game.ActionLeft
	}
	if g.Window.JustPressed(pixelgl.KeyRight) {
		a = game.ActionRight
	}
	if g.Window.JustPressed(pixelgl.KeyUp) {
		a = game.ActionTurn
	}
	g.renderBoard(tetris)
	g.renderScore(tetris)

	g.Window.Update()
	return a
}

// WindowClosed reports if window is closed
func (g *GameDisplayer) WindowClosed() bool {
	if g.Window == nil {
		return false
	}
	return g.Window.Closed()
}

func (g *GameDisplayer) renderScore(game game.Game) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	scoreTxt := text.New(pixel.V(scoreTextLocX, scoreTextLocY), basicAtlas)
	fmt.Fprintf(scoreTxt, "Score: %d", game.Score())
	scoreTxt.Draw(g.Window, pixel.IM.Scaled(scoreTxt.Orig, 2))
}

func (g *GameDisplayer) renderBoard(game game.Game) {
	board := game.Board()
	gameDisplayRowStart := 2
	for i := gameDisplayRowStart; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// Only draw active shapes
			if board[i][j] == 1 {
				blockX := gameDisplayX + j*blockSize
				blockY := gameDisplayY - i*blockSize
				g.SpriteActive.Draw(g.Window, pixel.IM.Scaled(pixel.ZV, 0.9).Moved(pixel.V(float64(blockX), float64(blockY))))
			}
		}
	}
}

// NewDisplayer creates a new game displayer
func NewDisplayer() (Displayer, error) {
	picData := pixel.PictureDataFromImage(NewBlockImg(color.White))
	spriteActive := pixel.NewSprite(picData, picData.Bounds())
	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}

	return &GameDisplayer{
		Window:       win,
		SpriteActive: spriteActive,
	}, nil
}
