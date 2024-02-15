package model

import "time"

type Configuration struct {
	Key       string     `gorm:"primaryKey;column:key;type:varchar(255);" json:"key"`
	Value     string     `gorm:"column:value;type:varchar(255);not null;" json:"value"`
	Desc      string     `gorm:"column:desc;type:varchar(255);" json:"desc"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
