package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (h *handler) GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenu(menuType)
	if err != nil {
		fmt.Print("got error %\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}