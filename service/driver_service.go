package service

import (
	"errors"
	"testBeGo/models"
	"testBeGo/repository"
)

type DriverService struct {
	Repo *repository.DriverRepository
}

func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{Repo: repo}
}

func (s *DriverService) GetAllDrivers() ([]models.Driver, error) {
	return s.Repo.GetAll()
}

func (s *DriverService) GetDriverByID(id uint) (models.Driver, error) {
	return s.Repo.GetByID(id)
}

func (s *DriverService) CreateDriver(driver *models.Driver) error {
	return s.Repo.Create(driver)
}

func (s *DriverService) UpdateDriver(id uint, updatedData *models.Driver) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("driver not found")
	}

	existing.Name = updatedData.Name
	existing.NIK = updatedData.NIK
	existing.PhoneNumber = updatedData.PhoneNumber
	existing.DailyRent = updatedData.DailyRent


	return s.Repo.Update(id, &existing)
}

func (s *DriverService) DeleteDriver(id uint) error {
	return s.Repo.Delete(id)
}
