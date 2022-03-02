package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 1366
	screenHeight = 768
)

var (
	worldImage *ebiten.Image
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Draw image.
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(worldImage, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func initialiseGame() {
	// initialise images from image files.
	var err error
	worldImage, _, err = ebitenutil.NewImageFromFile("./images/world.png")
	if err != nil {
		log.Println(err)
	}
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Game")

	initialiseGame()

	var myGame *Game
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
