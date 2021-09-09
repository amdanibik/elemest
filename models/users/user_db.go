package users

import (
	"time"

	"gorm.io/gorm"
)

/*
type User struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	Username  string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
*/

type User struct {
	Id        uint           `json:"id"`
	Name      string         `json:"string"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
