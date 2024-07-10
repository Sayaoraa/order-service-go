package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-tracking/internal/config"
	"order-tracking/internal/models"
)

// CreateOrder godoc
// @Summary Создание нового заказа
// @Description Создает новый заказ с указанными данными
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Данные заказа"
// @Success 201 {object} models.Order
// @Router /orders [post]
func CreateOrder(c echo.Context) error {
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return err
	}

	config.DB.Create(&order)
	return c.JSON(http.StatusCreated, order)
}

// GetOrders godoc
// @Summary Получение всех заказов
// @Description Возвращает список всех заказов
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func GetOrders(c echo.Context) error {
	var orders []models.Order
	config.DB.Find(&orders)
	return c.JSON(http.StatusOK, orders)
}

// UpdateOrder godoc
// @Summary Обновление заказа
// @Description Обновляет данные заказа по его ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "ID заказа"
// @Param order body models.Order true "Обновленные данные заказа"
// @Success 200 {object} models.Order
// @Router /orders/{id} [put]
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
