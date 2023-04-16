package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model `swaggerignore:"true"`
	Message    string `gorm:"not null" json:"message" valid:"required~Message is required" example:"string"`
	UserID     uint   `json:"user_id"`
	User       *User  `json:"user"`
	PhotoID    uint   `json:"photo_id"`
	Photo      *Photo `json:"photo"`
}

type CommentReq struct {
	Message string `json:"message"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
