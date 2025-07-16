package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/DanyaSokolov/subscription-service/docs" 

	"github.com/DanyaSokolov/subscription-service/internal/db/logger"
	"github.com/DanyaSokolov/subscription-service/internal/handler"
	"github.com/DanyaSokolov/subscription-service/internal/repository"
)

// @title           Subscription Service API
// @version         1.0
// @description     API для управления подписками
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@example.com

// @host      localhost:8080
// @BasePath  /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Логгер
	l := logger.NewLogger()

	// Репозиторий
	repo := repository.NewSubscriptionRepository(db)

	// Хендлер
	h := handler.NewSubscriptionHandler(repo, l)

	// Роутер
	r := gin.Default()

	// CRUD endpoints
	r.POST("/subscriptions", h.Create)
	r.GET("/subscriptions/:id", h.GetByID)
	r.PUT("/subscriptions/:id", h.Update)
	r.DELETE("/subscriptions/:id", h.Delete)
	r.GET("/subscriptions", h.List)

	// Общая сумма
	r.GET("/subscriptions/total-cost", h.TotalCost)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
