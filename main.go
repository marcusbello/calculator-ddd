package main

import (
	"calculator-ddd/database/postgresql"
	"calculator-ddd/http"
	"calculator-ddd/repository"
	"calculator-ddd/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.GetConfigs()

	postgresDb, err := postgresql.InitDatabase(config.PostgresConfig)

	r := gin.Default()
	calculateRepo := repository.NewCalculateRepository(postgresDb)
	calculateUseCase := usecase.NewCalculateUseCase(calculateRepo)

	http.NewCalculateHandler(r, calculateUseCase)

	err = r.Run(config.Port)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
