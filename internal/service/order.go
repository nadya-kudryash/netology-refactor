package service

import (
	"fmt"

	"example/solid/internal/notifications"
	"example/solid/internal/repository"
)

type OrderService struct {
	repo     repository.RepositoryWriter
	notifier notifications.Notifier
}

func NewOrderService(repo repository.RepositoryWriter, notifier notifications.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	productsStr := fmt.Sprintf("%v", products)

	err := s.repo.SaveOrder(customer, productsStr, total, "pending")
	if err != nil {
		return err
	}

	err = s.notifier.Send(customer)
	if err != nil {
		return err
	}

	return nil
}
