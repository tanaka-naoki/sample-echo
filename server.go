package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"sample-echo/handler"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/transecode", handler.MainPage())

    e.Start(":1323")
}