package port

import "github.com/JMjirapat/qrthrough-api/internal/core/model"

type ConfigurationRepo interface {
	All() ([]model.Configuration, error)
	Create(configuration *model.Configuration) error
	GetByKey(key string) (*model.Configuration, error)
	Updates(configuration []model.Configuration) error
}
