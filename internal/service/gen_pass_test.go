package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		expectError    error
	}{
		{
			name:           "success",
			expectedLength: 12,
			expectError:    nil,
		},
		{
			name:           "success with custom length 5",
			expectedLength: 5,
			expectError:    nil,
		},
		{
			name:           "success with custom length 120",
			expectedLength: 120,
			expectError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testSvc := NewGenPass()
			pass, err := testSvc.GeneratePassword(tc.expectedLength)
			assert.Equal(t, tc.expectError, err)
			assert.Equal(t, tc.expectedLength, len(pass))
		})
	}
}
