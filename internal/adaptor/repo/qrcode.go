package repo

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type qrCodeRepo struct {
	db *gorm.DB
}

func NewQRCodeRepo(db *gorm.DB) port.QRCodeRepo {
	return &qrCodeRepo{
		db: db,
	}
}

func (r qrCodeRepo) Create(qrcode *model.QRCode) error {
	return r.db.Create(qrcode).Error
}

func (r qrCodeRepo) GetById(id int64) (*model.QRCode, error) {
	var qrcode model.QRCode
	if err := r.db.
		Preload(clause.Associations).
		Take(&qrcode, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &qrcode, nil
}
