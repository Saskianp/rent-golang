package repository

import (
	"testBeGo/models"
	"gorm.io/gorm"
)

type DriverRepository struct {
	DB *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepository {
	return &DriverRepository{DB: db}
}

func (r *DriverRepository) GetAll() ([]models.Driver, error) {
	var driver []models.Driver
	err := r.DB.Find(&driver).Error
	return driver, err
}

func (r *DriverRepository) GetByID(id uint) (models.Driver, error) {
	var driver models.Driver
	err := r.DB.First(&driver, id).Error
	return driver, err
}

func (r *DriverRepository) Create(driver *models.Driver) error {
	return r.DB.Create(driver).Error
}

func (r *DriverRepository) Update(id uint, updatedData *models.Driver) error {
	return r.DB.Model(&models.Driver{}).Where("id = ?", id).Updates(updatedData).Error
}

func (r *DriverRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Driver{}, id).Error
}
