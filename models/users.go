package models

import (
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" valid:"required~Age is required,range(9|200)~Age have to be greater than or equal to 9"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
