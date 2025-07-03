package userservice

import (
	"time"

	"gorm.io/gorm"
)

type UsersOrm struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UsersOrm) TableName() string {
	return "users"
}
