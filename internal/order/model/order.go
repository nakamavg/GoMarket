package model

import "time"

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
}