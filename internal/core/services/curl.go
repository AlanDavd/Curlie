package services

import (
	"fmt"
	"strings"

	"curlie/internal/core/ports"
	"curlie/internal/core/domain"
)

// CurlService defines the interface for curl command generation.
type CurlService interface {
	GenerateCurlCommand(req *domain.CurlRequest) (string, error)
}

// curlService implements CurlService interface.
type curlService struct {
	repo ports.CurlRepository
}

func NewCurlService(repo ports.CurlRepository) *curlService {
	return &curlService{
		repo: repo,
	}
}

func (c *curlService) GenerateCurlCommand(req *domain.CurlRequest) (string, error) {
	var parts []string
	parts = append(parts, "curl")

	// Add method
	if req.Method != "" && req.Method != "GET" {
		parts = append(parts, "-X", req.Method)
	}
	// Add headers
	for key, value := range req.Headers {
		parts = append(parts, "-H", fmt.Sprintf("%s: %s", key, value))
	}
	// Add body if present
	if req.Body != "" {
		parts = append(parts, "-d", fmt.Sprintf("'%s'", req.Body))
	}

	// Build URL with query parameters
	url := req.URL
	if len(req.QueryParams) > 0 {
		queryParts := make([]string, 0)
		for key, value := range req.QueryParams {
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
		}
		url = fmt.Sprintf("%s?%s", url, strings.Join(queryParts, "&"))
	}
	parts = append(parts, fmt.Sprintf("'%s'", url))

	return strings.Join(parts, " "), nil
}
