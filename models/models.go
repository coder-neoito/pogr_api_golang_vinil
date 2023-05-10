package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"primaryKey; size:255"`
	UserName      string `gorm:"size:255; not null; unique"`
	FullName      string `gorm:"size:255"`
	Email         string `gorm:"size:255; not null; unique"`
	Password      string `gorm:"size:255; not null"`
	EmailVerified bool   `gorm:"default:false"`
	ProfileImgUrl string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Game struct {
	ID          string `gorm:"primaryKey; size:255"`
	Name        string `gorm:"size:255; not null; unique"`
	CompanyName string `gorm:"size:255; not null;"`
	Description string `gorm:"size:255"`
	ImageUrl    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserGameProfile struct {
	ID     string `gorm:"primaryKey; size:255"`
	UserID string `gorm:"not null; uniqueIndex:ugp_user_game"`
	GameID string `gorm:"not null; uniqueIndex:ugp_user_game"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"ForeignKey:UserID;References:ID"`
	Game Game `gorm:"ForeignKey:GameID;References:ID"`
}
