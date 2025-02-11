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
	var car models.Car
	var customer models.Customer
	var membership models.Membership
	var driver models.Driver

	// Ambil informasi mobil
	if err := r.db.First(&car, booking.CarID).Error; err != nil {
		return models.Booking{}, err
	}

	// Ambil informasi customer
	if err := r.db.First(&customer, booking.CustomerID).Error; err != nil {
		return models.Booking{}, err
	}

	// Ambil informasi membership customer (jika ada)
	if err := r.db.First(&membership, customer.MembershipID).Error; err != nil {
		// Jika tidak ada membership, asumsi diskon 0%
		membership.Discount = 0
	}

	// Hitung durasi sewa dalam hari
	duration := booking.EndRent.Sub(booking.StartRent).Hours() / 24 + 1
	if duration < 1 {
		duration = 1 // Minimal 1 hari
	}

	// Hitung diskon dari membership
	booking.Discount = (float64(car.DailyRent) * duration) * (membership.Discount / 100)

	// Hitung biaya driver jika ada
	if booking.DriverID != nil {
		if err := r.db.First(&driver, *booking.DriverID).Error; err != nil {
			return models.Booking{}, err
		}
		booking.DriverCost = float64(driver.DailyRent) * duration
	}

	// Hitung total biaya
	// booking.TotalCost = (float64(car.DailyRent) * duration) - booking.Discount + booking.DriverCost
	booking.TotalCost = float64(car.DailyRent) * duration

	// // Simpan booking ke database
	// err := r.db.Create(&booking).Error
	// return booking, err

		// Simpan booking ke database
	err := r.db.Create(&booking).Error
	if err != nil {
		return models.Booking{}, err
	}

	// Hitung insentif driver: (durasi * car.daily_rent) * 5%
	incentive := (duration * float64(car.DailyRent)) * 0.05

	// Simpan insentif ke tabel driver_incentives
	driverIncentive := models.DriverIncentive{
		BookingID: booking.ID,
		Incentive: incentive,
	}
	err = r.db.Create(&driverIncentive).Error
	if err != nil {
		return models.Booking{}, err
	}

	return booking, nil

}

func (r *bookingRepository) GetAllBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.
		Preload("Customer.Membership"). // Preload Membership untuk diskon
		Preload("Car").
		Preload("BookingType").
		Preload("Driver").
		Find(&bookings).Error

	return bookings, err
}

func (r *bookingRepository) GetBookingByID(id uint) (models.Booking, error) {
	var booking models.Booking
	err := r.db.
		Preload("Customer.Membership"). // Preload Membership untuk diskon
		Preload("Car").
		Preload("BookingType").
		Preload("Driver").
		First(&booking, id).Error // Gunakan `First` agar hanya ambil satu record

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
	if err != nil {
		return existingBooking, err
	}

	// Ambil data yang sudah diperbarui beserta relasi
	err = r.db.Preload("Customer").Preload("Customer.Membership").
		Preload("Car").
		Preload("BookingType").
		Preload("Driver").
		First(&existingBooking, id).Error

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
