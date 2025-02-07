package service

import (
	"errors"
	"testBeGo/models"
	"testBeGo/repository"
)

type BookingTypeService struct {
	Repo *repository.BookingTypeRepository
}

func NewBookingTypeService(repo *repository.BookingTypeRepository) *BookingTypeService {
	return &BookingTypeService{Repo: repo}
}

func (s *BookingTypeService) GetAllBookingTypes() ([]models.BookingType, error) {
	return s.Repo.GetAll()
}

func (s *BookingTypeService) GetBookingTypeByID(id uint) (models.BookingType, error) {
	return s.Repo.GetByID(id)
}

func (s *BookingTypeService) CreateBookingType(bookingType *models.BookingType) error {
	return s.Repo.Create(bookingType)
}

func (s *BookingTypeService) UpdateBookingType(id uint, updatedData *models.BookingType) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("booking type not found")
	}

	existing.BookingType = updatedData.BookingType
	existing.Description = updatedData.Description

	return s.Repo.Update(id, &existing)
}

func (s *BookingTypeService) DeleteBookingType(id uint) error {
	return s.Repo.Delete(id)
}
