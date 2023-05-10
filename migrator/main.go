package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vinilunni/pogr_api_golang_vinil/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInfo struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	Sslmode  string
	Timezone string
}

func DBInit() *gorm.DB {
	dbinf := &DbInfo{}

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

func main() {
	// Load env vars
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db := DBInit()
	Migrate(db)

	log.Println("Successfully migrated!")
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("err create on User")
	}
	if err := db.AutoMigrate(&models.Game{}); err != nil {
		panic("err create on Game")
	}
	if err := db.AutoMigrate(&models.UserGameProfile{}); err != nil {
		panic("err create on UserGameProfile")
	}
}
