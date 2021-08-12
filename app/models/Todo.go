package models

import "time"

type Todo struct {
	Id          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"size:255"`
	DueDate     time.Time `json:"due_date" gorm:"default:CURRENT_TIMESTAMP"`
}
