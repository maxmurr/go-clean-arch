package entities

import "github.com/maxmurr/go-clean-arch/internal/utils"

type (
	Floor struct {
		Id   uint   `gorm:"primary_key;auto_increment;not null" json:"id"`
		Name string `gorm:"unique;not null" json:"name"`
		utils.BaseModel
	}
)
