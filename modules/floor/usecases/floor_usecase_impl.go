package usecases

import (
	"sort"

	"github.com/go-playground/validator/v10"
	"github.com/maxmurr/go-clean-arch/internal/utils"
	"github.com/maxmurr/go-clean-arch/modules/floor/entities"
	"github.com/maxmurr/go-clean-arch/modules/floor/models"
	"github.com/maxmurr/go-clean-arch/modules/floor/repositories"
)

type FloorUsecaseImpl struct {
	FloorRepository repositories.FloorRepository
	Validate        *validator.Validate
}

func NewFloorUsecaseImpl(floorRepository repositories.FloorRepository, validate *validator.Validate) FloorUsecase {
	return &FloorUsecaseImpl{FloorRepository: floorRepository, Validate: validate}
}

// Create implements FloorUsecase.
func (f *FloorUsecaseImpl) Create(floor models.CreateFloorRequest) {
	err := f.Validate.Struct(floor)
	utils.ErrorPanic(err)
	floorData := entities.Floor{
		Name: floor.Name,
	}

	f.FloorRepository.Save(floorData)
}

// Delete implements FloorUsecase.
func (f *FloorUsecaseImpl) Delete(floorId uint) {
	f.FloorRepository.Delete(floorId)
}

// FindAll implements FloorUsecase.
func (f *FloorUsecaseImpl) FindAll() []models.FloorResponse {
	result := f.FloorRepository.FindAll()
	sort.Slice(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})

	var floorsResponse []models.FloorResponse
	for _, floor := range result {
		floorsResponse = append(floorsResponse, models.FloorResponse{
			Id:   floor.Id,
			Name: floor.Name,
		})
	}

	return floorsResponse
}

// FindById implements FloorUsecase.
func (f *FloorUsecaseImpl) FindById(floorId uint) models.FloorResponse {
	result, err := f.FloorRepository.FindById(floorId)
	utils.ErrorPanic(err)

	floorResponse := models.FloorResponse{
		Id:   result.Id,
		Name: result.Name,
	}

	return floorResponse
}

// Update implements FloorUsecase.
func (f *FloorUsecaseImpl) Update(floorData models.UpdateFloorRequest) {
	err := f.Validate.Struct(floorData)
	utils.ErrorPanic(err)

	floor := entities.Floor{
		Id:   floorData.Id,
		Name: floorData.Name,
	}

	f.FloorRepository.Update(floor)
}
