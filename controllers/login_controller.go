package controllers

import (
	"fmt"
	"net/http"

	"github.com/NaufalHSyahputra/alterra-agmc/lib/database"
	"github.com/NaufalHSyahputra/alterra-agmc/models"
	"github.com/labstack/echo/v4"
)

func LoginUsersController(c echo.Context) error {
	fmt.Println("Masuk sini?")
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
