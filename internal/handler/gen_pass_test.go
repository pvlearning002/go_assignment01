package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	gin "github.com/gin-gonic/gin"
	"github.com/pvlearning002/go_assignment01/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePasswordHandler_GenPass(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name             string
		setupRequest     func(ctx *gin.Context)
		setupMockService func(ctx context.Context) *mocks.GenPass
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "success",
			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/genpass", nil)
			},
			setupMockService: func(ctx context.Context) *mocks.GenPass {
				mockService := mocks.NewGenPass(t)
				mockService.On("GenPass", ctx).Return("mockedPassword", nil)
				return mockService
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: "mockedPassword",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Here you would call your GeneratePasswordHandler.GenPass method
			// and assert the results, similar to the service test.
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			tc.setupRequest(ctx)
			mockSvc := tc.setupMockService(ctx)
			testHandler := NewGenPass(mockSvc)
			testHandler.GeneratePassword(ctx)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}

}
