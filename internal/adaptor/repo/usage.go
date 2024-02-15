package repo

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) port.UsageRepo {
	return &logRepo{
		db: db,
	}
}

func (r logRepo) Create(log *model.Usage) error {
	return r.db.Create(log).Error
}

func (r logRepo) All(query domain.DashboardLogQuery) ([]model.Usage, int64, error) {
	var count int64
	var logs []model.Usage

	tx := r.db.Model(&model.Usage{}).Preload(clause.Associations)
	query.LogSearchByQuery(tx)
	tx.Count(&count)

	if err := tx.Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, count, nil
}

func (r logRepo) CountByQRCode(id int64) (int64, error) {
	var count int64
	var logs []model.Usage
	if err := r.db.Where("qrcode_id=?", id).Find(&logs).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
