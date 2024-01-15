package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxmurr/go-clean-arch/internal/utils"
	"github.com/maxmurr/go-clean-arch/modules/floor/models"
	"github.com/maxmurr/go-clean-arch/modules/floor/usecases"
)

type FloorHttpHandler struct {
	FloorUsecase usecases.FloorUsecase
}

func NewFloorHttpHandler(floorUsecase usecases.FloorUsecase) FloorHandler {
	return &FloorHttpHandler{
		FloorUsecase: floorUsecase,
	}
}

func (h *FloorHttpHandler) CreateFloor(ctx echo.Context) error {
	createFloorRequest := models.CreateFloorRequest{}
	utils.ReadRequest(ctx, &createFloorRequest)

	h.FloorUsecase.Create(createFloorRequest)

	return utils.WebResponse(
		ctx,
		http.StatusCreated,
		"Floor created successfully",
		nil,
	)
}

func (h *FloorHttpHandler) UpdateFloor(ctx echo.Context) error {
	request := models.UpdateFloorRequest{}
	floorIds, err := utils.GetParamIds(ctx)
	utils.ErrorPanic(err)

	utils.ReadRequest(ctx, &request)
	for _, id := range floorIds {
		request.Id = id
		h.FloorUsecase.Update(request)
	}

	return utils.WebResponse(
		ctx,
		http.StatusOK,
		"Floor updated successfully",
		nil,
	)
}

func (h *FloorHttpHandler) DeleteFloor(ctx echo.Context) error {
	floorIds, err := utils.GetParamIds(ctx)
	utils.ErrorPanic(err)

	for _, id := range floorIds {
		h.FloorUsecase.Delete(id)
	}

	return utils.WebResponse(
		ctx,
		http.StatusOK,
		"Floor deleted successfully",
		nil,
	)
}

func (h *FloorHttpHandler) GetAllFloor(ctx echo.Context) error {
	floors := h.FloorUsecase.FindAll()

	return utils.WebResponse(
		ctx,
		http.StatusOK,
		"Floor retrieved successfully",
		floors,
	)
}

func (h *FloorHttpHandler) GetFloorById(ctx echo.Context) error {
	floorIds, err := utils.GetParamIds(ctx)
	utils.ErrorPanic(err)

	var floors []models.FloorResponse
	for _, id := range floorIds {
		floor := h.FloorUsecase.FindById(id)
		floors = append(floors, floor)
	}

	return utils.WebResponse(
		ctx,
		http.StatusOK,
		"Floor retrieved successfully",
		floors,
	)
}
