package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CalculateGetRequest example
type CalculateGetRequest struct {
	First  string `form:"first" binding:"required"`
	Second string `form:"second" binding:"required"`
}

type CalculationHistory struct {
	gorm.Model
	FirstInteger  int
	SecondInteger int
	Sum           int
	Subtract      int
	Times         int
	Divide        float64
}

type CalculationHistoryResponse struct {
	ID            uint `json:"ID"`
	FirstInteger  int  `json:"firstInteger"`
	SecondInteger int  `json:"secondInteger"`
	Sum           int  `json:"sum"`
	Times         int  `json:"times"`
	Divide        int  `json:"divide"`
}

type CalculateUseCase interface {
	GetCalculationUc(ctx *gin.Context, a, b string) (int, int, int, float64, error)
	GetCalculationHistoryUc(ctx *gin.Context) ([]CalculationHistory, error)
}

type CalculateRepository interface {
	GetCalculationRepository(ctx *gin.Context, a, b int) (int, int, int, float64, error)
	Sum(a, b int) int
	Subtract(a, b int) int
	Times(a, b int) int
	Divide(a, b int) float64
	GetCalculationHistoryRepository(ctx *gin.Context) ([]CalculationHistory, error)
}
