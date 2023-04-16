package controllers

import (
	"net/http"
	"strconv"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateComment godoc
// @ID create-comment
// @Summary Single comment
// @Description Create single comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photoID path int true  "Photo ID"
// @Param CommentReq body models.CommentReq true "CommentReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /comments/{photoID} [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.CommentReq{}
	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	comment := models.Comment{
		Message: payload.Message,
		UserID:  uint(userData["id"].(float64)),
		PhotoID: uint(photoID),
	}

	err = db.Create(&comment).Error
	if err != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, comment)
}

// GetComment godoc
// @ID get-comment
// @Summary All comments
// @Description get all comments
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true "Photo ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos/{ID}/comments [get]
func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	result := db.Where("photo_id = ?", ID).Find(&comments)
	if result.Error != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, result.Error)
		return
	}

	helpers.SuccessResponse(ctx, comments)
}

// GetCommentById godoc
// @ID get-comment-by-id
// @Summary Single comment
// @Description get single comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true "Comment ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /comments/{ID} [get]
func GetCommentById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	comment := models.Comment{}
	res := db.First(&comment, ID)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, comment)
}

// UpdateComment godoc
// @ID update-comment
// @Summary Single comment
// @Description update single comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true "Comment ID"
// @Param CommentReq body models.CommentReq true "CommentReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /comments/{ID} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.CommentReq{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	comment := models.Comment{
		Message: payload.Message,
		UserID:  uint(userData["id"].(float64)),
	}

	comment.ID = uint(ID)

	res := db.Model(&comment).Where("id=?", ID).Updates(models.Comment{Message: comment.Message})
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, comment)
}

// DeleteComment godoc
// @ID delete-comment
// @Summary Single comment
// @Description delete single comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true "Comment ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /comments/{ID} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res := db.Delete(&comment, ID)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, "Deleted")
}
