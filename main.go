package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	personhttp "person-crud/internal/http"
	"person-crud/internal/http/middleware"
	"person-crud/internal/logic"
	"person-crud/internal/postgres"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using env variables")
	}

	logger := logrus.New()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to db: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Fatal("failed to ping db: ", err)
	}

	if err := postgres.InitSchema(db); err != nil {
		logger.Fatal("failed to init schema: ", err)
	}

	repo := postgres.NewPersonRepository(db)
	logic := logic.NewPersonLogic(repo)
	handler := personhttp.NewPersonHandler(logic)

	e := echo.New()
	e.Use(middleware.RequestLoggerMiddleware)

	e.POST("/person", handler.CreatePerson)
	e.GET("/person", handler.GetPersons)
	e.GET("/person/:id", handler.GetPerson)
	e.PUT("/person/:id", handler.UpdatePerson)
	e.DELETE("/person/:id", handler.DeletePerson)

	logger.Info("Server running at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
