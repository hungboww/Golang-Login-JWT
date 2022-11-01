package middlleware

import (
	"fmt"
	"gin-gonic/database"
	"gin-gonic/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func CheckLogin(c *gin.Context) {
	tokenValue, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check time exe
		if float64(time.Now().Unix()) > claims["exe"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	var user model.User
	database.DB.First(&user)
	c.Next()

}
