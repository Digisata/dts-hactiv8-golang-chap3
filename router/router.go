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

	photoRouter := r.Group("photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.GET("/:ID", middlewares.Authorization("Photo"), controllers.GetPhotoById)
		photoRouter.PUT("/:ID", middlewares.Authorization("Photo"), controllers.UpdatePhoto)
		photoRouter.DELETE("/:ID", middlewares.Authorization("Photo"), controllers.DeletePhoto)
	}

	socialMediaRouter := r.Group("social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.GET("/:ID", middlewares.Authorization("SocialMedia"), controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:ID", middlewares.Authorization("SocialMedia"), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:ID", middlewares.Authorization("SocialMedia"), controllers.DeleteSocialMedia)
	}

	commentRouter := r.Group("comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/:photoID", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.GET("/:ID", middlewares.Authorization("Comment"), controllers.GetCommentById)
		commentRouter.PUT("/:ID", middlewares.Authorization("Comment"), controllers.UpdateComment)
		commentRouter.DELETE("/:ID", middlewares.Authorization("Comment"), controllers.DeleteComment)
	}

	return r
}
