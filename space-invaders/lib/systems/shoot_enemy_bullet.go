package systems

import (
	"math/rand"

	gc "github.com/zillani/beginning-golang-game-programming/space-invaders/lib/components"
	"github.com/zillani/beginning-golang-game-programming/space-invaders/lib/resources"

	ecs "github.com/x-hgg-x/goecs/v2"
	ec "github.com/x-hgg-x/goecsengine/components"
	"github.com/x-hgg-x/goecsengine/loader"
	"github.com/x-hgg-x/goecsengine/math"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
)

var shootEnemyBulletFrame = ebiten.DefaultTPS

// ShootEnemyBulletSystem shoots enemy bullet
func ShootEnemyBulletSystem(world w.World) {
	shootEnemyBulletFrame--

	gameComponents := world.Components.Game.(*gc.Components)
	gameResources := world.Resources.Game.(*resources.Game)

	alienSet := world.Manager.Join(gameComponents.Alien, gameComponents.AlienMaster.Not())
	if alienSet.Empty() {
		return
	}

	if shootEnemyBulletFrame <= 0 {
		shootEnemyBulletFrame = int(ebiten.DefaultTPS / float64(gameResources.Difficulty) * rand.Float64())

		// Select random alien
		alienEntities := []ecs.Entity{}
		alienSet.Visit(ecs.Visit(func(entity ecs.Entity) {
			alienEntities = append(alienEntities, entity)
		}))
		alienEntity := alienEntities[rand.Intn(len(alienEntities))]
		alienHeight := gameComponents.Alien.Get(alienEntity).(*gc.Alien).Height
		alienTranslation := world.Components.Engine.Transform.Get(alienEntity).(*ec.Transform).Translation

		enemyBulletEntity := loader.AddEntities(world, world.Resources.Prefabs.(*resources.Prefabs).Game.EnemyBullet)
		for iEntity := range enemyBulletEntity {
			enemyBulletHeight := gameComponents.Bullet.Get(enemyBulletEntity[iEntity]).(*gc.Bullet).Height
			enemyBulletTransform := world.Components.Engine.Transform.Get(enemyBulletEntity[iEntity]).(*ec.Transform)
			enemyBulletTransform.Translation = math.Vector2{
				X: alienTranslation.X,
				Y: alienTranslation.Y - alienHeight/2 - enemyBulletHeight/2,
			}
		}
	}
}
