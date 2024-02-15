package model

import (
	"time"
)

type Usage struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	QRCodeID  int64      `gorm:"column:qrcode_id;not null" json:"qrcode_id"`
	AccountID int        `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`

	QRCode  QRCode  `gorm:"foreignKey:QRCodeID;onDelete:CASCADE" json:"qrcode"`
	Account Account `gorm:"foreignKey:AccountID;onDelete:CASCADE" json:"account"`
}
