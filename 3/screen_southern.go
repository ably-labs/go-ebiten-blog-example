package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

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
}

func updateSouthernScreen() {
	// If the Escape key is pressed, return to the title screen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}
}
