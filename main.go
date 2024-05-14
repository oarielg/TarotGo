package main

import (
	"github.com/oarielg/TarotGo/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/images", "images")

	e.GET("/", controller.IndexHandler)
	e.GET("/read", controller.ReadHandler)

	e.Logger.Fatal(e.Start(":8090"))
}
