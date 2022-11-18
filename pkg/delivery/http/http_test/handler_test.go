package http_test

import (
	calculatehandler "calculator-ddd/pkg/delivery/http"
	"calculator-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler_GetCalculation(t *testing.T) {
	gin.SetMode("test")
	//err := errors.New("error")
	t.Run("Get Data", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockGetCalculationUseCase := domain.NewMockCalculateUseCase(t)
		mockGetCalculationUseCase.On("GetCalculationUc", mock.Anything, "10", "10").Return(
			10, 10, 10, float64(10), nil)
		calculatehandler.NewCalculateHandler(engine, mockGetCalculationUseCase)
		req, err := http.NewRequest(http.MethodGet, "/calculate?first=10&second=10", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
