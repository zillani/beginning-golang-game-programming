package main

import (
	"math/rand"
	"time"

	"github.com/x-hgg-x/goecsengine/states"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// GameplayState is the main game state
type GameplayState struct{}

// OnPause method
func (st *GameplayState) OnPause(world w.World) {}

// OnResume method
func (st *GameplayState) OnResume(world w.World) {}

// OnStart method
func (st *GameplayState) OnStart(world w.World) {
	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Load game and text entities
	LoadEntities("metadata/start.toml", world)
	LoadEntities("metadata/text.toml", world)

	world.Resources.Game = NewGame()
}

// OnStop method
func (st *GameplayState) OnStop(world w.World) {
	world.Resources.Game = nil
	world.Manager.DeleteAllEntities()
}

// Update method
func (st *GameplayState) Update(world w.World) states.Transition {
	DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}
	return states.Transition{}
}
