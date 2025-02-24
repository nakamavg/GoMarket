package service

import (
	"github.com/nakama/Market/internal/order/engine"
	"github.com/nakama/Market/internal/order/model"
	"github.com/nakama/Market/internal/order/repository"
)

type OrderService struct {
    repo     repository.OrderRepository
    engine   *engine.MatchingEngine
}

func NewOrderService(repo repository.OrderRepository, engine *engine.MatchingEngine) *OrderService {
    return &OrderService{
        repo: repo,
        engine: engine,
    }
}

func (s *OrderService) GetOrderBook(symbol string) (*model.OrderBook, error) {
    return s.repo.GetOrderBook(symbol)
}

func (s *OrderService) PlaceOrder(order *model.Order) error {
    if err := s.repo.Create(order); err != nil {
        return err
    }

    book, err := s.repo.GetOrderBook(order.Symbol)
    if err != nil {
        return err
    }

    s.engine.Match(book)
    return nil
}