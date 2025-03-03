package handler

import (
	"net/http"

	"github.com/alandavd/curlie/internal/domain/curl"
	"github.com/gin-gonic/gin"
)

type CurlHandler struct {
	service curl.Service
}

func NewCurlHandler(service curl.Service) *CurlHandler {
	return &CurlHandler{
		service: service,
	}
}

func (h *CurlHandler) GenerateCurl(c *gin.Context) {
	var request curl.Request
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

	command, err := h.service.GenerateCurlCommand(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"command": command,
	})
} 