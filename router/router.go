package router

import (
	"hacktiv8-golang-final-project/controllers"
	"hacktiv8-golang-final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

		userRouter.Use(middlewares.Authentication())

		userRouter.PUT("/", controllers.UserUpdate)
		userRouter.DELETE("/", controllers.UserDelete)
	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())

		photosRouter.POST("/", controllers.PhotoCreate)
		photosRouter.GET("/", controllers.PhotoGetAll)
		photosRouter.PUT("/:photoId", controllers.PhotoUpdate)
		photosRouter.DELETE("/:photoId", controllers.PhotoDelete)
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())

		commentsRouter.POST("/", controllers.CommentCreate)
		commentsRouter.GET("/", controllers.CommentGetAll)
		commentsRouter.PUT("/:commentId", controllers.CommentUpdate)
		commentsRouter.DELETE("/:commentId", controllers.CommentDelete)
	}

	socialMediasRouter := r.Group("/socialMedias")
	{
		socialMediasRouter.Use(middlewares.Authentication())

		socialMediasRouter.POST("/", controllers.SocialMediasCreate)
		socialMediasRouter.GET("/", controllers.SocialMediaGetAll)
		socialMediasRouter.PUT("/:socialMediaId", controllers.SocialMediaUpdate)
		socialMediasRouter.DELETE("/:socialMediaId", controllers.SocialMediaDelete)
	}

	return r
}
