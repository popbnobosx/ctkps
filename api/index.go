package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Hello(c echo.Context) error {
	return c.Redirect(302, "https://xn--68j5e377y.com/")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//go Rule()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/hello", Hello)

	e.ServeHTTP(w, r)
}
