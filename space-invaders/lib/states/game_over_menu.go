package states

import (
	"fmt"

	"github.com/zillani/beginning-golang-game-programming/space-invaders/lib/resources"
	g "github.com/zillani/beginning-golang-game-programming/space-invaders/lib/systems"

	ecs "github.com/x-hgg-x/goecs/v2"
	"github.com/x-hgg-x/goecsengine/loader"
	"github.com/x-hgg-x/goecsengine/states"
	w "github.com/x-hgg-x/goecsengine/world"
)

// GameOverState is the game over menu state
type GameOverState struct {
	difficulty   resources.Difficulty
	gameOverMenu []ecs.Entity
	selection    int
}

//
// Menu interface
//

func (st *GameOverState) getSelection() int {
	return st.selection
}

func (st *GameOverState) setSelection(selection int) {
	st.selection = selection
}

func (st *GameOverState) confirmSelection() states.Transition {
	switch st.selection {
	case 0:
		// Restart
		return states.Transition{Type: states.TransReplace, NewStates: []states.State{&GameplayState{game: resources.NewGame(st.difficulty)}}}
	case 1:
		// Main Menu
		return states.Transition{Type: states.TransReplace, NewStates: []states.State{&MainMenuState{}}}
	case 2:
		// Exit
		return states.Transition{Type: states.TransQuit}
	}
	panic(fmt.Errorf("unknown selection: %d", st.selection))
}

func (st *GameOverState) getMenuIDs() []string {
	return []string{"restart", "main_menu", "exit"}
}

func (st *GameOverState) getCursorMenuIDs() []string {
	return []string{"cursor_restart", "cursor_main_menu", "cursor_exit"}
}

//
// State interface
//

// OnPause method
func (st *GameOverState) OnPause(world w.World) {}

// OnResume method
func (st *GameOverState) OnResume(world w.World) {}

// OnStart method
func (st *GameOverState) OnStart(world w.World) {
	prefabs := world.Resources.Prefabs.(*resources.Prefabs)
	st.gameOverMenu = append(st.gameOverMenu, loader.AddEntities(world, prefabs.Menu.GameOverMenu)...)
}

// OnStop method
func (st *GameOverState) OnStop(world w.World) {
	world.Manager.DeleteEntities(st.gameOverMenu...)
}

// Update method
func (st *GameOverState) Update(world w.World) states.Transition {
	g.SoundSystem(world)

	return updateMenu(st, world)
}
