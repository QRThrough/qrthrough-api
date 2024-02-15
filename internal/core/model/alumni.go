package model

import "time"

type Alumni struct {
	ID        int        `gorm:"primaryKey;column:student_code;" json:"student_code"`
	Firstname string     `gorm:"column:firstname;type:varchar(256);not null;" json:"firstname"`
	Lastname  string     `gorm:"column:lastname;type:varchar(256);" json:"lastname"`
	Tel       string     `gorm:"column:tel;type:varchar(15);" json:"tel"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
