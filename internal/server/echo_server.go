package server

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxmurr/go-clean-arch/config"

	floorHttpHandlers "github.com/maxmurr/go-clean-arch/modules/floor/handlers"
	floorRepositories "github.com/maxmurr/go-clean-arch/modules/floor/repositories"
	floorUsecases "github.com/maxmurr/go-clean-arch/modules/floor/usecases"
	"gorm.io/gorm"
)

const TIMEOUT = 30 * time.Second

type EchoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &EchoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *EchoServer) Start() error {
	validate := validator.New()
	s.InitializeFloorHandler(validate)
	s.app.Use(middleware.Logger())

	router := s.app.Group("/api/v1")
	router.GET("", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	defer stop()
	errShutdown := make(chan error, 1)
	go s.Shutdown(ctx, errShutdown)
	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	err := s.app.Start(serverUrl)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	err = <-errShutdown
	if err != nil {
		return err
	}
	return nil
}

func (s *EchoServer) Shutdown(ctx context.Context, errShutdown chan error) {
	<-ctx.Done()

	ctxTimeout, stop := context.WithTimeout(context.Background(), TIMEOUT)
	defer stop()

	err := s.app.Shutdown(ctxTimeout)
	switch err {
	case nil:
		errShutdown <- nil
	case context.DeadlineExceeded:
		errShutdown <- nil
	default:
		errShutdown <- err
	}
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
