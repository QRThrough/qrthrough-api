package domain

import (
	"fmt"

	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"gorm.io/gorm"
)

type DashboardService interface {
	GetRole(uid string) (model.Role, error)
	AllConfiguration() ([]dto.Configuration, error)
	UpdateConfiguration([]dto.Configuration) error
}

type QueryType string
type QueryOrder string
type QuerySort string

const (
	TYPE_STUDENT_CODE QueryType = "STUDENT CODE"
	TYPE_NAME         QueryType = "NAME"
	TYPE_TEL          QueryType = "TEL"
)

const (
	ORDER_STUDENT_CODE QueryOrder = "STUDENT CODE"
	ORDER_NAME         QueryOrder = "NAME"
	ORDER_TEL          QueryOrder = "TEL"
	ORDER_DATE         QueryOrder = "DATE"
)

const (
	SORT_ASC  QuerySort = "ASC"
	SORT_DESC QuerySort = "DESC"
)

type DashboardUserQuery struct {
	Type  QueryType  `query:"type"`
	Value string     `query:"value"`
	Order QueryOrder `query:"order"`
	Sort  QuerySort  `query:"sort"`
}

func (obj DashboardUserQuery) AccountSearchByQuery(tx *gorm.DB) *gorm.DB {
	query := ""
	order := ""

	switch obj.Order {
	case ORDER_STUDENT_CODE:
		order = fmt.Sprintf("%v %v", "account_id", obj.Sort)
	case ORDER_NAME:
		order = fmt.Sprintf("%v %v", "(firstname,lastname)", obj.Sort)
	case ORDER_TEL:
		order = fmt.Sprintf("%v %v", "tel", obj.Sort)
	case ORDER_DATE:
		order = fmt.Sprintf("%v %v", "CAST(created_at AS DATE)", obj.Sort)
	}

	switch obj.Type {
	case TYPE_STUDENT_CODE:
		query = "CAST(account_id AS TEXT) LIKE ?"
	case TYPE_NAME:
		query = "CONCAT(firstname, ' ', lastname) LIKE ?"
	case TYPE_TEL:
		query = "tel LIKE ?"
	default:
		return tx
	}

	searchValue := fmt.Sprintf("%%%v%%", obj.Value)

	return tx.Where(query, searchValue).Order(order)
}
