package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pvlearning002/go_assignment01/internal/config"
	services_mocks "github.com/pvlearning002/go_assignment01/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	// Implement your test logic here
	t.Parallel()
	testCases := []struct {
		name             string
		setupRequest     func(c *gin.Context)
		setupMockService func(ctx context.Context) *services_mocks.HealthCheckMock
		expectedStatus   int
		expectedResponse bool
	}{
		{
			name: "normal case",
			setupRequest: func(c *gin.Context) {
				// Setup your request here
				c.Request = httptest.NewRequest(http.MethodGet, config.HEALTH_CHECK_GET_PATH, nil)
			},
			setupMockService: func(ctx context.Context) *services_mocks.HealthCheckMock {
				mockService := services_mocks.NewHealthCheckMock(t)
				// Setup your mock service here
				response := mockService.GenerateHealthCheck()

				mockService.On("GenerateHealthCheck").Return(response)
				return mockService
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Implement your test logic here
			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)
			tc.setupRequest(ginCtx)
			mockService := tc.setupMockService(context.Background())
			handler := NewHealthCheck(mockService)
			handler.GenerateHealthCheck(ginCtx)
			if rec.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rec.Code)
			}
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String() == `{"message":config.OK_STATUS,"serviceName":config.BOOKMARK_SERVICE,"instanceID":config.HEALTH_CHECK_ID_DEFAULT}`)
		})
	}
}
