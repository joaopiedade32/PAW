package main

import (
	"bookapi/controller"
	"bookapi/service"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	service.InitializeBooks()

	v1 := router.Group("/api/v1")
	{
		book := v1.Group("/book")
		{
			book.GET("/", controller.GetAllBooks)
			book.POST("/", controller.InsertBook)
			book.GET("/:id", controller.GetBookById)
			book.PUT("/:id", controller.UpdateBook)
			book.DELETE("/:id", controller.DeleteBook)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers)
			user.POST("/", controller.Register)
			user.GET("/:id", controller.Profile)
			user.PUT("/:id", controller.UpdateProfile)
			user.DELETE("/:id", controller.DeleteAccount)
		}

	}

	router.Run(":3000")
}
