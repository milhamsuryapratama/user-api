package domain

import (
	"context"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Username    string `tag:"username" json:"username" form:"username" validate:"min=3"`
	Password    string `tag:"password" json:"password" form:"password" validate:"min=6"`
	NamaLengkap string `tag:"nama_lengkap" json:"nama_lengkap" form:"nama_lengkap" validate:"min=3"`
	Foto        string `tag:"foto" json:"foto" form:"foto"`
}

// UserEntity ...
type UserEntity interface {
	Get(ctx context.Context) ([]User, error)
	Create(user User) (User, error)
	Update(id string, user User) (User, error)
	Delete(id string) (User, error)
	Show(id string) (User, error)
	// Login(username string, password string) ()
}

// UserRepository ...
type UserRepository interface {
	Get(ctx context.Context) (res []User, err error)
	Create(user User) (User, error)
	Update(id string, user User) (User, error)
	Delete(id string) (User, error)
	Show(id string) (User, error)
}
