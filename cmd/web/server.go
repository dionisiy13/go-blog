package main

import (
	"github.com/dionisiy13/go-web/pkg/config"
	"github.com/dionisiy13/go-web/pkg/handlers"
	"github.com/labstack/echo/v4"
	sessionM "github.com/spazzymoto/echo-scs-session"
)

const PORT = ":8090"

func InitServer() {
	// init app
	config.GetAppConfig()

	// server
	e := echo.New()

	middlewares(e)

	// routes
	e.GET("/", handlers.Home)
	e.GET("/about", handlers.About)
	e.GET("/favicon.ico", handlers.DoNothing)

	// start server
	e.Logger.Fatal(e.Start(PORT))
}

func middlewares(e *echo.Echo) {
	// middlewares
	a := config.GetAppConfig()
	e.Use(sessionM.LoadAndSave(a.Session))
}
