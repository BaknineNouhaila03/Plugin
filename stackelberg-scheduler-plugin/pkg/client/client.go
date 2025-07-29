package stackelberg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Request structure for the Python service
type AllocateRequest struct {
	TotalCPU    float64                `json:"total_cpu"`
	TotalMemory float64                `json:"total_memory"`
	Params      map[string]interface{} `json:"params"`
}

// Response structure from the Python service
type AllocateResponse struct {
	Allocations     map[string]interface{} `json:"allocations"`
	Prices          map[string]interface{} `json:"prices"`
	PlatformUtility float64                `json:"platform_utility"`
	Metrics         map[string]interface{} `json:"metrics"`
	Success         bool                   `json:"success"`
	Error           string                 `json:"error,omitempty"`
}

// Client configuration
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}



// NewClient creates a new Stackelberg client
func NewClient(baseURL string) *Client {
	if baseURL == "" {
		baseURL = "http://localhost:5000"
	}

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CallPythonSidecar calls the Python service to get resource allocation
func CallPythonSidecar(totalCPU, totalMemory float64, params map[string]interface{}) (*AllocateResponse, error) {
	client := NewClient("")
	return client.Allocate(totalCPU, totalMemory, params)
}

// Allocate calls the /stackelberg/allocate endpoint
func (c *Client) Allocate(totalCPU, totalMemory float64, params map[string]interface{}) (*AllocateResponse, error) {
	// Prepare request payload
	request := AllocateRequest{
		TotalCPU:    totalCPU,
		TotalMemory: totalMemory,
		Params:      params,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/stackelberg/allocate", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result AllocateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if the service returned an error
	if !result.Success {
		return nil, fmt.Errorf("service error: %s", result.Error)
	}

	return &result, nil
}

// Health checks if the Python service is running
func (c *Client) Health() error {
	url := fmt.Sprintf("%s/health", c.BaseURL)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service unhealthy, status: %d", resp.StatusCode)
	}

	return nil
}
