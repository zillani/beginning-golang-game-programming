package main

import (
	"github.com/x-hgg-x/goecsengine/loader"
	"github.com/x-hgg-x/goecsengine/utils"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/pelletier/go-toml"
	ecs "github.com/x-hgg-x/goecs/v2"
)

type gameComponentList struct {
	Gopher *Gopher
	Sticky *Sticky
}

type entity struct {
	Components gameComponentList
}

type entityGameMetadata struct {
	Entities []entity `toml:"entity"`
}

func loadGameComponents(entityMetadataPath string, world w.World) []interface{} {
	var entityGameMetadata entityGameMetadata
	tree, err := toml.LoadFile(entityMetadataPath)
	utils.LogError(err)
	utils.LogError(tree.Unmarshal(&entityGameMetadata))

	gameComponentList := make([]interface{}, len(entityGameMetadata.Entities))
	for iEntity, entity := range entityGameMetadata.Entities {
		gameComponentList[iEntity] = entity.Components
	}
	return gameComponentList
}

// LoadEntities creates entities with components from a TOML file
func LoadEntities(entityMetadataPath string, world w.World) []ecs.Entity {
	gameComponentList := loadGameComponents(entityMetadataPath, world)
	return loader.LoadEntities(entityMetadataPath, world, gameComponentList)
}
