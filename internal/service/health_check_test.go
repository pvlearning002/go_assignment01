package service

import (
	"testing"

	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/pvlearning002/go_assignment01/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Implement your test logic here
	t.Parallel()
	testCases := []struct {
		name                string
		expectedMessage     string
		expectedServiceName string
		expectedInstanceId  string
		expectError         error
	}{
		{
			name:                "Health check returns OK",
			expectedMessage:     config.OK_STATUS,
			expectedServiceName: config.BOOKMARK_SERVICE,
			expectedInstanceId:  config.HEALTH_CHECK_ID_DEFAULT,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the HealthCheck function and compare the result with tc.expectedStatus
			t.Parallel()
			svc := NewHealthCheck()
			var result *model.HealthCheck = svc.GenerateHealthCheck()
			assert.Equal(t, tc.expectedMessage, result.Message)
			assert.Equal(t, tc.expectedServiceName, result.ServiceName)
			assert.Equal(t, tc.expectedInstanceId, result.InstanceID)
		})
	}
}
