package main

import (
	"gin-gonic/controllers"
	"gin-gonic/database"
	"gin-gonic/initializers"
	"gin-gonic/middlleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDatabase()
	initializers.Migration()
}
func main() {
	r := gin.Default()
	database.ConnectDatabase()

	api := r.Group("api")
	{
		api.GET("/", controllers.GetBlog)
		api.GET("/:id", controllers.ProductDetail)
		api.POST("/", controllers.AddProduct)
		api.PATCH("/:id", controllers.EditProduct)
		api.POST("/login", controllers.Login)
		api.POST("/create", controllers.Register)
		api.GET("/info", middlleware.CheckLogin, controllers.GetInforUser)
		api.POST("/post", middlleware.CheckLogin, controllers.AddForum)
	}

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
