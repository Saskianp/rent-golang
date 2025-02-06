package service

import (
	"errors"
	"testBeGo/models"
	"testBeGo/repository"
)

type CustomerService struct {
	Repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{Repo: repo}
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.Repo.GetAll()
}

func (s *CustomerService) GetCustomerByID(id uint) (models.Customer, error) {
	return s.Repo.GetByID(id)
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
	return s.Repo.Create(customer)
}

func (s *CustomerService) UpdateCustomer(id uint, updatedData *models.Customer) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("customer not found")
	}

	existing.Name = updatedData.Name
	existing.NIK = updatedData.NIK
	existing.PhoneNumber = updatedData.PhoneNumber

	return s.Repo.Update(id, &existing)
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	return s.Repo.Delete(id)
}
