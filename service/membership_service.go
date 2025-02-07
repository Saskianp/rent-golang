package service

import (
	"errors"
	"testBeGo/models"
	"testBeGo/repository"
)

type MembershipService struct {
	Repo *repository.MembershipRepository
}

func NewMembershipService(repo *repository.MembershipRepository) *MembershipService {
	return &MembershipService{Repo: repo}
}

func (s *MembershipService) GetAllMemberships() ([]models.Membership, error) {
	return s.Repo.GetAll()
}

func (s *MembershipService) GetMembershipByID(id uint) (models.Membership, error) {
	return s.Repo.GetByID(id)
}

func (s *MembershipService) CreateMembership(membership *models.Membership) error {
	return s.Repo.Create(membership)
}

func (s *MembershipService) UpdateMembership(id uint, updatedData *models.Membership) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("membership not found")
	}

	existing.MembershipName = updatedData.MembershipName
	existing.Discount = updatedData.Discount

	return s.Repo.Update(id, &existing)
}

func (s *MembershipService) DeleteMembership(id uint) error {
	return s.Repo.Delete(id)
}
