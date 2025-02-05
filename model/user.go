package model

import "time"

type User struct {
	ID           string    `gorm:"primary_key;type:varchar(255);not null" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	NickName     string    `gorm:"type:varchar(255)" json:"nick_name"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"password_hash"`
	Phone        *string   `gorm:"type:varchar(20)" json:"phone"`
	Email        *string   `gorm:"type:varchar(100)" json:"email"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserRecord struct {
	ID            string `gorm:"primary_key;type:varchar(255);not null" json:"id"`
	UserID        string `gorm:"type:varchar(255);not null" json:"user_id"`
	StoryID       string `gorm:"type:varchar(255);not null" json:"story_id"`
	QuestionCount int    `gorm:"type:int;not null" json:"question_count"`
}
