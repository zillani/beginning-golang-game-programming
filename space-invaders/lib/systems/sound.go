package systems

import (
	"github.com/zillani/beginning-golang-game-programming/space-invaders/lib/loader"
	"github.com/zillani/beginning-golang-game-programming/space-invaders/lib/resources"

	w "github.com/x-hgg-x/goecsengine/world"
)

// SoundSystem manages sound
func SoundSystem(world w.World) {
	if world.Resources.InputHandler.Actions[resources.EnableDisableSoundAction] {
		audioPlayers := *world.Resources.AudioPlayers
		if audioPlayers["music"].Volume() == 0 {
			loader.EnableSound(world)
		} else {
			loader.DisableSound(world)
		}
	}
}
