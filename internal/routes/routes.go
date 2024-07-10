package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	handlers2 "order-tracking/internal/handlers"
	"order-tracking/internal/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/login", handlers2.Login)
	e.POST("/register", handlers2.Register)

	orders := e.Group("/orders")
	orders.Use(middleware.JWT())
	orders.POST("", handlers2.CreateOrder)
	orders.GET("", handlers2.GetOrders)
	orders.PUT("/:id", handlers2.UpdateOrder)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
