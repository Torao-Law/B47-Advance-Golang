package repository

import "dumbmerch/models"

type AuthRepository interface {
	Register(user models.User) (models.User, error)
}

func (r *repository) Register(users models.User) (models.User, error) {
	err := r.db.Create(&users).Error

	return users, err
}
