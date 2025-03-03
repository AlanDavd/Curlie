package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"curlie/internal/core/domain"
)

// MockCurlService is a mock implementation of the CurlService interface
type MockCurlService struct {
	mock.Mock
}

func (m *MockCurlService) GenerateCurlCommand(req *domain.CurlRequest) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func TestGenerateCurl(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		request        domain.CurlRequest
		mockResponse   string
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful curl command generation",
			request: domain.CurlRequest{
				Method: "GET",
				URL:    "https://api.example.com",
				Headers: map[string]string{
					"Authorization": "Bearer token123",
				},
			},
			mockResponse:   "curl -H 'Authorization: Bearer token123' 'https://api.example.com'",
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"command": "curl -H 'Authorization: Bearer token123' 'https://api.example.com'",
			},
		},
		{
			name: "missing URL",
			request: domain.CurlRequest{
				Method: "GET",
				URL:    "",
			},
			mockResponse:   "",
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "URL is required",
			},
		},
		{
			name: "complex request with all fields",
			request: domain.CurlRequest{
				Method: "POST",
				URL:    "https://api.example.com/data",
				Headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer token123",
				},
				Body: `{"key": "value"}`,
				QueryParams: map[string]string{
					"param1": "value1",
				},
			},
			mockResponse:   "curl -X POST -H 'Content-Type: application/json' -H 'Authorization: Bearer token123' -d '{\"key\": \"value\"}' 'https://api.example.com/data?param1=value1'",
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"command": "curl -X POST -H 'Content-Type: application/json' -H 'Authorization: Bearer token123' -d '{\"key\": \"value\"}' 'https://api.example.com/data?param1=value1'",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock service
			mockService := new(MockCurlService)
			if tt.request.URL != "" {
				mockService.On("GenerateCurlCommand", &tt.request).Return(tt.mockResponse, tt.mockError)
			}

			// Create a new handler with the mock service
			handler := NewCurlHandler(mockService)

			// Create a new Gin router
			router := gin.New()
			router.POST("/api/curl", handler.GenerateCurl)

			// Create the request body
			body, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest(http.MethodPost, "/api/curl", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, response)

			// Verify that the mock expectations were met
			mockService.AssertExpectations(t)
		})
	}
}
