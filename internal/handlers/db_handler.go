package handlers

import (
	"fmt"
	"os"

	"github.com/vinilunni/pogr_api_golang_vinil/internal/common/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dbinf := &types.DbInfo{}

	dbinf.Host = os.Getenv("POSTGRES_HOST")
	dbinf.User = os.Getenv("POSTGRES_USER")
	dbinf.Password = os.Getenv("POSTGRES_PASSWORD")
	dbinf.Name = os.Getenv("POSTGRES_NAME")
	dbinf.Port = os.Getenv("POSTGRES_PORT")
	dbinf.Sslmode = os.Getenv("POSTGRES_SSLMODE")
	dbinf.Timezone = os.Getenv("POSTGRES_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbinf.Host, dbinf.User, dbinf.Password, dbinf.Name, dbinf.Port, dbinf.Sslmode, dbinf.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB connection err")
	}

	return db
}
