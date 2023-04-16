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

// CreatePhoto godoc
// @ID create-photo
// @Summary Single photo
// @Description Create single photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param PhotoReq body models.PhotoReq true "PhotoReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.PhotoReq{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	photo := models.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserID:   uint(userData["id"].(float64)),
	}

	err = db.Create(&photo).Error
	if err != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, photo)
}

// GetPhoto godoc
// @ID get-photo
// @Summary All photos
// @Description get all photos
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos [get]
func GetPhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photos := []models.Photo{}

	result := db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&photos)
	if result.Error != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, result.Error)
		return
	}

	helpers.SuccessResponse(ctx, photos)
}

// GetPhotoById godoc
// @ID get-photo-by-id
// @Summary Single photo
// @Description get single photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "Photo ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos/{ID} [get]
func GetPhotoById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	photo := models.Photo{}
	res := db.First(&photo, ID)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, photo)
}

// UpdatePhoto godoc
// @ID update-photo
// @Summary Single photo
// @Description update single photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "Photo ID"
// @Param PhotoReq body models.PhotoReq true "PhotoReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos/{ID} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.PhotoReq{}
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

	photo := models.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserID:   uint(userData["id"].(float64)),
	}

	photo.ID = uint(ID)

	res := db.Model(&photo).Where("id=?", ID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL})
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, photo)
}

// DeletePhoto godoc
// @ID delete-photo
// @Summary Single photo
// @Description delete single photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "Photo ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /photos/{ID} [delete]
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res := db.Delete(&photo, ID)
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
