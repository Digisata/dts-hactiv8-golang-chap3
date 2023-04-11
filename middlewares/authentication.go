package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func Authorization(table string) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		var res *gorm.DB
		var domainUserID uint

		switch table {
		case "Photo":
			domain := models.Photo{}
			res = db.Select("user_id").First(&domain, uint(ID))
			domainUserID = domain.UserID
		case "SocialMedia":
			domain := models.SocialMedia{}
			res = db.Select("user_id").First(&domain, uint(ID))
			domainUserID = domain.UserID
		case "Comment":
			domain := models.Comment{}
			res = db.Select("user_id").First(&domain, uint(ID))
			domainUserID = domain.UserID
		}

		if res.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("%s with ID %d not found", table, ID),
			})
			return
		}

		if domainUserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   fmt.Sprintf("You are not allowed to access this %s", table),
			})
			return
		}

		c.Next()
	}
}
