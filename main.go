package main

import (
	"github.com/cassiusbessa/create-service/handlers"
	"github.com/cassiusbessa/create-service/repositories"
	"github.com/sirupsen/logrus"
)

func main() {
	r := handlers.Router()
	repositories.Repo.Ping()
	r.POST("/services/:company", handlers.CreateService)
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
