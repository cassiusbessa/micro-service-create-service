package main

import (
	"log"
	"net/http"

	entity "github.com/cassiusbessa/create-service/entity"
	errors "github.com/cassiusbessa/create-service/errors"
	repository "github.com/cassiusbessa/create-service/repository"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message interface{} `json:"message"`
}

var _, cancel = repository.MongoConnection()

func CreateService(c *gin.Context) {
	var service entity.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := service.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := c.Param("company")
	if err := repository.CreateService(db, service); err != nil {
		repository.SaveError(db, *errors.NewError(http.StatusInternalServerError, "Error Mongo creating service", "CreateService", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Service created successfully"})

	defer cancel()
}

func main() {
	r := gin.Default()
	r.POST("/service/:company", CreateService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
