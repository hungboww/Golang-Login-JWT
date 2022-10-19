package controllers

import (
	"gin-gonic/database"
	"gin-gonic/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	var FormData struct {
		Email    string
		Password string
	}
	if c.Bind(&FormData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check for "})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(FormData.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check for password"})
		return
	}
	//create
	user := model.User{
		Email:    FormData.Email,
		Password: string(hash),
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No create account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func Login(c *gin.Context) {
	var FormData struct {
		Email    string
		Password string
	}
	if c.Bind(&FormData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check for "})
		return
	}
	var user model.User
	database.DB.First(&user, "Email=?", FormData.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No account "})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(FormData.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Check email of password"})
	}
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user.ID",
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error create token"})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
