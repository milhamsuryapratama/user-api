package main

import (
	"net/http"
	"user-api/config"
	"user-api/controllers"
	"user-api/domain"
	"user-api/handler"
	"user-api/repositories"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestHeader struct {
	token string
}

type db struct {
	Con *gorm.DB
}

func main() {
	r := gin.New()
	db := config.Connect()

	repo := repositories.NewUserRepository(db)
	entity := controllers.NewUserEntity(repo)

	api := r.Group("/api")
	handler.UserHandlerFunc(api, entity)
	r.Run()
}

// Login ...
func Login(c *gin.Context) {
	var user domain.User
	var u *db
	isLogin := u.Con.Where("username = ? AND password = ?", c.PostForm("username"), c.PostForm("password")).Find(&user)

	if isLogin != nil {
		return
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, error := sign.SignedString([]byte("secret"))
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": error.Error(),
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
