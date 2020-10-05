package domain

import (
	"context"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Username    string
	Password    string
	NamaLengkap string
	Foto        string
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
