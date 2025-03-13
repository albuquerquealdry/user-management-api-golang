package dto

import "time"

type UserResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Birthday    string    `json:"birthday"`
	Address     string    `json:"address"`
	Postalcode  int       `json:"postalcode"`
	CPF         int       `json:"cpf"`
	Nationality string    `json:"nationality"`
	Score       int       `json:"score"`
	Status      string    `json:"status"`
	MotherName  string    `json:"mother_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
