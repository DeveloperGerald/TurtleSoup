package model

import "time"

type Story struct {
	ID        string    `gorm:"primary_key;type:varchar(255);not null" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Riddle    string    `gorm:"type:text;not null" json:"riddle"`
	Answer    string    `gorm:"type:text;not null" json:"answer"`
	Owner     string    `gorm:"type:varchar(255);not null" json:"owner"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}
