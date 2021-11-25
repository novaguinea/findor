package models

import(
	// "gorm.io/gorm"
	"time"

)

type Users struct {
	ID			uint32		`json:"id" gorm:"primary_key:auto_increment"`
	Name		string		`json:"name" gorm:"size:255;not null"`
	Email		string		`json:"email" gorm:"size:255;not null;unique"`
	Password	string		`json:"password" gorm:"size:255;not null"`
	CreateAt	time.Time	`json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt	time.Time	`json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`	
}