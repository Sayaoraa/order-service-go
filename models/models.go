package models

import (
	"gorm.io/gorm"
	"time"
	_ "time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Password  string     `gorm:"not null" json:"-"`
}

type Order struct {
	gorm.Model
	OrderID      string
	UserID       uint
	Status       string
	TrackingInfo string
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Order{})
}
