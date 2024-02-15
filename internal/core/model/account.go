package model

import (
	"time"

	"gorm.io/gorm"
)

type Flag string

const (
	FLAG_NOTFOUND Flag = "NOTFOUND"
	FLAG_FOUND    Flag = "FOUND"
	FLAG_EDIT     Flag = "EDIT"
)

type Role string

const (
	ROLE_USER      Role = "USER"
	ROLE_MODERATOR Role = "MODERATOR"
	ROLE_ADMIN     Role = "ADMIN"
)

type Key string

const (
	ACCOUNT_ID = "account_id"
	LINE_ID    = "line_id"
)

type Account struct {
	ID        int             `gorm:"primaryKey;column:account_id;" json:"account_id"`
	LineID    string          `gorm:"column:line_id;type:varchar(64);not null;unique" json:"line_id"`
	Firstname string          `gorm:"column:firstname;type:varchar(256);not null;" json:"firstname"`
	Lastname  string          `gorm:"column:lastname;type:varchar(256);not null;" json:"lastname"`
	Tel       string          `gorm:"column:tel;type:varchar(15);not null;unique" json:"tel"`
	Flag      Flag            `gorm:"column:flag;not null" json:"flag"`
	Role      Role            `gorm:"column:role;default:'USER';not null;" json:"role"`
	IsActive  bool            `gorm:"column:is_active;type:boolean;default:true;not null" json:"is_active"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
