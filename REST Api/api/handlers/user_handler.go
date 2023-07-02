package handler

import "github.com/gin-gonic/gin"

type UserData struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var Data UserData

	if c.Bind(&Data) != nil {
		c.JSON()
	}
}
