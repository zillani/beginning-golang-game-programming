package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)


// Position is a set of coordinates in 2-D plan
type Position struct {
	X, Y float32
}

// GetCenter returns the center position on screen
func GetCenter(screen *ebiten.Image) Position {
	w, h := screen.Size()
	return Position{
		X: float32(w / 2),
		Y: float32(h / 2),
	}
}

// GameState is an enum that represents all possible game states
type GameState byte

const (
	StartState GameState = iota
	ControlsState
	PlayState
	InterState
	PauseState
	GameOverState
)

var (
	BgColor  = color.Black
	ObjColor = color.RGBA{120, 226, 160, 255}
)

func GetVertexIndex(x, y, w, h float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff
	x0 := x
	y0 := y
	x1 := x + w
	y1 := y + h

	return []ebiten.Vertex{
		{
			DstX:   x0,
			DstY:   y0,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1,
			DstY:   y0,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0,
			DstY:   y1,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1,
			DstY:   y1,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2, 1, 2, 3}
}

func GetRGBA(R, G, B int) color.RGBA {
	return color.RGBA{uint8(R), uint8(G), uint8(B), 255}
}

func DrawRectangle(screen *ebiten.Image){
	//v, i := rect(50, 50,100,100, GetRGBA(255, 45, 137))
	//screen.DrawTriangles(v, i, src, nil)
}