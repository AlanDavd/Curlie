package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"curlie/internal/core/services"
	"curlie/internal/core/domain"
)

type curlHandler struct {
	service services.CurlService
}

func NewCurlHandler(service services.CurlService) *curlHandler {
	return &curlHandler{
		service: service,
	}
}

func (h *curlHandler) GenerateCurl(c *gin.Context) {
	var request domain.CurlRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if request.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	// Set default method if not provided
	if request.Method == "" {
		request.Method = "GET"
	}

	command, err := h.service.GenerateCurlCommand(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"command": command,
	})
}
