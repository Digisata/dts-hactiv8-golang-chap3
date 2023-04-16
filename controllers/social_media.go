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

// CreateSocialMedia godoc
// @ID create-social-media
// @Summary Single social media
// @Description Create single social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param SocialMediaReq body models.SocialMediaReq true "SocialMediaReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /social-media [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.SocialMediaReq{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	socialMedia := models.SocialMedia{
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
		UserID:         uint(userData["id"].(float64)),
	}

	err = db.Create(&socialMedia).Error
	if err != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, socialMedia)
}

// GetSocialMedia godoc
// @ID get-social-media
// @Summary All social-media
// @Description get all social-media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /social-media [get]
func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedias := []models.SocialMedia{}

	res := db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&socialMedias)
	if res.Error != nil {
		helpers.FailResponse(ctx, http.StatusInternalServerError, res.Error)
		return
	}

	helpers.SuccessResponse(ctx, socialMedias)
}

// GetSocialMediaById godoc
// @ID get-social-media-by-id
// @Summary Single social media
// @Description get single social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "SocialMedia ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /social-media/{ID} [get]
func GetSocialMediaById(ctx *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	socialMedia := models.SocialMedia{}
	res := db.First(&socialMedia, ID)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, socialMedia)
}

// UpdateSocialMedia godoc
// @ID update-social-media
// @Summary Single social media
// @Description update single social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "SocialMedia ID"
// @Param SocialMediaReq body models.SocialMediaReq true "SocialMediaReq"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /social-media/{ID} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	payload := models.SocialMediaReq{}
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

	socialMedia := models.SocialMedia{
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
		UserID:         uint(userData["id"].(float64)),
	}

	socialMedia.ID = uint(ID)

	res := db.Model(&socialMedia).Where("id=?", ID).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaURL: socialMedia.SocialMediaURL})
	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(ctx, http.StatusNotFound, res.Error)
			return
		}

		helpers.FailResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.SuccessResponse(ctx, socialMedia)
}

// DeleteSocialMedia godoc
// @ID delete-social-media
// @Summary Single social media
// @Description delete single social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ID path int true  "SocialMedia ID"
// @Success 200 {object} models.SuccessResult
// @Failure 400 {object} models.BadRequestResult
// @Failure 401 {object} models.UnauthorizedResult
// @Failure 404 {object} models.NotFoundResult
// @Failure 500 {object} models.InternalServerErrorResult
// @Router /social-media/{ID} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}
	ID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res := db.Delete(&socialMedia, ID)
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
