package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	errs "github.com/JMjirapat/qrthrough-api/pkg/errors"
	"gorm.io/gorm"
)

type lineService struct {
	qrRepo            port.QRCodeRepo
	accRepo           port.AccountRepo
	configurationRepo port.ConfigurationRepo
}

func NewLineService(qrRepo port.QRCodeRepo, accRepo port.AccountRepo, configurationRepo port.ConfigurationRepo) domain.LineService {
	return &lineService{
		qrRepo:            qrRepo,
		accRepo:           accRepo,
		configurationRepo: configurationRepo,
	}
}

func (s lineService) CreateQR(id int64, uid string) error {
	expireTime := time.Now().Add(5 * time.Minute)

	account, err := s.accRepo.Get(model.LINE_ID, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(errs.LW_NOT_FOUND_USER_DESC)
		}
		return errors.New(errs.LW_SYSTEM_ERROR_DESC)
	}

	open, err := s.configurationRepo.GetByKey("Open")
	if err != nil {
		return errors.New(errs.LW_CANT_CHECK_OC_TIME_DESC)
	}

	close, err := s.configurationRepo.GetByKey("Close")
	if err != nil {
		return errors.New(errs.LW_CANT_CHECK_OC_TIME_DESC)
	}

	if !account.IsActive {
		return errors.New(errs.LW_NOT_ACTIVE_USER_DESC)
	}

	openTime, err1 := time.Parse("15:04", open.Value)
	closeTime, err2 := time.Parse("15:04", close.Value)

	// Check for parsing errors.
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing time:", err1, err2)
		return errors.New(errs.LW_CANT_CHECK_OC_TIME_DESC)
	}

	currentTime := time.Now().UTC().Add(time.Hour * 7)

	todayOpenTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), openTime.Hour(), openTime.Minute(), 0, 0, currentTime.Location())
	todayCloseTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), closeTime.Hour(), closeTime.Minute(), 0, 0, currentTime.Location())

	enable, err := s.configurationRepo.GetByKey("Enable")
	if err != nil {
		return errors.New(errs.LW_SYSTEM_ERROR_DESC)
	}

	if enable.Value != "true" {
		return errors.New(errs.LW_SYSTEM_OFF_DESC)
	}

	if !todayOpenTime.Before(currentTime) || !todayCloseTime.After(currentTime) {
		return fmt.Errorf(errs.LW_OUT_OF_TIME_DESC, openTime.Format("15:04"), closeTime.Format("15:04"))
	}

	qrcode := model.QRCode{
		ID:        id,
		AccountID: account.ID,
		ExpireAt:  &expireTime,
	}

	if err = s.qrRepo.Create(&qrcode); err != nil {
		return errors.New(errs.LW_SYSTEM_ERROR_DESC)
	}
	return nil
}
