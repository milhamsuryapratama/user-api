package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"user-api/config"
	"user-api/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// UserHandler ...
type UserHandler struct {
	UserEntity domain.UserEntity
	Conn       *gorm.DB
}

type requestHeader struct {
	token string
}

// UserHandlerFunc ...
func UserHandlerFunc(r *gin.RouterGroup, user domain.UserEntity) {
	handler := &UserHandler{
		UserEntity: user,
		Conn:       config.Connect(),
	}

	r.POST("login", handler.Login)

	// r.Use(isAuthorized)

	r.GET("/user", isAuthorized, handler.GetUser)
	r.POST("/user", isAuthorized, handler.CreateUser)
	r.PUT("/user/:id", isAuthorized, handler.UpdateUser)
	r.DELETE("/user/:id", isAuthorized, handler.DeleteUser)
}

func isAuthorized(c *gin.Context) {
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

// Login ...
func (u *UserHandler) Login(c *gin.Context) {
	var user domain.User
	isLogin := u.Conn.Where("username = ? AND password = ?", c.PostForm("username"), c.PostForm("password")).First(&user)

	if user.Username == "" {
		fmt.Println(isLogin)
		c.JSON(400, gin.H{
			"message": "Periksa Login Anda",
		})

		c.Abort()
		return
	}

	GenerateJWT(user.Username, c)
}

// GetUser ...
func (u *UserHandler) GetUser(c *gin.Context) {
	users, _ := u.UserEntity.Get(c)
	c.JSON(200, gin.H{
		"data": users,
	})
}

// CreateUser ...
func (u *UserHandler) CreateUser(c *gin.Context) {
	file, _ := c.FormFile("foto")

	// isValidImage := helper.MimeFromIncipit([]byte(file))

	// if isValidIMage == "" {
	// 	c.JSON(201, gin.H{
	// 		"message": "error",
	// 	})
	// 	c.Abort()
	// 	return
	// }

	var k domain.User
	if err := c.ShouldBind(&k); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "photos/" + file.Filename

	user := domain.User{
		NamaLengkap: c.PostForm("nama_lengkap"),
		Username:    c.PostForm("username"),
		Password:    c.PostForm("password"),
		Foto:        file.Filename,
	}

	c.SaveUploadedFile(file, path)

	isCreated, _ := u.UserEntity.Create(user)
	c.JSON(201, gin.H{
		"data": isCreated,
	})
}

// UpdateUser ...
func (u *UserHandler) UpdateUser(c *gin.Context) {
	oldUser, _ := u.UserEntity.Show(c.Param("id"))

	file, _ := c.FormFile("foto")

	// Set Folder untuk menyimpan filenya
	path := "photos/" + file.Filename

	user := domain.User{
		NamaLengkap: c.PostForm("nama_lengkap"),
		Username:    c.PostForm("username"),
		Password:    c.PostForm("password"),
		Foto:        file.Filename,
	}

	oldFile := "photos/" + oldUser.Foto
	os.Remove(oldFile)
	c.SaveUploadedFile(file, path)

	isUpdated, _ := u.UserEntity.Update(c.Param("id"), user)

	c.JSON(200, gin.H{
		"data": isUpdated,
	})
}

// DeleteUser ...
func (u *UserHandler) DeleteUser(c *gin.Context) {
	isDeleted, _ := u.UserEntity.Delete(c.Param("id"))
	c.JSON(200, gin.H{
		"data": isDeleted,
	})
}

// ShowUser ...
func (u *UserHandler) ShowUser(c *gin.Context) {
	user, _ := u.UserEntity.Show(c.Param("id"))
	c.JSON(200, gin.H{
		"data": user,
	})
}
