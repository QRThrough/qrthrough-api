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

type scannerService struct {
	qrRepo            port.QRCodeRepo
	logRepo           port.UsageRepo
	scannerRepo       port.ScannerRepo
	configurationRepo port.ConfigurationRepo
}

func NewScannerService(qrRepo port.QRCodeRepo, logRepo port.UsageRepo, scannerRepo port.ScannerRepo, configurationRepo port.ConfigurationRepo) domain.ScannerService {
	return &scannerService{
		qrRepo:            qrRepo,
		logRepo:           logRepo,
		scannerRepo:       scannerRepo,
		configurationRepo: configurationRepo,
	}
}

func (s scannerService) VerifyQR(id int64) error {
	open, err := s.configurationRepo.GetByKey("Open")
	if err != nil {
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}

	close, err := s.configurationRepo.GetByKey("Close")
	if err != nil {
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}

	openTime, err1 := time.Parse("15:04", open.Value)
	closeTime, err2 := time.Parse("15:04", close.Value)

	// Check for parsing errors.
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing time:", err1, err2)
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}

	currentTime := time.Now().UTC().Add(time.Hour * 7)

	todayOpenTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), openTime.Hour(), openTime.Minute(), 0, 0, currentTime.Location())
	todayCloseTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), closeTime.Hour(), closeTime.Minute(), 0, 0, currentTime.Location())

	enable, err := s.configurationRepo.GetByKey("Enable")
	if err != nil {
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}

	if enable.Value != "true" {
		return errors.New(errs.SC_SYSTEM_OFF_CODE)
	}

	if !todayOpenTime.Before(currentTime) || !todayCloseTime.After(currentTime) {
		return errors.New(errs.SC_OUT_OF_TIME_CODE)
	}

	qrcode, err := s.qrRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(errs.SC_QRCODE_NOT_FOUND_CODE)
		}
		return err
	}

	if (*qrcode.ExpireAt).Before(currentTime) {
		return errors.New(errs.SC_QRCODE_EXPIRED_CODE)
	}

	count, err := s.logRepo.CountByQRCode(qrcode.ID)
	if err != nil {
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}

	if count >= 5 {
		return errors.New(errs.SC_QRCODE_USED_UP_CODE)
	}

	usage := model.Usage{
		AccountID: qrcode.AccountID,
		QRCodeID:  qrcode.ID,
	}
	if err = s.logRepo.Create(&usage); err != nil {
		return errors.New(errs.SYSTEM_ERROR_CODE)
	}
	return nil
}

func (s scannerService) CheckExistedScanner(mac string) bool {
	return s.scannerRepo.CheckExistedScanner(mac)
}
