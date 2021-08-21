package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zillani/beginning-golang-game-programming/ebiten/util"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	InitBallRadius = 10.0
)

type Game struct {
	count int
}

// Ball is a final ball
type Ball struct {
	util.Position
	Radius    float32
	XVelocity float32
	YVelocity float32
	Color     color.Color
	Img       *ebiten.Image
}

func setBallPixels(c color.Color, ballImg *ebiten.Image) {
	// TODO: set pixels for round effect
	ballImg.Fill(c)
}

func (b *Game) Update(screen *ebiten.Image) {

}

func (b *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shapes (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}