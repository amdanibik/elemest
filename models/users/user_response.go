package users

import (
	"time"
)

type UserToken struct {
	Id        uint      `json:"id"`
	Name      string    `json:"string"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Token     string    `json:"token"`
}

type UserResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type UserResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UserResponseCount struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Jumlah  int64  `json:"jumlah"`
}

type UserResponseToken struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    UserToken `json:"data"`
}
