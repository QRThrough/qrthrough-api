package repo

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type scannerRepo struct {
	db *gorm.DB
}

func NewScannerRepo(db *gorm.DB) port.ScannerRepo {
	return &scannerRepo{
		db: db,
	}
}

func (r scannerRepo) CheckExistedScanner(mac string) bool {
	var scanner model.Scanner
	if err := r.db.
		Preload(clause.Associations).
		Take(&scanner, "mac=?", mac).Error; err != nil {
		return false
	}

	return true
}
