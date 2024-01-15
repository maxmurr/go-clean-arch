package repositories

import (
	"github.com/maxmurr/go-clean-arch/internal/utils"
	"github.com/maxmurr/go-clean-arch/modules/floor/entities"
	"gorm.io/gorm"
)

type FloorRepositoryImpl struct {
	db *gorm.DB
}

func NewFloorRepositoryImpl(db *gorm.DB) FloorRepository {
	return &FloorRepositoryImpl{db: db}
}

// Delete implements FloorRepository.
func (f *FloorRepositoryImpl) Delete(floorId uint) {
	var floor entities.Floor
	result := f.db.Where("id = ?", floorId).Delete(&floor)
	utils.ErrorPanic(result.Error)
}

// FindAll implements FloorRepository.
func (f *FloorRepositoryImpl) FindAll() []entities.Floor {
	var floors []entities.Floor
	result := f.db.Find(&floors)
	utils.ErrorPanic(result.Error)
	return floors
}

// FindById implements FloorRepository.
func (f *FloorRepositoryImpl) FindById(floorId uint) (floor entities.Floor, err error) {
	result := f.db.Find(&floor, floorId)
	if result.Error != nil {
		return floor, result.Error
	}
	return floor, nil
}

// Save implements FloorRepository.
func (f *FloorRepositoryImpl) Save(floor entities.Floor) {
	result := f.db.Create(&floor)
	utils.ErrorPanic(result.Error)
}

// Update implements FloorRepository.
func (f *FloorRepositoryImpl) Update(floor entities.Floor) {
	var floorData = entities.Floor{
		Name: floor.Name,
	}
	result := f.db.Model(&floor).Updates(floorData)
	utils.ErrorPanic(result.Error)
}
