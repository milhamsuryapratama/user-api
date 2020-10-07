package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// IsAuthorized ...
func IsAuthorized(c *gin.Context) {
	godotenv.Load()
	if c.Request.Header["Token"] != nil {

		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})

			c.Abort()
			return
		}

		if token.Valid {
			// sign := token.Claims.(jwt.MapClaims)
			// fmt.Println(sign["authorized"])
			c.Next()
		}
	} else {

		c.JSON(400, gin.H{
			"message": "Token Invalid",
		})

		c.Abort()
		return
	}
}

// GenerateJWT ...
func GenerateJWT(user string, c *gin.Context) {
	godotenv.Load()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = c.PostForm("username")
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		// fmt.Errorf("Something Went Wrong: %s", err.Error())
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
