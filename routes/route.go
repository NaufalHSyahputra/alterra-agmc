package routes

import (
	"github.com/NaufalHSyahputra/alterra-agmc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// v := e.Group("/v1/")

	// e.Any("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "Hello World!",
	// 	})
	// })

	// v.Any("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "Hello World!",
	// 	})
	// })

	usersRoute := e.Group("/v1/users")
	usersRoute.GET("", controllers.GetUsersController)
	usersRoute.GET("/:id", controllers.GetUserController)
	usersRoute.POST("", controllers.CreateUserController)
	usersRoute.PUT("/:id", controllers.UpdateUserController)
	usersRoute.DELETE("/:id", controllers.DeleteUserController)

	booksRoute := e.Group("/v1/books")
	booksRoute.GET("", controllers.GetBooksController)
	booksRoute.GET("/:id", controllers.GetBookController)
	booksRoute.POST("", controllers.CreateBookController)
	booksRoute.PUT("/:id", controllers.UpdateBookController)
	booksRoute.DELETE("/:id", controllers.DeleteBookController)

	return e
}
