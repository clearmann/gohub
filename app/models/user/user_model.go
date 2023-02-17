package user

import "gohub/app/models"

type User struct {
	models.BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	password string `json:"password"`
	models.CommonTimestampsField
}
