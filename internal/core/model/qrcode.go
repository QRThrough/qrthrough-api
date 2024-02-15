package model

import (
	"time"
)

type QRCode struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	AccountID int        `gorm:"column:account_id;not null" json:"account_id"`
	ExpireAt  *time.Time `gorm:"column:expire_at;not null" json:"expire_at"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`

	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
