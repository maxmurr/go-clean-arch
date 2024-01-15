package database

import (
	"fmt"

	"github.com/maxmurr/go-clean-arch/config"
	"github.com/maxmurr/go-clean-arch/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) database.Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &PostgresDatabase{db: db}
}

func (p *PostgresDatabase) Getdb() *gorm.DB {
	return p.db
}
