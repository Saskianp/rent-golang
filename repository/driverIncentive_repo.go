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
	err := r.DB.
		Preload("Booking").
		Preload("Booking.Customer").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Car").
		Preload("Booking.BookingType").
		Preload("Booking.Driver").
		Find(&driverIncentive).Error
	return driverIncentive, err
}

func (r *DriverIncentiveRepository) GetByID(id uint) (models.DriverIncentive, error) {
	var driverIncentive models.DriverIncentive
	err := r.DB.
		Preload("Booking").
		Preload("Booking.Customer").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Car").
		Preload("Booking.BookingType").
		Preload("Booking.Driver").
		First(&driverIncentive, id).Error
	return driverIncentive, err
}


func (r *DriverIncentiveRepository) Delete(id uint) error {
	return r.DB.Delete(&models.DriverIncentive{}, id).Error
}
