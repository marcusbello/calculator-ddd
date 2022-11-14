package postgresql

import (
	"calculator-ddd/pkg/domain"
	"calculator-ddd/pkg/utils"
	"fmt"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config utils.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
		config.Host, config.User, config.Password, config.Db, config.Port,
	)

	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{})
	db.AutoMigrate(domain.CalculationHistory{})

	return db, err
}
