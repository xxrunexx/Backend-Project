package route

import (
	// Import Framework Echo
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// Import Files
	c_account "movie-api/controller/account"
)

func Router() *echo.Echo {
	// Initiate Echo & Middleware
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// Initiate Route
	e.GET("/account", c_account.GetAllAccount)
	// e.POST("/account", AddAccount)
	// e.PUT("/account/:id", UpdateAccount)

	return e
}
