package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/zillani/beginning-golang-game-programming/ebiten/util"
	"image"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	emptyImage = ebiten.NewImage(3, 3)
)

func init() {
	emptyImage.Fill(color.White)
}

type Game struct {
	count int
}


func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rect := image.Rect(1, 1, 3, 3)
	img := emptyImage.SubImage(rect)
	imgCast := img.(*ebiten.Image)
	rgba := util.GetRGBA(255, 45, 137)
	v, i := util.GetVertexIndex(50, 50,100,100, rgba)
	screen.DrawTriangles(v, i, imgCast, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
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