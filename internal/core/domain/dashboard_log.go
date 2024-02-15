package domain

import (
	"fmt"

	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"gorm.io/gorm"
)

type DashboardLogService interface {
	All(DashboardLogQuery) (*dto.AllLogsResponseBody, error)
}

type DashboardLogQuery struct {
	Type  QueryType  `query:"type"`
	Value string     `query:"value"`
	Order QueryOrder `query:"order"`
	Sort  QuerySort  `query:"sort"`
}

func (obj DashboardLogQuery) LogSearchByQuery(tx *gorm.DB) *gorm.DB {
	query := ""
	order := ""

	switch obj.Order {
	case ORDER_STUDENT_CODE:
		order = fmt.Sprintf("%v %v", "accounts.account_id", obj.Sort)
	case ORDER_NAME:
		order = fmt.Sprintf("%v %v", "(accounts.firstname,accounts.lastname)", obj.Sort)
	case ORDER_TEL:
		order = fmt.Sprintf("%v %v", "accounts.tel", obj.Sort)
	case ORDER_DATE:
		order = fmt.Sprintf("%v %v", "created_at", obj.Sort)
	}

	switch obj.Type {
	case TYPE_STUDENT_CODE:
		query = "CAST(accounts.account_id AS TEXT) LIKE ?"
	case TYPE_NAME:
		query = "CONCAT(accounts.firstname, ' ', accounts.lastname) LIKE ?"
	case TYPE_TEL:
		query = "accounts.tel LIKE ?"
	default:
		return tx
	}

	searchValue := fmt.Sprintf("%%%v%%", obj.Value)

	return tx.Joins("JOIN accounts ON accounts.account_id = usages.account_id").Where(query, searchValue).Order(order)
}
