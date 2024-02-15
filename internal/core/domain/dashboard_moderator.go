package domain

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
)

type DashboardModeratorService interface {
	All(DashboardUserQuery) (*dto.AllAccountsResponseBody, error)
	Update(id int, body dto.AccountRequestBody) error
	Delete(id int) error
}
