package service

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
)

type dashboardService struct {
	accountRepo       port.AccountRepo
	configurationRepo port.ConfigurationRepo
}

func NewDashboardService(accountRepo port.AccountRepo, configurationRepo port.ConfigurationRepo) domain.DashboardService {
	return &dashboardService{
		accountRepo:       accountRepo,
		configurationRepo: configurationRepo,
	}
}

func (s dashboardService) GetRole(uid string) (model.Role, error) {
	account, err := s.accountRepo.Get(model.LINE_ID, uid)
	if err != nil {
		return "", err
	}
	return account.Role, nil
}

func (s dashboardService) AllConfiguration() ([]dto.Configuration, error) {
	configurations, err := s.configurationRepo.All()
	if err != nil {
		return nil, err
	}

	var configurationsResult []dto.Configuration
	for _, configuration := range configurations {
		configurationResult := dto.Configuration{
			Key:   configuration.Key,
			Desc:  configuration.Desc,
			Value: configuration.Value,
		}
		configurationsResult = append(configurationsResult, configurationResult)
	}

	return configurationsResult, nil
}

func (s dashboardService) UpdateConfiguration(body []dto.Configuration) error {
	var configurations []model.Configuration
	for _, configuration := range body {
		configurationMapped := model.Configuration{
			Key:   configuration.Key,
			Desc:  configuration.Desc,
			Value: configuration.Value,
		}
		configurations = append(configurations, configurationMapped)
	}

	if err := s.configurationRepo.Updates(configurations); err != nil {
		return err
	}
	return nil
}
