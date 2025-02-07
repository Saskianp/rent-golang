package repository

import (
	"testBeGo/models"
	"gorm.io/gorm"
)

type MembershipRepository struct {
	DB *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) *MembershipRepository {
	return &MembershipRepository{DB: db}
}

func (r *MembershipRepository) GetAll() ([]models.Membership, error) {
	var membership []models.Membership
	err := r.DB.Find(&membership).Error
	return membership, err
}

func (r *MembershipRepository) GetByID(id uint) (models.Membership, error) {
	var membership models.Membership
	err := r.DB.First(&membership, id).Error
	return membership, err
}

func (r *MembershipRepository) Create(membership *models.Membership) error {
	return r.DB.Create(membership).Error
}

func (r *MembershipRepository) Update(id uint, updatedData *models.Membership) error {
	return r.DB.Model(&models.Membership{}).Where("id = ?", id).Updates(updatedData).Error
}

func (r *MembershipRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Membership{}, id).Error
}
