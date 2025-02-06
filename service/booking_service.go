package service

import (
	"testBeGo/models"
	"testBeGo/repository"
)

type BookingService interface {
	CreateBooking(booking models.Booking) (models.Booking, error)
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id uint) (models.Booking, error)
	UpdateBooking(id uint, booking models.Booking) (models.Booking, error)
	DeleteBooking(id uint) error
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(booking models.Booking) (models.Booking, error) {
	return s.repo.CreateBooking(booking)
}

func (s *bookingService) GetAllBookings() ([]models.Booking, error) {
	return s.repo.GetAllBookings()
}

func (s *bookingService) GetBookingByID(id uint) (models.Booking, error) {
	return s.repo.GetBookingByID(id)
}

func (s *bookingService) UpdateBooking(id uint, booking models.Booking) (models.Booking, error) {
	return s.repo.UpdateBooking(id, booking)
}

func (s *bookingService) DeleteBooking(id uint) error {
	return s.repo.DeleteBooking(id)
}
