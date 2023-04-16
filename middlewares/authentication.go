package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(c *gin.Context) {
	claims, err := helpers.VerifyToken(c)
	if err != nil {
		helpers.FailResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Set("userData", claims)
	c.Next()
}

func Authorization(c *gin.Context) {
	db := database.GetDB()
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, err)
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var res *gorm.DB
	var domainUserID uint

	path := c.Request.URL.Path
	domain := ""

	if strings.Contains(path, "photos") && !strings.Contains(path, "comments") {
		model := models.Photo{}
		res = db.Select("user_id").First(&model, uint(ID))
		domainUserID = model.UserID
		domain = "Photo"
	} else if strings.Contains(path, "social-media") {
		model := models.SocialMedia{}
		res = db.Select("user_id").First(&model, uint(ID))
		domainUserID = model.UserID
		domain = "Social Media"
	} else {
		model := models.Comment{}
		res = db.Select("user_id").First(&model, uint(ID))
		domainUserID = model.UserID
		domain = "Comment"
	}

	if res.Error != nil {
		if res.RowsAffected == 0 {
			helpers.FailResponse(c, http.StatusNotFound, fmt.Errorf("%s with ID %d not found", domain, ID))
			return
		}

		helpers.FailResponse(c, http.StatusInternalServerError, err)
		return
	}

	if domainUserID != userID {
		helpers.FailResponse(c, http.StatusUnauthorized, fmt.Errorf("you are not allowed to acces %s with ID %d", domain, ID))
		return
	}

	c.Next()
}
