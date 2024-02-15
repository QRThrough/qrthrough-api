package repo

import (
	"fmt"

	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) port.AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r accountRepo) Create(account *model.Account) error {
	if err := r.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r accountRepo) Get(key model.Key, value interface{}) (*model.Account, error) {
	var account model.Account
	conditionKey := fmt.Sprintf("%s=?", key)
	if err := r.db.Take(&account, conditionKey, value).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r accountRepo) All(query domain.DashboardUserQuery, isModerator bool) ([]model.Account, int64, error) {
	var count int64
	var accounts []model.Account

	var queryRole string

	if isModerator {
		queryRole = "role = 'ADMIN' OR role = 'MODERATOR'"
	} else {
		queryRole = "role = 'USER'"
	}

	tx := r.db.Model(&model.Account{}).Preload(clause.Associations)
	query.AccountSearchByQuery(tx.Where(queryRole))
	tx.Count(&count)

	if err := tx.Find(&accounts).Error; err != nil {
		return nil, 0, err
	}

	return accounts, count, nil
}

func (r accountRepo) UpdateById(id int, account *model.Account) error {
	return r.db.Where("account_id=?", id).Select("Firstname", "Lastname", "Tel", "IsActive", "Role").Updates(&account).Error
}

func (r accountRepo) DeleteById(id int) error {
	return r.db.Where("account_id=?", id).Delete(&model.Account{}).Error
}
