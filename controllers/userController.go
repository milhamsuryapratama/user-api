package controllers

import (
	"context"
	"user-api/domain"
)

// UserEntity ...
type UserEntity struct {
	userRepo domain.UserRepository
}

// NewUserEntity ...
func NewUserEntity(u domain.UserRepository) domain.UserEntity {
	return &UserEntity{
		userRepo: u,
	}
}

// Get ...
func (k *UserEntity) Get(ctx context.Context) (res []domain.User, err error) {
	res, err = k.userRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return
}

// Create ...
func (k *UserEntity) Create(usr domain.User) (user domain.User, err error) {
	user, err = k.userRepo.Create(usr)
	if err != nil {
		return user, nil
	}

	return
}

// Update ...
func (k *UserEntity) Update(id string, usr domain.User) (user domain.User, err error) {
	user, err = k.userRepo.Update(id, usr)
	return
}

// Delete ...
func (k *UserEntity) Delete(id string) (user domain.User, err error) {
	user, err = k.userRepo.Delete(id)
	return
}

// Show ...
func (k *UserEntity) Show(id string) (user domain.User, err error) {
	user, err = k.userRepo.Show(id)
	return
}
