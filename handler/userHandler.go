package handler

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"user-api/config"
	"user-api/domain"
	"user-api/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	r.GET("/user", helper.IsAuthorized, handler.GetUser)
	r.POST("/user", helper.IsAuthorized, handler.CreateUser)
	r.PUT("/user/:id", helper.IsAuthorized, handler.UpdateUser)
	r.DELETE("/user/:id", helper.IsAuthorized, handler.DeleteUser)
}

// Login ...
func (u *UserHandler) Login(c *gin.Context) {
	var user domain.User
	hashPassword := md5.New()
	hashPassword.Write([]byte(c.PostForm("password")))
	isLogin := u.Conn.Where("username = ? AND password = ?", c.PostForm("username"), hex.EncodeToString(hashPassword.Sum(nil))).First(&user)

	if isLogin.Error != nil {
		// fmt.Println(isLogin)
		c.JSON(400, gin.H{
			"message": "Periksa Login Anda",
		})

		c.Abort()
		return
	}

	helper.GenerateJWT(user.Username, c)
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

	if c.Request.Method != "POST" {
		c.JSON(204, gin.H{
			"message": "Method not allowed",
		})
		c.Abort()
		return
	}

	file, null := c.FormFile("foto")

	if null != nil {
		c.JSON(400, gin.H{
			"message": "Images must been added",
		})
		c.Abort()
		return
	}

	var isImages bool = helper.ImageValidation(file.Header.Get("Content-Type"))

	// fmt.Println(file.Header.Get("Content-Type"))
	// fmt.Println(isImages)

	if !isImages {
		c.JSON(400, gin.H{
			"message": "Images not valid",
		})
		c.Abort()
		return
	}
	// isValidImage := helper.MimeFromIncipit(file.Filename)

	// if isValidImage == "" {
	// 	c.JSON(400, gin.H{
	// 		"message": "error",
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// var k domain.User
	// if err := c.ShouldBind(&k); err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	c.Abort()
	// 	return
	// }

	// Set Folder untuk menyimpan filenya
	path := "photos/" + file.Filename

	hashPassword := md5.New()
	hashPassword.Write([]byte(c.PostForm("password")))

	user := domain.User{
		NamaLengkap: c.PostForm("nama_lengkap"),
		Username:    c.PostForm("username"),
		Password:    hex.EncodeToString(hashPassword.Sum(nil)),
		Foto:        file.Filename,
	}

	errs := validator.New().Struct(user)
	if errs != nil {
		var formError = map[string]string{}
		for _, err := range errs.(validator.ValidationErrors) {
			formError[err.Field()] = err.Tag()
			// fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			// fmt.Println(err.Field())     // by passing alt name to ReportError like below
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.)
		}
		c.JSON(400, gin.H{
			"error": formError,
		})

		c.Abort()
		return
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
	if oldUser.NamaLengkap == "" {
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		c.Abort()
		return
	}

	file, null := c.FormFile("foto")

	if null != nil {
		c.JSON(400, gin.H{
			"message": "Images must been added",
		})
		c.Abort()
		return
	}

	var isImages bool = helper.ImageValidation(file.Header.Get("Content-Type"))

	// fmt.Println(file.Header.Get("Content-Type"))
	// fmt.Println(isImages)

	if !isImages {
		c.JSON(400, gin.H{
			"message": "Images not valid",
		})
		c.Abort()
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "photos/" + file.Filename

	hashPassword := md5.New()
	hashPassword.Write([]byte(c.PostForm("password")))

	user := domain.User{
		NamaLengkap: c.PostForm("nama_lengkap"),
		Username:    c.PostForm("username"),
		Password:    hex.EncodeToString(hashPassword.Sum(nil)),
		Foto:        file.Filename,
	}

	errs := validator.New().Struct(user)
	if errs != nil {
		c.JSON(400, gin.H{
			"error": errs.Error(),
		})

		c.Abort()
		return
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
	if c.Request.Method != "DELETE" {
		c.JSON(204, gin.H{
			"message": "Method not allowed",
		})
		c.Abort()
		return
	}

	isDeleted, _ := u.UserEntity.Delete(c.Param("id"))
	if isDeleted.NamaLengkap == "" {
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		c.Abort()
		return
	}

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
