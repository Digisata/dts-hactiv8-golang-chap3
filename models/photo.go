package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title" validate:"required-Title is required"`
	Caption  string `json:"caption" validate:"required-Caption is required"`
	PhotoURL string `json:"photo_url" validate:"required-Photo URL is required"`
	UserID   uint   `json:"user_id"`
	User     *User  `json:"user"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
