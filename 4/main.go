package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 1366
	screenHeight = 768
	fontDpi      = 72
)

var (
	state    gameState
	mainFont font.Face
	white    = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
)

func init() {
	state = titleScreen

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mainFont = truetype.NewFace(tt, &truetype.Options{
		Size:    32,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})
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
