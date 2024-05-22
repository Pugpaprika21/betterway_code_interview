package main

import (
	"log"

	"github.com/Pugpaprika21/question10/db"
	"github.com/Pugpaprika21/question10/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	dbc, err := db.New().UseDB()
	if err != nil {
		panic(err)
	}

	migrate := db.NewMigration(dbc)
	migrate.Run()

	router.EchoSetupRouter(e, dbc)

	e.Logger.Fatal(e.Start(":8081"))
}
