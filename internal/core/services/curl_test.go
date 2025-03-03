package services

import (
	"testing"

	"curlie/internal/core/domain"
)

func TestGenerateCurlCommand(t *testing.T) {
	tests := []struct {
		name     string
		request  *domain.CurlRequest
		expected string
	}{
		{
			name: "GET request with no params",
			request: &domain.CurlRequest{
				Method: "GET",
				URL:    "https://api.example.com",
			},
			expected: "curl 'https://api.example.com'",
		},
		{
			name: "POST request with headers and body",
			request: &domain.CurlRequest{
				Method: "POST",
				URL:    "https://api.example.com/data",
				Headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer token123",
				},
				Body: `{"key": "value"}`,
			},
			expected: "curl -X POST -H Content-Type: application/json -H Authorization: Bearer token123 -d '{\"key\": \"value\"}' 'https://api.example.com/data'",
		},
		{
			name: "GET request with query params",
			request: &domain.CurlRequest{
				Method: "GET",
				URL:    "https://api.example.com",
				QueryParams: map[string]string{
					"param1": "value1",
					"param2": "value2",
				},
			},
			expected: "curl 'https://api.example.com?param1=value1&param2=value2'",
		},
	}

	service := NewCurlService(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GenerateCurlCommand(tt.request)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
