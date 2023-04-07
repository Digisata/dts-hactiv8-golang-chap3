package router

import (
	"github.com/Digisata/dts-hactiv8-golang-chap3/controllers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetProduct)
		productRouter.PUT("/:productID", middlewares.AdminAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.AdminAuthorization(), controllers.DeleteProduct)
		// productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
