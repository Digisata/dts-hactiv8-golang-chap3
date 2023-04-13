package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" valid:"optional"`
	PhotoURL string `gorm:"not null" json:"photo_url" valid:"required~Photo URL is required"`
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
