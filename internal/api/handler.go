package api

import (
    "net/http"
    "time"
    "market/internal/order/model"
    "market/internal/order/service"
    "market/internal/stock"
    "github.com/gin-gonic/gin"
)

type StockHandler struct {
    repo stock.Repository
}

func NewStockHandler(repo stock.Repository) *StockHandler {
    return &StockHandler{repo: repo}
}

func (h *StockHandler) CreateStock(c *gin.Context) {
    var request struct {
        Symbol       string  `json:"simbolo"`
        InitialPrice float64 `json:"precio_inicial"`
    }
    
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    
    newStock := &stock.Stock{
        Symbol:      request.Symbol,
        CurrentPrice: request.InitialPrice,
    }
    
    if err := h.repo.Create(newStock); err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, newStock)
}

func (h *StockHandler) GetStockPrice(c *gin.Context) {
    symbol := c.Param("simbolo")
    stock, err := h.repo.Get(symbol)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, stock)
}

type OrderHandler struct {
    service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
    return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var request struct {
        UserID     int     `json:"usuario_id"`
        Type       string  `json:"tipo"`
        Symbol     string  `json:"simbolo"`
        Quantity   int     `json:"cantidad"`
        LimitPrice float64 `json:"precio_limite"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    orderType := model.OrderType(request.Type)
    if orderType != model.Buy && orderType != model.Sell {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order type"})
        return
    }

    newOrder := &model.Order{
        UserID:     request.UserID,
        Type:       orderType,
        Symbol:     request.Symbol,
        Quantity:   request.Quantity,
        LimitPrice: request.LimitPrice,
        CreatedAt:  time.Now(),
    }

    if err := h.service.PlaceOrder(newOrder); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newOrder)
}

func (h *OrderHandler) GetOrderBook(c *gin.Context) {
    symbol := c.Query("simbolo")
    if symbol == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "symbol parameter required"})
        return
    }

    book, err := h.service.GetOrderBook(symbol)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    response := gin.H{
        "ordenes_compra":  formatOrders(book.BuyOrders),
        "ordenes_venta": formatOrders(book.SellOrders),
    }
    
    c.JSON(http.StatusOK, response)
}

func formatOrders(orders []*model.Order) []gin.H {
    var result []gin.H
    for _, o := range orders {
        result = append(result, gin.H{
            "usuario_id":    o.UserID,
            "cantidad":      o.Quantity,
            "precio_limite": o.LimitPrice,
            "creado_en":     o.CreatedAt.Format(time.RFC3339),
        })
    }
    return result
}