package model

import (
	"time"

	"gorm.io/gorm"
)

type CredentialDB struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Todo struct {
	gorm.Model
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Session struct {
	gorm.Model
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type ToggleTodoReq struct {
	ID   int  `json:"id"`
	Done bool `json:"done"`
}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
