package models

import (
	"time"
)

type User struct {
	ID    uint   `json:"id" gorm:"primary_key";"AUTO_INCREMENT"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	Name  string `json:"name" gorm:"max:256"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
