package handler

import (
	"fmt"
	"net/http"
	"os"
	"user-api/config"
	"user-api/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

	r.Use(authMiddleware)

	r.GET("/user", handler.GetUser)
	r.POST("/user", handler.CreateUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
}

func authMiddleware(c *gin.Context) {
	tokenString := requestHeader{}
	if err := c.ShouldBindHeader(&tokenString); err != nil {
		c.JSON(200, err)
	}
	token, err := jwt.Parse(tokenString.token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("ilham"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}

// Login ...
func (u *UserHandler) Login(c *gin.Context) {
	var user domain.User
	u.Conn.Where("username = ? AND password = ?", c.PostForm("username"), c.PostForm("password")).Find(&user)

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
