package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"order-tracking/config"
	"order-tracking/models"
	"order-tracking/routes"
	_ "order-tracking/swagger" // Это важно для автогенерации документации
)

// @title Order Tracking API
// @version 1.0
// @description This is an API for tracking orders.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config.InitDB()
	models.MigrateDB(config.DB)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
