package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"log"
)

var (
	southernImage *ebiten.Image
)

func initialiseSouthernScreen() {
	// initialise image from image files.
	var err error
	southernImage, _, err = ebitenutil.NewImageFromFile("./images/southern.png")
	if err != nil {
		log.Println(err)
	}
}

func drawSouthernScreen(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(southernImage, opts)

	text.Draw(screen, "The Southern Hemisphere", mainFont, screenWidth/2, screenHeight/4, white)
}

func updateSouthernScreen() {
	// If the Escape key is pressed, return to the title screen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}
}
