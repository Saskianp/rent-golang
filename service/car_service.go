package service

import (
	"testBeGo/models"
	"testBeGo/repository"
)

type CarService struct {
	Repo *repository.CarRepository
}

func NewCarService(repo *repository.CarRepository) *CarService {
	return &CarService{Repo: repo}
}

func (s *CarService) GetAllCars() ([]models.Car, error) {
	return s.Repo.GetAllCars()
}

func (s *CarService) GetCarByID(id uint) (*models.Car, error) {
	return s.Repo.GetCarByID(id)
}

func (s *CarService) CreateCar(car *models.Car) error {
	return s.Repo.CreateCar(car)
}

func (s *CarService) UpdateCar(id uint, car *models.Car) error {
	return s.Repo.UpdateCar(id, car)
}

func (s *CarService) DeleteCar(id uint) error {
	return s.Repo.DeleteCar(id)
}
