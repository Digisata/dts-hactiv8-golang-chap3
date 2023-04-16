package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model     `swaggerignore:"true"`
	Name           string `gorm:"not null" json:"name" valid:"required~Name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint   `json:"user_id"`
	User           *User  `json:"user"`
}

type SocialMediaReq struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
