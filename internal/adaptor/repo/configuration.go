package repo

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type configurationRepo struct {
	db *gorm.DB
}

func NewConfigurationRepo(db *gorm.DB) port.ConfigurationRepo {
	return &configurationRepo{
		db: db,
	}
}

func (r configurationRepo) All() ([]model.Configuration, error) {
	var configurations []model.Configuration
	if err := r.db.Preload(clause.Associations).Find(&configurations).Error; err != nil {
		return nil, err
	}
	return configurations, nil
}

func (r configurationRepo) Create(configuration *model.Configuration) error {
	if err := r.db.Create(configuration).Error; err != nil {
		return err
	}
	return nil
}

func (r configurationRepo) GetByKey(key string) (*model.Configuration, error) {
	var configuration model.Configuration
	if err := r.db.Take(&configuration, "key=?", key).Error; err != nil {
		return nil, err
	}
	return &configuration, nil
}

func (r configurationRepo) Updates(configurations []model.Configuration) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, config := range configurations {
		if err := tx.Model(&model.Configuration{}).Where("key = ?", config.Key).Updates(config).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
