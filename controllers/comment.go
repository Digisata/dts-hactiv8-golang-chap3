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

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}
	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	comment.UserID = uint(userData["id"].(float64))
	comment.PhotoID = uint(photoID)

	err = db.Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	result := db.Where("photo_id = ?", ID).Find(&comments)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func GetCommentById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	comment := models.Comment{}
	err = db.First(&comment, ID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	comment.ID = uint(ID)
	comment.UserID = uint(userData["id"].(float64))

	res := db.Model(&comment).Where("id=?", ID).Updates(models.Comment{Message: comment.Message})
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": fmt.Sprintf("Comment with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	res := db.Delete(&comment, ID)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": fmt.Sprintf("Comment with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
