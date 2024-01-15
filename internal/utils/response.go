package utils

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(ctx echo.Context, code int, message string, data interface{}) error {
	return ctx.JSON(code, &Response{
		Code:    code,
		Status:  "Success",
		Message: message,
		Data:    data,
	})
}
