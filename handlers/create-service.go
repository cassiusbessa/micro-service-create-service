package handlers

import (
	"net/http"

	"github.com/cassiusbessa/create-service/entities"
	"github.com/cassiusbessa/create-service/logs"
	"github.com/cassiusbessa/create-service/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateService(c *gin.Context) {
	var service entities.Service
	db := c.Param("company")
	logrus.Warnf("Creating Service on %v", db)
	if err := c.BindJSON(&service); err != nil {
		logrus.Errorf("Error decoding Service %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := service.Validate(); err != nil {
		logrus.Errorf("Error validating Service on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := repositories.CreateService(db, service); err != nil {
		logrus.Errorf("Error Inserting Service to MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer logs.Elapsed("Create Service")()
	c.JSON(http.StatusCreated, gin.H{"message": "Service created successfully"})
}
