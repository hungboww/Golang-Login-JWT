package initializers

import (
	"gin-gonic/database"
	"gin-gonic/model"
)

func Migration() {
	database.DB.AutoMigrate(&model.Product{}, &model.User{})
}
