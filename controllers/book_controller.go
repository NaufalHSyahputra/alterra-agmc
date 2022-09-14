package controllers

import (
	"net/http"

	"github.com/NaufalHSyahputra/alterra-agmc/models"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books := models.InitBookData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}
func GetBookController(c echo.Context) error {
	books := models.InitBookData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}
func CreateBookController(c echo.Context) error {
	books := models.InitBookData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}
func UpdateBookController(c echo.Context) error {
	books := models.InitBookData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}
func DeleteBookController(c echo.Context) error {
	books := models.InitBookData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}
