package handlers

import "github.com/labstack/echo/v4"

type FloorHandler interface {
	CreateFloor(c echo.Context) error
	GetAllFloor(c echo.Context) error
	GetFloorById(c echo.Context) error
	UpdateFloor(c echo.Context) error
	DeleteFloor(c echo.Context) error
}
