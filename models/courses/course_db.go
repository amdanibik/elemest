package courses

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Category  string         `json:"category"`
	Level     string         `json:"level"`
	Price     int            `json:"price"`
	IsFree    int            `json:"isfree"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
