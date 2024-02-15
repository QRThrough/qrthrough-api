package service

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
)

type dashboardModeratorService struct {
	accountRepo port.AccountRepo
}

func NewDashboardModeratorService(accountRepo port.AccountRepo) domain.DashboardModeratorService {
	return &dashboardModeratorService{
		accountRepo: accountRepo,
	}
}

func (s dashboardModeratorService) All(query domain.DashboardUserQuery) (*dto.AllAccountsResponseBody, error) {
	moderators, count, err := s.accountRepo.All(query, true)
	if err != nil {
		return nil, err
	}

	return &dto.AllAccountsResponseBody{
		Count:    count,
		Accounts: moderators,
	}, nil
}
func (s dashboardModeratorService) Update(id int, body dto.AccountRequestBody) error {

	account := model.Account{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Tel:       body.Tel,
		IsActive:  body.IsActive,
		Role:      body.Role,
	}

	if err := s.accountRepo.UpdateById(id, &account); err != nil {
		return err
	}

	return nil
}

func (s dashboardModeratorService) Delete(id int) error {
	if err := s.accountRepo.DeleteById(id); err != nil {
		return err
	}
	return nil
}
