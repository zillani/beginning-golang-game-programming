package states

import (
	gloader "github.com/x-hgg-x/sokoban-go/lib/loader"
	"github.com/x-hgg-x/sokoban-go/lib/resources"
	g "github.com/x-hgg-x/sokoban-go/lib/systems"

	"github.com/x-hgg-x/goecsengine/states"
	"github.com/x-hgg-x/goecsengine/utils"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// GameplayState is the main game state
type GameplayState struct{}

// OnStart method
func (st *GameplayState) OnStart(world w.World) {
	// Define package name
	packageName := "xsokoban"

	// Load package
	utils.LogError(gloader.LoadPackage(packageName, world))

	// Load game
	game := resources.NewGame(world, packageName)
	world.Resources.Game = game

	// Load last played level
	levelNum := 0
	tree := resources.LoadSaveFile(world)
	if savedCurrentLevel, ok := tree.Get("CurrentLevel").(int64); ok {
		currentLevel := int(savedCurrentLevel) - 1
		if currentLevel != game.Level.CurrentNum && 0 <= currentLevel && currentLevel < game.LevelCount {
			levelNum = currentLevel
		}
	}

	resources.InitLevel(world, levelNum)
}

// OnPause method
func (st *GameplayState) OnPause(world w.World) {}

// OnResume method
func (st *GameplayState) OnResume(world w.World) {}

// OnStop method
func (st *GameplayState) OnStop(world w.World) {
	world.Manager.DeleteAllEntities()
	resources.SaveLevel(world)
	world.Resources.Game = nil
}

// Update method
func (st *GameplayState) Update(world w.World) states.Transition {
	g.SwitchLevelSystem(world)
	g.UndoSystem(world)
	g.MoveSystem(world)
	g.SaveSystem(world)
	g.InfoSystem(world)
	g.GridTransformSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	gameResources := world.Resources.Game.(*resources.Game)
	switch gameResources.StateEvent {
	case resources.StateEventLevelComplete:
		gameResources.StateEvent = resources.StateEventNone
		return states.Transition{Type: states.TransPush, NewStates: []states.State{&LevelCompleteState{}}}
	}

	return states.Transition{}
}
