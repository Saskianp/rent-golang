package repository

import (
	"testBeGo/models"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) GetByID(id uint) (models.Customer, error) {
	var customer models.Customer
	err := r.DB.First(&customer, id).Error
	return customer, err
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.DB.Create(customer).Error
}

func (r *CustomerRepository) Update(id uint, updatedData *models.Customer) error {
	return r.DB.Model(&models.Customer{}).Where("id = ?", id).Updates(updatedData).Error
}

func (r *CustomerRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Customer{}, id).Error
}
