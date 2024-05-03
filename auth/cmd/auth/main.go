package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status string
			Path   string
			Info   string
		}{Status: "OK", Path: "/", Info: "Version 2"})
	})

	e.GET("/readiness", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status string
			Path   string
		}{Status: "OK", Path: "/readiness"})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Status string
			Path   string
		}{Status: "OK", Path: "/health"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
