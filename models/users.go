package models

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"not null;uniqueIndex" json:"username" valid:"required~Username is required"`
	Email      string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password   string `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age        int    `gorm:"not null" json:"age" valid:"required~Age is required,range(9|200)~Age have to be greater than or equal to 9"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	salt := 10
	arrByte := []byte(u.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(arrByte, salt)
	if err != nil {
		return
	}
	u.Password = string(hashedPass)

	return
}
