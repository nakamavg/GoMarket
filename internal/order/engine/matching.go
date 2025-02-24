package engine

import (
    "math"
    "github.com/nakama/Market/internal/order/model"
    "github.com/nakama/Market/internal/stock"
)

type MatchingEngine struct {
    stockRepo stock.Repository
}

func NewMatchingEngine(stockRepo stock.Repository) *MatchingEngine {
    return &MatchingEngine{stockRepo: stockRepo}
}

func (e *MatchingEngine) Match(book *model.OrderBook) {
    for {
        if len(book.BuyOrders) == 0 || len(book.SellOrders) == 0 {
            break
        }

        bestBuy := book.BuyOrders[0]
        bestSell := book.SellOrders[0]

        if bestBuy.LimitPrice < bestSell.LimitPrice {
            break
        }

        quantity := math.Min(float64(bestBuy.Quantity), float64(bestSell.Quantity))
        transactionPrice := bestSell.LimitPrice

        bestBuy.Quantity -= int(quantity)
        bestSell.Quantity -= int(quantity)

        // Eliminar órdenes completadas
        if bestBuy.Quantity == 0 {
            book.BuyOrders = book.BuyOrders[1:]
        }
        if bestSell.Quantity == 0 {
            book.SellOrders = book.SellOrders[1:]
        }

        e.stockRepo.UpdatePrice(bestBuy.Symbol, transactionPrice)
    }
}