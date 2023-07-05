package repository

import "dumbmerch/models"

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var Products []models.Product
	err := r.db.Preload("User").Find(&Products).Error

	return Products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var Product models.Product
	err := r.db.Preload("User").First(&Product, ID).Error

	return Product, err
}

func (r *repository) CreateProduct(Product models.Product) (models.Product, error) {
	err := r.db.Create(&Product).Error

	return Product, err
}
