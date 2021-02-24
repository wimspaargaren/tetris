package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/wimspaargaren/tetris/pkg/displayer"
	"github.com/wimspaargaren/tetris/pkg/game"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	displayer, err := displayer.NewDisplayer()
	if err != nil {
		panic(err)
	}

	tetris := game.NewGame()
	lastStep := time.Now()
	for !displayer.WindowClosed() {
		action := displayer.Render(tetris)
		// Step every 200 milliseconds
		if time.Since(lastStep) > time.Millisecond*200 {
			lastStep = time.Now()
			result := tetris.Step(action)
			if result.GameOver {
				fmt.Println("Game over")
				break
			}
		} else {
			tetris.DoAction(action)
		}
	}
}
