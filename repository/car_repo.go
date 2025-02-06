package repository

import (
	"testBeGo/models"

	"gorm.io/gorm"
)

type CarRepository struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{DB: db}
}

func (r *CarRepository) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	err := r.DB.Find(&cars).Error
	return cars, err
}

func (r *CarRepository) GetCarByID(id uint) (*models.Car, error) {
	var car models.Car
	err := r.DB.First(&car, id).Error
	return &car, err
}

func (r *CarRepository) CreateCar(car *models.Car) error {
	return r.DB.Create(car).Error
}

func (r *CarRepository) UpdateCar(id uint, car *models.Car) error {
	return r.DB.Model(&models.Car{}).Where("id = ?", id).Updates(car).Error
}

func (r *CarRepository) DeleteCar(id uint) error {
	return r.DB.Delete(&models.Car{}, id).Error
}
