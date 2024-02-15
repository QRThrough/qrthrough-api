package service

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
)

type dashboardLogService struct {
	repo port.UsageRepo
}

func NewDashboardLogService(repo port.UsageRepo) domain.DashboardLogService {
	return &dashboardLogService{
		repo: repo,
	}
}

func (s dashboardLogService) All(query domain.DashboardLogQuery) (*dto.AllLogsResponseBody, error) {
	logs, count, err := s.repo.All(query)
	if err != nil {
		return nil, err
	}

	return &dto.AllLogsResponseBody{
		Count: count,
		Logs:  logs,
	}, nil
}
