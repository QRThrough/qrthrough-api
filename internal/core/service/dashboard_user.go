package service

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
)

type dashboardUserService struct {
	accountRepo port.AccountRepo
}

func NewDashboardUserService(accountRepo port.AccountRepo) domain.DashboardUserService {
	return &dashboardUserService{
		accountRepo: accountRepo,
	}
}

func (s dashboardUserService) All(query domain.DashboardUserQuery) (*dto.AllAccountsResponseBody, error) {
	users, count, err := s.accountRepo.All(query, false)
	if err != nil {
		return nil, err
	}

	return &dto.AllAccountsResponseBody{
		Count:    count,
		Accounts: users,
	}, nil
}

func (s dashboardUserService) Update(id int, body dto.AccountRequestBody) error {

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

func (s dashboardUserService) Delete(id int) error {
	if err := s.accountRepo.DeleteById(id); err != nil {
		return err
	}
	return nil
}
