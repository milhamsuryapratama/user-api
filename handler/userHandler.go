package handler

import (
	"net/http"
	"os"
	"user-api/config"
	"user-api/domain"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler ...
type UserHandler struct {
	UserEntity domain.UserEntity
	Conn       *gorm.DB
}

// Credential ...
type Credential struct {
	Username string
	Password string
}

// UserHandlerFunc ...
func UserHandlerFunc(r *gin.RouterGroup, user domain.UserEntity) {
	handler := &UserHandler{
		UserEntity: user,
		Conn:       config.Connect(),
	}

	r.GET("/user", handler.GetUser)
	r.POST("/user", handler.CreateUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
	r.POST("login", handler.Login)
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
