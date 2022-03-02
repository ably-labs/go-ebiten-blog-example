package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	northernImage *ebiten.Image
)

func initialiseNorthernScreen() {
	// initialise image from image files.
	var err error
	northernImage, _, err = ebitenutil.NewImageFromFile("./images/northern.png")
	if err != nil {
		log.Println(err)
	}
}

func drawNorthernScreen(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(northernImage, opts)

	text.Draw(screen, "The Northern Hemisphere", mainFont, screenWidth/2, screenHeight/4, white)
}

func updateNorthernScreen() {
	// If the Escape key is pressed, return to the title screen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}
}
