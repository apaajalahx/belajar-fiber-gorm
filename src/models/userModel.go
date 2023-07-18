package models

import "time"

type UserModel struct {
	Id        int       `gorm:"PrimaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
