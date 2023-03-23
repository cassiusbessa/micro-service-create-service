package main

import (
	"github.com/cassiusbessa/create-service/handlers"
	"github.com/cassiusbessa/create-service/logs"
	"github.com/cassiusbessa/create-service/repositories"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	r := handlers.Router()
	repositories.Repo.Ping()
	r.POST("/services/:company", handlers.CreateService)
	r.StaticFile("/logs", "./logs/logs.log")
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
