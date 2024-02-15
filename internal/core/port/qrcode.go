package port

import "github.com/JMjirapat/qrthrough-api/internal/core/model"

type QRCodeRepo interface {
	Create(qrcode *model.QRCode) error
	GetById(id int64) (*model.QRCode, error)
}
