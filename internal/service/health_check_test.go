package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Implement your test logic here
	cfg := config.GetConfig()
	t.Parallel()
	testCases := []struct {
		name                string
		expectedMessage     string
		expectedServiceName string
		expectedInstanceId  uuid.UUID
		expectError         error
	}{
		{
			name:                "Health check returns OK",
			expectedMessage:     config.OK_STATUS,
			expectedServiceName: config.BOOKMARK_SERVICE,
			expectedInstanceId:  uuid.New(),
			expectError:         nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the HealthCheck function and compare the result with tc.expectedStatus
			t.Parallel()
			svc := NewHealthCheck()
			result, err := svc.GenerateHealthCheck()
			assert.Equal(t, tc.expectedMessage, config.OK_STATUS)
			assert.Equal(t, tc.expectedServiceName, config.BOOKMARK_SERVICE)
			assert.IsType(t, uuid.UUID{}, result.InstanceID)
			assert.ErrorIs(t, err, tc.expectError)
		})
	}
}
