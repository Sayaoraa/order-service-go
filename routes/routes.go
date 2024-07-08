package routes

import (
	"github.com/labstack/echo/v4"
	"order-tracking/handlers"
	"order-tracking/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register)

	r := e.Group("/orders")
	r.Use(middleware.JWT())
	r.POST("", handlers.CreateOrder)
	r.GET("", handlers.GetOrders)
	r.PUT("/:id", handlers.UpdateOrder)
}
