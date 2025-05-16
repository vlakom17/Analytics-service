package main

import (
	"context"
	"log"
	nethttp "net/http"

	"github.com/vlakom17/analytics-service/internal/config"
	"github.com/vlakom17/analytics-service/internal/infrastructure/db"
	"github.com/vlakom17/analytics-service/internal/kafka"
	"github.com/vlakom17/analytics-service/internal/repository/postgres"
	"github.com/vlakom17/analytics-service/internal/service"
	"github.com/vlakom17/analytics-service/internal/transport/http"
)

func main() {

	cfg := config.LoadConfig()

	dbConn, err := db.NewPostgresConnection(&cfg.DB)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbConn.Close()

	FactRepo := postgres.NewFactRepo(dbConn)
	FactService := service.NewFactService(FactRepo)

	consumer := kafka.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, FactService)
	go consumer.Start(context.Background())

	router := http.NewRouter(dbConn, cfg.Admin.Secret)

	log.Println("Сервер запущен на порту 8080")
	if err := nethttp.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
