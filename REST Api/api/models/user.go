package models

type User struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
