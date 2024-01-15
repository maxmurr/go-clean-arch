package usecases

import (
	"github.com/maxmurr/go-clean-arch/modules/floor/models"
)

type FloorUsecase interface {
	Create(floor models.CreateFloorRequest)
	Update(floor models.UpdateFloorRequest)
	Delete(floorId uint)
	FindAll() []models.FloorResponse
	FindById(floorId uint) models.FloorResponse
}
