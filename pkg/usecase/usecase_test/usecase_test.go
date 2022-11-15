package usecase_test

import (
	"calculator-ddd/pkg/domain"
	"calculator-ddd/pkg/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCalculateUseCase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	//err := errors.New("error")
	t.Run("Get Empty Data", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationHistoryRepository", c).Return([]domain.CalculationHistory{}, nil).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		uc, err := u.GetCalculationHistoryUc(c)
		assert.NoError(t, err)
		assert.NotNil(t, uc)
	})
}
