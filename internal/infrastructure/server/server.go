package server

import (
	"net/http"

	"github.com/alandavd/curlie/internal/application/service"
	"github.com/alandavd/curlie/internal/infrastructure/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.Default()

	// Initialize dependencies
	curlService := service.NewCurlService(nil) // No repository for now
	curlHandler := handler.NewCurlHandler(curlService)

	// Load HTML templates
	engine.LoadHTMLGlob("internal/infrastructure/ui/templates/*")

	// Serve static files
	engine.Static("/static", "internal/infrastructure/ui/static")

	// Routes
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	engine.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.html", nil)
	})

	engine.GET("/terms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.html", nil)
	})

	api := engine.Group("/api")
	{
		api.POST("/curl", curlHandler.GenerateCurl)
	}

	return &Server{
		engine: engine,
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
} 