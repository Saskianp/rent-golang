package repository

import (
	"testBeGo/models"
	"gorm.io/gorm"
)

type DriverIncentiveRepository struct {
	DB *gorm.DB
}

func NewDriverIncentiveRepository(db *gorm.DB) *DriverIncentiveRepository {
	return &DriverIncentiveRepository{DB: db}
}

func (r *DriverIncentiveRepository) GetAll() ([]models.DriverIncentive, error) {
	var driverIncentive []models.DriverIncentive
	err := r.DB.Find(&driverIncentive).Error
	return driverIncentive, err
}

func (r *DriverIncentiveRepository) GetByID(id uint) (models.DriverIncentive, error) {
	var driverIncentive models.DriverIncentive
	err := r.DB.First(&driverIncentive, id).Error
	return driverIncentive, err
}

func (r *DriverIncentiveRepository) Create(driverIncentive *models.DriverIncentive) error {
	return r.DB.Create(driverIncentive).Error
}

func (r *DriverIncentiveRepository) Update(id uint, updatedData *models.DriverIncentive) error {
	return r.DB.Model(&models.DriverIncentive{}).Where("id = ?", id).Updates(updatedData).Error
}

func (r *DriverIncentiveRepository) Delete(id uint) error {
	return r.DB.Delete(&models.DriverIncentive{}, id).Error
}
