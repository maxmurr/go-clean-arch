package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxmurr/go-clean-arch/config"
	"github.com/maxmurr/go-clean-arch/internal/server"

	floorHttpHandlers "github.com/maxmurr/go-clean-arch/modules/floor/handlers"
	floorRepositories "github.com/maxmurr/go-clean-arch/modules/floor/repositories"
	floorUsecases "github.com/maxmurr/go-clean-arch/modules/floor/usecases"
	"gorm.io/gorm"
)

type EchoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) server.Server {
	return &EchoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *EchoServer) Start() {
	validate := validator.New()

	s.InitializeFloorHandler(validate)

	router := s.app.Group("/api/v1")
	router.GET("", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	s.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))

}

func (s *EchoServer) InitializeFloorHandler(validator *validator.Validate) {
	// Initialize all layers
	floorRepository := floorRepositories.NewFloorRepositoryImpl(s.db)

	floorUsecase := floorUsecases.NewFloorUsecaseImpl(floorRepository, validator)

	floorHttpHandler := floorHttpHandlers.NewFloorHttpHandler(floorUsecase)

	// Routers
	floorRouters := s.app.Group("/api/v1/floors")
	floorRouters.GET("", floorHttpHandler.GetAllFloor)
	floorRouters.GET("/:id", floorHttpHandler.GetFloorById)
	floorRouters.POST("", floorHttpHandler.CreateFloor)
	floorRouters.PUT("/:id", floorHttpHandler.UpdateFloor)
	floorRouters.DELETE("/:id", floorHttpHandler.DeleteFloor)
}
