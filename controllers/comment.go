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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	comment.UserID = uint(userData["id"].(float64))
	comment.PhotoID = uint(photoID)

	err = db.Create(&comment).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comments := []models.Comment{}

	result := db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&comments)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func GetCommentById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	comment := models.Comment{}
	err = db.First(&comment, ID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	comment.ID = uint(ID)
	comment.UserID = uint(userData["id"].(float64))

	res := db.Model(&comment).Where("id=?", ID).Updates(models.Comment{Message: comment.Message})
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := db.Delete(&comment, ID)
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Comment with ID %d not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
