package services

import (
	"context"

	"mfo-service/internal/domain"
	"mfo-service/internal/repositories"
)

type ColdUsersRepository interface {
	FindByPhone(ctx context.Context, phone int) (*domain.ColdUsers, error)
}

type ColdUsersService struct {
	repo *repositories.ColdUsersRepository
}

func NewColdUsersService(repo *repositories.ColdUsersRepository) *ColdUsersService {
	return &ColdUsersService{repo: repo}
}

func (s *ColdUsersService) GetColdUsers(ctx context.Context, phone int) (*domain.ColdUsers, error) {
	return s.repo.FindByPhone(ctx, phone)
}
