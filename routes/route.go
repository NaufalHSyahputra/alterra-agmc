package routes

import (
	"github.com/NaufalHSyahputra/alterra-agmc/constants"
	"github.com/NaufalHSyahputra/alterra-agmc/controllers"
	"github.com/NaufalHSyahputra/alterra-agmc/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)

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
	usersRoute.GET("", controllers.GetUsersController, middleware.JWT([]byte(constants.SECRET_JWT)))
	usersRoute.POST("/login", controllers.LoginUsersController)
	usersRoute.GET("/:id", controllers.GetUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	usersRoute.POST("", controllers.CreateUserController)
	usersRoute.PUT("/:id", controllers.UpdateUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	usersRoute.DELETE("/:id", controllers.DeleteUserController, middleware.JWT([]byte(constants.SECRET_JWT)))

	booksRoute := e.Group("/v1/books")
	booksRoute.GET("", controllers.GetBooksController)
	booksRoute.GET("/:id", controllers.GetBookController)
	booksRoute.POST("", controllers.CreateBookController, middleware.JWT([]byte(constants.SECRET_JWT)))
	booksRoute.PUT("/:id", controllers.UpdateBookController, middleware.JWT([]byte(constants.SECRET_JWT)))
	booksRoute.DELETE("/:id", controllers.DeleteBookController, middleware.JWT([]byte(constants.SECRET_JWT)))

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserDetailControllers)

	return e
}
