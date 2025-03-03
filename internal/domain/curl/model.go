package curl

// Request represents the curl command request parameters
type Request struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	QueryParams map[string]string `json:"queryParams"`
}

// Service defines the interface for curl command generation
type Service interface {
	GenerateCurlCommand(req Request) (string, error)
}

// Repository interface (for future persistence if needed)
type Repository interface {
	SaveCommand(command string) error
	GetCommands() ([]string, error)
} 