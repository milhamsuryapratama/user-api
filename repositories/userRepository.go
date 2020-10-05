package repositories

import (
	"context"
	"user-api/domain"

	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &UserRepository{Conn}
}

// Get ...
func (r *UserRepository) Get(ctx context.Context) (res []domain.User, err error) {
	var users []domain.User
	r.Conn.Find(&users)
	return users, nil
}

// Create ...
func (r *UserRepository) Create(usr domain.User) (user domain.User, err error) {
	r.Conn.Create(&usr)
	return usr, nil
}

// Update ...
func (r *UserRepository) Update(id string, usr domain.User) (user domain.User, err error) {
	var us domain.User
	r.Conn.First(&us, id)
	us.NamaLengkap = usr.NamaLengkap
	us.Username = usr.Username
	us.Password = usr.Password
	us.Foto = usr.Foto
	r.Conn.Save(&us)
	return us, nil
}

// Delete ...
func (r *UserRepository) Delete(id string) (user domain.User, err error) {
	var usr domain.User
	r.Conn.Delete(&usr, id)
	return usr, nil
}

// Show ...
func (r *UserRepository) Show(id string) (user domain.User, err error) {
	var usr domain.User
	r.Conn.First(&usr, id)
	return usr, nil
}
