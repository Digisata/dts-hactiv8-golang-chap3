package controllers

import (
	"errors"
	"net/http"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create-user
// @Summary Register User
// @Description Create new user
// @Tags User
// @Accept json
// @Produce json
// @Param RegisterReq body models.RegisterReq true "RegisterReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /users/register [post]
func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	payload := models.RegisterReq{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
	}

	err = db.Create(&user).Error
	if err != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, user)
}

// LoginUser godoc
// @ID login-user
// @Summary Login User
// @Description Login with existing user
// @Tags User
// @Accept json
// @Produce json
// @Param LoginReq body models.LoginReq true "LoginReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /users/login [post]
func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	payload := models.LoginReq{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	password := payload.Password

	user := models.User{}

	res := db.Where("email = ?", payload.Email).Take(&user)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, errors.New("invalid email or password"))
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if !helpers.PasswordValid(user.Password, password) {
		helpers.FailResponse(ctx, http.StatusNotFound, errors.New("invalid email or password"))
		return
	}

	token, err := helpers.GenerateToken(user.ID, payload.Email)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, gin.H{
		"token": token,
	})
}
