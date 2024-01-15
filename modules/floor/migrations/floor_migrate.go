package main

import (
	"github.com/maxmurr/go-clean-arch/config"
	database "github.com/maxmurr/go-clean-arch/internal/database"
	postgres "github.com/maxmurr/go-clean-arch/internal/database/postgres"
	"github.com/maxmurr/go-clean-arch/modules/floor/entities"
)

func main() {
	cfg := config.GetConfig()

	db := postgres.NewPostgresDatabase(&cfg)

	FloorMigrate(db)
}

func FloorMigrate(db database.Database) {
	db.Getdb().Migrator().CreateTable(&entities.Floor{})
	db.Getdb().CreateInBatches([]entities.Floor{
		{Name: "Floor 1"},
		{Name: "Floor 2"},
		{Name: "Floor 3"},
		{Name: "Floor 4"},
	}, 4)
}
