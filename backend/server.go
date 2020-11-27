package main

import (
	"youtube-managed-app/backend/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	//Routes
	routes.Init(e)
	//Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
