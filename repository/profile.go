package repository

import (
	"dumbmerch/models"
)

type Profile interface {
	GetProfile(ID int) (models.Profile, error)
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}
