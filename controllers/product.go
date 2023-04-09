package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Digisata/dts-hactiv8-golang-chap3/database"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	products := []models.Product{}
	var result *gorm.DB

	if userData["role"] == "admin" {
		result = db.Find(&products)
	}

	if userData["role"] == "user" {
		result = db.Where("user_id = ?", uint(userData["id"].(float64))).Find(&products)
	}

	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func GetProductById(ctx *gin.Context) {
	db := database.GetDB()
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product := models.Product{}
	err = db.First(&product, productID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.ID = uint(productID)
	product.UserID = uint(userData["id"].(float64))

	res := db.Model(&product).Where("id=?", productID).Updates(models.Product{Title: product.Title, Description: product.Description})
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Product with ID %d not found", productID),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	product := models.Product{}
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := db.Delete(&product, productID)
	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Product with ID %d not found", productID),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
