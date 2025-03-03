package domain

// CurlRequest represents the curl command request parameters.
type CurlRequest struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	QueryParams map[string]string `json:"queryParams"`
}
