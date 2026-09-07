package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pvlearning002/go_assignment01/internal/api"
	"github.com/pvlearning002/go_assignment01/internal/config"
)

func TestGenerateHealthCheckEndpoint(t *testing.T) {
	// Implement your integration test logic here
	t.Parallel()
	testCases := []struct {
		name string
		// Add other fields as needed for your test cases
		setupTestApiEngine func(api api.Engine) *httptest.ResponseRecorder
		expectedStatus     int
		expectedResponse   bool
	}{
		{
			name: "normal case",
			// Initialize other fields as needed
			setupTestApiEngine: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest(http.MethodGet, config.HEALTH_CHECK_GET_PATH, nil)
				rec := httptest.NewRecorder()
				api.ServeHTTP(rec, req)
				return rec
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Implement your test logic here
			testAPI := api.NewEngine()
			rec := tc.setupTestApiEngine(testAPI)
			if rec.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rec.Code)
			}
			if rec.Body.String() == `{"message":config.OK_STATUS,"serviceName":config.BOOKMARK_SERVICE,"instanceID":config.HEALTH_CHECK_ID_DEFAULT}` != tc.expectedResponse {
				t.Errorf("expected response %v, got %s", tc.expectedResponse, rec.Body.String())
			}

		})
	}
}
