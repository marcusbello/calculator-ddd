package http

import (
	"calculator-ddd/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalculateHandler struct {
	calculateUseCase domain.CalculateUseCase
}

func NewCalculateHandler(r *gin.Engine, calculateUc domain.CalculateUseCase) {
	handler := &CalculateHandler{calculateUseCase: calculateUc}

	r.GET("/calculate", handler.GetCalculation)
	r.GET("/calculate/history", handler.GetCalculationHistory)
}

func (h CalculateHandler) GetCalculation(c *gin.Context) {
	var request domain.CalculateGetRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	sum, sub, times, div, err := h.calculateUseCase.GetCalculationUc(c, request.First, request.Second)
	c.JSON(http.StatusOK, utils.Response{
		Data:  delivery.MapCalculateResponse{sum, sub, times, div},
		Error: nil,
	})
}

func (h CalculateHandler) GetCalculationHistory(c *gin.Context) {
	res, err := h.calculateUseCase.GetCalculationHistoryUc(c)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Data:  delivery.MapCalculateHistoryResponse(res),
		Error: nil,
	})
}
