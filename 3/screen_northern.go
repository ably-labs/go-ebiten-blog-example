package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
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
}

func updateNorthernScreen() {
	// If the Escape key is pressed, return to the title screen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		state = titleScreen
	}
}
