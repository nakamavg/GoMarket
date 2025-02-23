package stock

import (
    "errors"
    "sync"
)

type Stock struct {
    Symbol      string
    CurrentPrice float64
}

type Repository interface {
    Create(stock *Stock) error
    Get(symbol string) (*Stock, error)
    UpdatePrice(symbol string, price float64) error
}

type MemoryRepository struct {
    stocks map[string]*Stock
    mu     sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
    return &MemoryRepository{
        stocks: make(map[string]*Stock),
    }
}

func (r *MemoryRepository) Create(stock *Stock) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.stocks[stock.Symbol]; exists {
        return errors.New("stock already exists")
    }
    
    r.stocks[stock.Symbol] = stock
    return nil
}

func (r *MemoryRepository) Get(symbol string) (*Stock, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    stock, exists := r.stocks[symbol]
    if !exists {
        return nil, errors.New("stock not found")
    }
    return stock, nil
}

func (r *MemoryRepository) UpdatePrice(symbol string, price float64) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    stock, exists := r.stocks[symbol]
    if !exists {
        return errors.New("stock not found")
    }
    
    stock.CurrentPrice = price
    return nil
}