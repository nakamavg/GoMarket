package order

import (
	"sync"
	"time"
)

type OrderType string

const (
	Buy  OrderType = "compra"
	Sell OrderType = "venta"
)

type Order struct {
	ID         int
	UserID     int
	Type       OrderType
	Symbol     string
	Quantity   int
	LimitPrice float64
	CreatedAt  time.Time
}

type OrderBook struct {
	BuyOrders  []*Order
	SellOrders []*Order
	mu         sync.RWMutex
}

func (r *MemoryRepository) matchOrders(symbol string) {
	book := r.orderBooks[symbol]
	if book == nil || len(book.BuyOrders) == 0 || len(book.SellOrders) == 0 {
		return
	}

	for len(book.BuyOrders) > 0 && len(book.SellOrders) > 0 {
		bestBuy := book.BuyOrders[0]
		bestSell := book.SellOrders[0]

		if bestBuy.LimitPrice < bestSell.LimitPrice {
			break
		}

		quantity := bestBuy.Quantity
		if bestSell.Quantity < quantity {
			quantity = bestSell.Quantity
		}

		bestBuy.Quantity -= quantity
		bestSell.Quantity -= quantity

		if bestBuy.Quantity == 0 {
			book.BuyOrders = book.BuyOrders[1:]
		}
		if bestSell.Quantity == 0 {
			book.SellOrders = book.SellOrders[1:]
		}

		r.stockRepo.UpdatePrice(symbol, bestSell.LimitPrice)
	}
}
