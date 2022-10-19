package controllers

import (
	"gin-gonic/database"
	"gin-gonic/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func ViewGonic(c *gin.Context) {
//	var product []model.Product
//	db, err := database.ConnectDatabase()
//	if err != nil {
//		log.Println(err, product, db)
//	}
//	c.JSON(http.StatusOK, gin.H{"data": "Ahihihihihihihihihihihihihihihihihihih"})
//}

func GetBlog(c *gin.Context) {
	var blogs []model.Product
	database.DB.Find(&blogs)
	c.JSON(http.StatusOK, &blogs)
}
func ProductDetail(c *gin.Context) {
	var blogs []model.Product
	database.DB.Find(&blogs)
	c.JSON(http.StatusOK, &blogs)
}
func AddProduct(c *gin.Context) {
	var blogs []model.Product
	database.DB.Find(&blogs)
	c.JSON(http.StatusOK, &blogs)
}
func DeleteProduct(c *gin.Context) {
	var delete_product model.Product
	database.DB.Where("id = ?", c.Param("id")).Delete(&delete_product)
	c.JSON(http.StatusOK, &delete_product)
}
func EditProduct(c *gin.Context) {
	var blogs []model.Product
	database.DB.Find(&blogs)
	c.JSON(http.StatusOK, &blogs)
}
