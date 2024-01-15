package utils

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetParamIds(ctx echo.Context) ([]uint, error) {
	idStr := ctx.Param("id")
	idParts := strings.Split(idStr, ",")

	ids := make([]uint, len(idParts))
	for i, idPart := range idParts {
		id, err := strconv.Atoi(idPart)
		if err != nil {
			return nil, err
		}
		ids[i] = uint(id)
	}

	return ids, nil
}
