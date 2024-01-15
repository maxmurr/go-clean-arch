package main

import (
	"github.com/maxmurr/go-clean-arch/config"
	database "github.com/maxmurr/go-clean-arch/internal/database/postgres"
	server "github.com/maxmurr/go-clean-arch/internal/server/echo"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.Getdb()).Start()
}
