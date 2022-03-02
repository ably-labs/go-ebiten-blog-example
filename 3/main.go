package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 1366
	screenHeight = 768
)

var (
	state gameState
)

func init() {
	state = titleScreen
}

type Game struct{}

func (g *Game) Update() error {
	// Handle updates for each game state.
	switch state {
	case titleScreen:
		updateTitleScreen()
	case northernScreen:
		updateNorthernScreen()
	case southernScreen:
		updateSouthernScreen()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Handle drawing for each game state.
	switch state {
	case titleScreen:
		drawTitleScreen(screen)
	case northernScreen:
		drawNorthernScreen(screen)
	case southernScreen:
		drawSouthernScreen(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Game")

	initialiseTitleScreen()
	initialiseNorthernScreen()
	initialiseSouthernScreen()

	var myGame *Game
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
