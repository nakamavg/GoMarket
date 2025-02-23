package repository

import (
    "sync"
    "errors"
    "market/internal/order/model"
)

type OrderRepository interface {
    Create(order *model.Order) error
    GetOrderBook(symbol string) (*model.OrderBook, error)
}

type MemoryRepository struct {
    orderBooks map[string]*model.OrderBook
    mu         sync.RWMutex
}

func NewMemoryRepository() OrderRepository {
    return &MemoryRepository{
        orderBooks: make(map[string]*model.OrderBook),
    }
}

func (r *MemoryRepository) Create(order *model.Order) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    symbol := order.Symbol
    if _, exists := r.orderBooks[symbol]; !exists {
        r.orderBooks[symbol] = &model.OrderBook{
            BuyOrders:  []*model.Order{},
            SellOrders: []*model.Order{},
        }
    }

    book := r.orderBooks[symbol]
    if order.Type == model.Buy {
        book.BuyOrders = append(book.BuyOrders, order)
    } else {
        book.SellOrders = append(book.SellOrders, order)
    }

    return nil
}

func (r *MemoryRepository) GetOrderBook(symbol string) (*model.OrderBook, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    book, exists := r.orderBooks[symbol]
    if !exists {
        return nil, errors.New("símbolo no encontrado")
    }
    return book, nil
}