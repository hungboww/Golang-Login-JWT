package middlleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
)

func CheckLogin(c *gin.Context) {
	tokenValue, err := c.Cookie("Authorization")
	signingKey := []byte(os.Getenv("SECRET"))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if token.Valid {
		fmt.Println("ahihi")
	}
	fmt.Println("token: ", token)
	fmt.Println(jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU", func(token *jwt.Token) (interface{}, error) {
		return tokenValue, nil
	}))
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	fmt.Println("ahihi")
	//	// check time exe
	//	if float64(time.Now().Unix()) > claims["exe"].(float64) {
	//		c.AbortWithStatus(http.StatusUnauthorized)
	//	}
	//	var user model.User
	//	ax := database.DB.First(&user, claims["sub"])
	//	fmt.Println(ax)
	//	if user.ID == 0 {
	//		c.AbortWithStatus(http.StatusUnauthorized)
	//	}
	//
	//	c.Set("user", user)
	//	c.Next()
	//
	//	fmt.Println(claims["foo"], claims["nbf"])
	//} else {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//}

}
