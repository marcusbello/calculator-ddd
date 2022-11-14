package main

import (
	"calculator-ddd/pkg/database/postgresql"
	"calculator-ddd/pkg/database/redis_cache"
	"calculator-ddd/pkg/http"
	"calculator-ddd/pkg/repository"
	"calculator-ddd/pkg/usecase"
	"calculator-ddd/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := utils.GetConfigs()

	postgresDb, err := postgresql.InitDatabase(config.PostgresConfig)
	redisDb, err := redis_cache.InitRedisCache(config.RedisConfig)

	log.Println("successful db and cache connection")
	r := gin.Default()
	calculateRepo := repository.NewCalculateRepository(postgresDb, redisDb)
	calculateUseCase := usecase.NewCalculateUseCase(calculateRepo)

	http.NewCalculateHandler(r, calculateUseCase)

	err = r.Run(config.Port)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
