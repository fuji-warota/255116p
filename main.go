package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, ping())
	})
  e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, pong())
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func ping() string {
	return "pong"
}

func pong() string {
	return "ping"
}
