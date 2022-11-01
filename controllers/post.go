package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AddForum(c *gin.Context) {
	var FormData struct {
		UserID      int
		Description string
		Title       string
		View        int
		Stt         int
	}

	if c.Bind(&FormData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check for "})
		return
	}
	//var forum model.Forum
	uid := os.Getuid()
	fmt.Println("Id", uid)
}
