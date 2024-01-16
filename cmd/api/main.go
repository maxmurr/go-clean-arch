package main

import (
	"github.com/maxmurr/go-clean-arch/config"
	"github.com/maxmurr/go-clean-arch/internal/database"
	"github.com/maxmurr/go-clean-arch/internal/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.Getdb()).Start()
}
