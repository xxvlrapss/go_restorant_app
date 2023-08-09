package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Print("got error %\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	orderData, err := h.restoUsecase.Order(request)
	if err != nil {
		fmt.Print("got error %\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": orderData,
		})

}

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderID := c.Param("orderID")

	orderData, err := h.restoUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		OrderID: orderID,
	})
	if err != nil {
		fmt.Println("got error %\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
	})
}
return c.JSON(http.StatusOK, map[string]interface{}{
	"data": orderData,
})
	}