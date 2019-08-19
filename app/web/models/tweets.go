package models

type Tweet struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Tweet  string `json:"tweet"`
	Email  string `json:"email"`
}
