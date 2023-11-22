package service

import (
	"context"
	"test_task/internal/repository"
)

type BinanceService struct {
	repository *repository.BinanceRepository
}

func NewService(repository *repository.BinanceRepository) *BinanceService {
	return &BinanceService{
		repository: repository,
	}
}

func (service *BinanceService) GetPairs(ctx context.Context, count int) ([]string, error) {
	return service.repository.GetPairs(ctx, count)
}

func (service *BinanceService) GetPrice(ctx context.Context, pair string) (string, error) {
	return service.repository.GetPrice(ctx, pair)
}
