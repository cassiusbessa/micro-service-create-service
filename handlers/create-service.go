package handlers

import (
	"net/http"

	"github.com/cassiusbessa/create-service/entities"
	"github.com/cassiusbessa/create-service/repositories"
	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context) {
	var service entities.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := service.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	db := c.Param("company")
	if err := repositories.CreateService(db, service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Service created successfully"})
}
