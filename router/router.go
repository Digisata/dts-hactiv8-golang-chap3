package router

import (
	"github.com/Digisata/dts-hactiv8-golang-chap3/controllers"
	"github.com/Digisata/dts-hactiv8-golang-chap3/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")

	userRouter := v1.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	photoRouter := v1.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.GET("/:ID", controllers.GetPhotoById)
		photoRouter.PUT("/:ID", middlewares.Authorization, controllers.UpdatePhoto)
		photoRouter.DELETE("/:ID", middlewares.Authorization, controllers.DeletePhoto)
		photoRouter.GET("/:ID/comments", middlewares.Authorization, controllers.GetComment)
	}

	socialMediaRouter := v1.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.GET("/:ID", controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:ID", middlewares.Authorization, controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:ID", middlewares.Authorization, controllers.DeleteSocialMedia)
	}

	commentRouter := v1.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication)
		commentRouter.POST("/:photoID", controllers.CreateComment)
		// commentRouter.GET("/", controllers.GetComment)
		commentRouter.GET("/:ID", controllers.GetCommentById)
		commentRouter.PUT("/:ID", middlewares.Authorization, controllers.UpdateComment)
		commentRouter.DELETE("/:ID", middlewares.Authorization, controllers.DeleteComment)
	}

	return r
}
