package service

import (
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

func (s *DriverIncentiveService) DeleteDriverIncentive(id uint) error {
	return s.Repo.Delete(id)
}
