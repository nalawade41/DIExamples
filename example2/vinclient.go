package example2

import "testing"

// Abstraction for client
type VINAPIClient interface {
	IsValidVinCode(code string) bool
}

type webAPIClient struct {
	apiURL string
	apiKey string
	// .. internals go here ...
}

func NewVINAPIClient(apiURL, apiKey string) *webAPIClient {
	return &webAPIClient{apiURL, apiKey}
}

func (client *webAPIClient) IsValidVinCode(code string) bool {
	// calls external API and returns something more useful
	return true
}

// Service layer
type VINService struct {
	client VINAPIClient
}

type VINServiceConfig struct {
	// more configuration values
}

// Cretes service using configuration and client
func NewVINService(config *VINServiceConfig, apiClient VINAPIClient) *VINService {

	// apiClient is created elsewhere and injected here
	return &VINService{apiClient}
}

// Now while testing  we could just do as below
// we will create mock client and use it for testing
type mockAPIClient struct {
	apiCalls int
}

func NewMockAPIClient() *mockAPIClient {
	return &mockAPIClient{}
}

func (client *mockAPIClient) IsValidVinCode(code string) bool {
	client.apiCalls++
	return true
}

// Test case for example
func TestVIN_IsEuropeanr(t *testing.T) {
	apiClient := NewMockAPIClient()
	service := NewVINService(&VINServiceConfig{}, apiClient)
	_ = service.client.IsValidVinCode("testcode")
	if apiClient.apiCalls != 1 {
		t.Errorf("unexpected number of API calls: %d", apiClient.apiCalls)
	}
}
