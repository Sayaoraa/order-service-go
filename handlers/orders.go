package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-tracking/config"
	"order-tracking/models"
)

func CreateOrder(c echo.Context) error {
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return err
	}

	config.DB.Create(&order)
	return c.JSON(http.StatusCreated, order)
}

func GetOrders(c echo.Context) error {
	var orders []models.Order
	config.DB.Find(&orders)
	return c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	order := new(models.Order)
	if err := config.DB.First(&order, id).Error; err != nil {
		return echo.ErrNotFound
	}

	if err := c.Bind(order); err != nil {
		return err
	}

	config.DB.Save(&order)
	return c.JSON(http.StatusOK, order)
}
