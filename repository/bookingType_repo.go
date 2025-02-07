package repository

import (
	"testBeGo/models"
	"gorm.io/gorm"
)

type BookingTypeRepository struct {
	DB *gorm.DB
}

func NewBookingTypeRepository(db *gorm.DB) *BookingTypeRepository {
	return &BookingTypeRepository{DB: db}
}

func (r *BookingTypeRepository) GetAll() ([]models.BookingType, error) {
	var bookingType []models.BookingType
	err := r.DB.Find(&bookingType).Error
	return bookingType, err
}

func (r *BookingTypeRepository) GetByID(id uint) (models.BookingType, error) {
	var bookingType models.BookingType
	err := r.DB.First(&bookingType, id).Error
	return bookingType, err
}

func (r *BookingTypeRepository) Create(bookingType *models.BookingType) error {
	return r.DB.Create(bookingType).Error
}

func (r *BookingTypeRepository) Update(id uint, updatedData *models.BookingType) error {
	return r.DB.Model(&models.BookingType{}).Where("id = ?", id).Updates(updatedData).Error
}

func (r *BookingTypeRepository) Delete(id uint) error {
	return r.DB.Delete(&models.BookingType{}, id).Error
}
