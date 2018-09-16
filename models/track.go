package models

type Track struct {
	Id      uint   `json:"id"`
	User_id uint   `json:"user_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
