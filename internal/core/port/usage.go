package port

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
)

type UsageRepo interface {
	Create(*model.Usage) error
	All(query domain.DashboardLogQuery) ([]model.Usage, int64, error)
	CountByQRCode(id int64) (int64, error)
}
