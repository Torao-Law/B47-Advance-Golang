package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type Profile interface {
	GetProfile(ID int) (models.Profile, error)
}

func ProfileRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}
