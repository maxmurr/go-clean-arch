package repositories

import "github.com/maxmurr/go-clean-arch/modules/floor/entities"

type FloorRepository interface {
	FindAll() ([]entities.Floor)
	FindById(floorId uint) (floor entities.Floor, err error)
	Save(floor entities.Floor)
	Update(floor entities.Floor)
	Delete(floorId uint)
}
