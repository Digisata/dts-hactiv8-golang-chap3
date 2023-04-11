package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" validate:"required-Name is required"`
	SocialMediaURL string `json:"social_media_url" validate:"required-Social Media URL is required"`
	UserID         uint   `json:"user_id"`
	User           *User  `json:"user"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
