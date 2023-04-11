package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo.UserID = uint(userData["id"].(float64))

	err = db.Create(&photo).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func GetPhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photos := []models.Photo{}

	result := db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&photos)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func GetPhotoById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo := models.Photo{}
	err = db.First(&photo, ID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photo := models.Photo{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo.ID = uint(ID)
	photo.UserID = uint(userData["id"].(float64))

	res := db.Model(&photo).Where("id=?", ID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL})
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Photo with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := db.Delete(&photo, ID)
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Photo with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
