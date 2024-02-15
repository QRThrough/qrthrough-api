package port

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
)

type AccountRepo interface {
	Create(account *model.Account) error
	Get(key model.Key, value interface{}) (*model.Account, error)
	All(query domain.DashboardUserQuery, isModerator bool) ([]model.Account, int64, error)
	UpdateById(id int, account *model.Account) error
	DeleteById(id int) error
}
