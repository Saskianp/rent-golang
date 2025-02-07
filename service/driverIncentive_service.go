package service

import (
	"errors"
	"testBeGo/models"
	"testBeGo/repository"
)

type DriverIncentiveService struct {
	Repo *repository.DriverIncentiveRepository
}

func NewDriverIncentiveService(repo *repository.DriverIncentiveRepository) *DriverIncentiveService {
	return &DriverIncentiveService{Repo: repo}
}

func (s *DriverIncentiveService) GetAllDriverIncentives() ([]models.DriverIncentive, error) {
	return s.Repo.GetAll()
}

func (s *DriverIncentiveService) GetDriverIncentiveByID(id uint) (models.DriverIncentive, error) {
	return s.Repo.GetByID(id)
}

func (s *DriverIncentiveService) CreateDriverIncentive(driverIncentive *models.DriverIncentive) error {
	return s.Repo.Create(driverIncentive)
}

func (s *DriverIncentiveService) UpdateDriverIncentive(id uint, updatedData *models.DriverIncentive) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("driverIncentive not found")
	}

	// existing.Incentive = updatedData.Incentive

	return s.Repo.Update(id, &existing)
}

func (s *DriverIncentiveService) DeleteDriverIncentive(id uint) error {
	return s.Repo.Delete(id)
}
