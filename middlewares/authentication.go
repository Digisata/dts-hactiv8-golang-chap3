package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/helpers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func Authorization(table string) gin.HandlerFunc {
	tables := map[string]interface{}{
		"Photo":       models.Photo{},
		"SocialMedia": models.SocialMedia{},
		"Comment":     models.Comment{},
	}

	return func(c *gin.Context) {
		db := database.GetDB()
		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "Unauthorized",
				"message": "Invalid ID data type",
			})
			return
		}
		model, ok := tables[table]
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "Unauthorized",
				"message": fmt.Sprintf("Invalid table %s", table),
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		res := db.Model(model).Select("user_id").First(model, uint(ID))

		var domainUserID uint
		switch v := model.(type) {
		case *models.Photo:
			domainUserID = v.UserID
		case *models.SocialMedia:
			domainUserID = v.UserID
		case *models.Comment:
			domainUserID = v.UserID
		}

		if res.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": fmt.Sprintf("%s with ID %d not found", table, ID),
			})
			return
		}

		if domainUserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "Forbidden",
				"message": fmt.Sprintf("You are not allowed to access this %s", table),
			})
			return
		}

		c.Next()
	}
}
