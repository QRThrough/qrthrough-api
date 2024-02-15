package model

type Scanner struct {
	Mac      string `gorm:"primaryKey;column:mac;type:varchar(20)" json:"id"`
	Desc     string `gorm:"column:desc;type:varchar(255);not null;" json:"desc"`
	IsActive bool   `gorm:"column:is_active;type:boolean;default:true;not null" json:"is_active"`
}
