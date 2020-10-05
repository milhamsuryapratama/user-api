package main

import (
	"user-api/config"
	"user-api/controllers"
	"user-api/handler"
	"user-api/repositories"

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
	r := gin.Default()
	db := config.Connect()

	repo := repositories.NewUserRepository(db)
	entity := controllers.NewUserEntity(repo)

	api := r.Group("/api")
	// api.POST("/login", Login)
	// api.Use(authMiddleware)
	// {
	// 	handler.UserHandlerFunc(api, entity)
	// }
	handler.UserHandlerFunc(api, entity)
	r.Run()
}

// func authMiddleware(c *gin.Context) {
// 	tokenString := requestHeader{}
// 	if err := c.ShouldBindHeader(&tokenString); err != nil {
// 		c.JSON(200, err)
// 	}
// 	token, err := jwt.Parse(tokenString.token, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte("ilham"), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println(claims["foo"], claims["nbf"])
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// Login ...
// // func Login(c *gin.Context) {
// 	var user domain.User
// 	isLogin := u.Con.Where("username = ? AND password = ?", c.PostForm("username"), c.PostForm("password")).Find(&user)

// 	if isLogin != nil {
// 		return
// 	}

// 	sign := jwt.New(jwt.GetSigningMethod("HS256"))
// 	token, error := sign.SignedString([]byte("secret"))
// 	if error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": error.Error(),
// 		})
// 		c.Abort()
// 	}
// 	c.JSON(200, gin.H{
// 		"token": token,
// 	})
// }
