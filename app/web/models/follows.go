package models

type Follow struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	To     int `json:"to"`
}
