package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedKey    string
		expectingError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			expectedKey:    "valid_api_key",
			expectingError: false,
		},
		{
			name:           "Missing Authorization Header",
			headers:        http.Header{},
			expectedKey:    "",
			expectingError: true,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Bearer some_token"},
			},
			expectedKey:    "",
			expectingError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if tt.expectingError && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tt.expectingError && err != nil {
				t.Errorf("did not expect an error but got one: %v", err)
			}
			if apiKey != tt.expectedKey {
				t.Errorf("expected API key %s, but got %s", tt.expectedKey, apiKey)
			}
		})
	}
}
