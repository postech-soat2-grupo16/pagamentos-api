package external

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDialector() gorm.Dialector {
	connStr := os.Getenv("DATABASE_URL")
	pgDialector := postgres.Open(connStr)
	fmt.Printf("DB connected: %v\n", pgDialector.Name())
	return pgDialector
}
