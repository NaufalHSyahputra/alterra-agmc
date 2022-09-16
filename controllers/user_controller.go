package controllers

import (
	"net/http"
	"strconv"

	"github.com/NaufalHSyahputra/alterra-agmc/lib/database"
	"github.com/NaufalHSyahputra/alterra-agmc/models"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	users, e := database.GetUser(int(id))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user_id, e := database.CreateUser(user.Name, user.Email, user.Password)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"user_id": user_id,
	})
}

func UpdateUserController(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	user_id, e := database.UpdateUser(int(id), user.Name, user.Email, user.Password)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"user_id": user_id,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	e := database.DeleteUser(int(id))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
