package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		input          http.Header
		expectedString string
		expectedErr    error
	}{
		{"Case 1", http.Header{}, ErrNoAuthHeaderIncluded.Error(), ErrNoAuthHeaderIncluded},
		{"Case 2", http.Header{"Authorization": []string{"ApiKey"}}, errors.New("malformed authorization header").Error(), errors.New("malformed authorization header")},
		{"Case 3", http.Header{"Authorization": []string{"ApiKey my-api-key"}}, "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAPIKey(tt.input)

			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("expected error %v, got %v. result was %q", tt.expectedErr, err, result)
			}
		})
	}
}
