package models

type CreateFloorRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateFloorRequest struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type FloorResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
