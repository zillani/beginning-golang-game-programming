package main

import (
	r "github.com/x-hgg-x/goecsengine/resources"
	_ "image/png"

	"github.com/x-hgg-x/goecsengine/loader"
	s "github.com/x-hgg-x/goecsengine/states"
	"github.com/x-hgg-x/goecsengine/utils"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 600
	windowHeight = 600
)

type mainGame struct {
	world        w.World
	stateMachine s.StateMachine
}

func (game *mainGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (game *mainGame) Update() error {
	game.stateMachine.Update(game.world)
	return nil
}

func (game *mainGame) Draw(screen *ebiten.Image) {
	game.stateMachine.Draw(game.world, screen)
}

func main() {
	world := w.InitWorld(nil)

	// Init screen dimensions
	world.Resources.ScreenDimensions = &r.ScreenDimensions{Width: windowWidth, Height: windowHeight}

	// Load sprite sheets
	spriteSheets := loader.LoadSpriteSheets("spritesheets.toml")
	world.Resources.SpriteSheets = &spriteSheets

	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Demo")

	utils.LogError(ebiten.RunGame(&mainGame{world, s.Init(&GameplayState{}, world)}))
}
