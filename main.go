package main

import (
    "github.com/nakama/Market/internal/api"
    "github.com/nakama/Market/internal/order/repository"
    "github.com/nakama/Market/internal/order/engine"
    "github.com/nakama/Market/internal/order/service"
    "github.com/nakama/Market/internal/stock"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Inicializar dependencias
    stockRepo := stock.NewMemoryRepository()
    orderRepo := repository.NewMemoryRepository()
    matchingEngine := engine.NewMatchingEngine(stockRepo)
    orderService := service.NewOrderService(orderRepo, matchingEngine)

    // Handlers
    stockHandler := api.NewStockHandler(stockRepo)
    orderHandler := api.NewOrderHandler(orderService)

    // Configurar rutas
    r.POST("/acciones", stockHandler.CreateStock)
    r.GET("/acciones/:simbolo", stockHandler.GetStockPrice)
    r.POST("/ordenes", orderHandler.CreateOrder)
    r.GET("/ordenes", orderHandler.GetOrderBook)

    r.Run("0.0.0.0:8080")
}