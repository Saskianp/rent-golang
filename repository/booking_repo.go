package repository

import (
	"testBeGo/models"
	"errors"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking models.Booking) (models.Booking, error)
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id uint) (models.Booking, error)
	UpdateBooking(id uint, booking models.Booking) (models.Booking, error)
	DeleteBooking(id uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking models.Booking) (models.Booking, error) {
    // Menarik data mobil berdasarkan CarID untuk mendapatkan DailyRent
    var car models.Car
    if err := r.db.First(&car, booking.CarID).Error; err != nil {
        return models.Booking{}, err
    }

    // Menghitung durasi sewa dalam hari
    duration := booking.EndRent.Sub(booking.StartRent).Hours() / 24 // dalam hari
    if duration < 1 {
        duration = 1 // Minimal 1 hari
    }

    // Menghitung total biaya sewa berdasarkan durasi dan harga sewa per hari
    booking.TotalCost = duration * float64(car.DailyRent)

    // Menyimpan booking ke dalam database
    err := r.db.Create(&booking).Error
    return booking, err
}

func (r *bookingRepository) GetAllBookings() ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.db.Preload("Customer").Preload("Car").Find(&bookings).Error
    return bookings, err
}

func (r *bookingRepository) GetBookingByID(id uint) (models.Booking, error) {
	var booking models.Booking
	err := r.db.Preload("Customer").Preload("Car").First(&booking, id).Error

	// Cek jika data tidak ditemukan
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Booking{}, errors.New("booking not found")
	}

	return booking, err
}

func (r *bookingRepository) UpdateBooking(id uint, booking models.Booking) (models.Booking, error) {
	var existingBooking models.Booking

	// Cek apakah booking dengan ID tersebut ada
	err := r.db.First(&existingBooking, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Booking{}, errors.New("booking not found")
	}

	// Update booking
	err = r.db.Model(&existingBooking).Updates(booking).Error
	return existingBooking, err
}

func (r *bookingRepository) DeleteBooking(id uint) error {
	var booking models.Booking

	// Cek apakah booking dengan ID tersebut ada
	err := r.db.First(&booking, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("booking not found")
	}

	// Hapus booking
	return r.db.Delete(&booking).Error
}
