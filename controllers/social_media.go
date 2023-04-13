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

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	err = db.Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)
}

func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedias := []models.SocialMedia{}

	result := db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&socialMedias)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

func GetSocialMediaById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	socialMedia := models.SocialMedia{}
	err = db.First(&socialMedia, ID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	socialMedia.ID = uint(ID)
	socialMedia.UserID = uint(userData["id"].(float64))

	res := db.Model(&socialMedia).Where("id=?", ID).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaURL: socialMedia.SocialMediaURL})
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": fmt.Sprintf("SocialMedia with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	res := db.Delete(&socialMedia, ID)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": fmt.Sprintf("SocialMedia with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
