package http

import (
	"calculator-ddd/docs"
	"calculator-ddd/pkg/domain"
	"calculator-ddd/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type CalculateHandler struct {
	calculateUseCase domain.CalculateUseCase
}

// @title          Calculator API
// @version        1.0
// @description    This is a sample server calculator server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3330
// @BasePath /api/v1

func NewCalculateHandler(r *gin.Engine, calculateUc domain.CalculateUseCase) {
	handler := &CalculateHandler{calculateUseCase: calculateUc}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/calculate", handler.GetCalculation)
			eg.GET("/calculate/history", handler.GetCalculationHistory)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//r.GET("/calculate", handler.GetCalculation)
	//r.GET("/calculate/history", handler.GetCalculationHistory)
}

// GetCalculation godoc
// @Summary Calculate
// @Schemes
// @Description calculate two integers
// @Tags        calculator
// @Produce     json
// @Param       first  query    string true "first input"
// @Param       second query    string true "second input"
// @Success     200    {object} utils.Response
// @Router      /calculate [get]
func (h CalculateHandler) GetCalculation(c *gin.Context) {
	var request domain.CalculateGetRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	sum, sub, times, div, err := h.calculateUseCase.GetCalculationUc(c, request.First, request.Second)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Data:  fmt.Sprintf(`{Addition: %v, Subtraction:%v, Multiply: %v, Division:%v}`, sum, sub, times, div),
		Error: nil,
	})
}

// GetCalculationHistory godoc
// @Summary Calculation History
// @Schemes
// @Description get the latest calculations from cache
// @Tags        calculator
// @Produce     json
// @Success     200 {object} utils.Response
// @Router      /calculate/history [get]
func (h CalculateHandler) GetCalculationHistory(c *gin.Context) {
	res, err := h.calculateUseCase.GetCalculationHistoryUc(c)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Data:  fmt.Sprintf("%v", res),
		Error: nil,
	})
}
