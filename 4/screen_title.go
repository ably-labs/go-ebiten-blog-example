package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

var (
	worldImage *ebiten.Image
)

func initialiseTitleScreen() {
	// initialise image from image files.
	var err error
	worldImage, _, err = ebitenutil.NewImageFromFile("./images/world.png")
	if err != nil {
		log.Println(err)
	}
}

func drawTitleScreen(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(worldImage, opts)
}

func updateTitleScreen() {
	// Detect if N key or S key is pressed.
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		state = northernScreen
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		state = southernScreen
	}
}
