package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ReadRequest(ctx echo.Context, result interface{}) {
	decoder := json.NewDecoder(ctx.Request().Body)
	err := decoder.Decode(result)
	ErrorPanic(err)
}

func WriteResponse(ctx echo.Context, result interface{}) {
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response().Writer)
	err := encoder.Encode(result)
	ErrorPanic(err)
}
