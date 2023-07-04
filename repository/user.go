package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	// SELECT * FROM users WHERE id = 2

	return user, err
}

func (r *repository) CreateUser(users models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO `users`(`name`, `email`, `password`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)", users.Name, users.Email, users.Password, users.CreatedAt, users.UpdatedAt).Error

	return users, err
}

func (r *repository) UpdateUser(users models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE `users` SET `name`=?,`email`=?,`password`=?,`updated_at`=? WHERE id = ?", users.Name, users.Email, users.Password, users.UpdatedAt, ID).Scan(&users).Error

	return users, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
