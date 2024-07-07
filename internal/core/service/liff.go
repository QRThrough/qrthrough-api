package service

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	errs "github.com/JMjirapat/qrthrough-api/pkg/errors"
	"gorm.io/gorm"
)

type liffService struct {
	accRepo    port.AccountRepo
	alumniRepo port.AlumniRepo
}

func NewLiffService(accRepo port.AccountRepo, alumniRepo port.AlumniRepo) domain.LiffService {
	return &liffService{
		accRepo:    accRepo,
		alumniRepo: alumniRepo,
	}
}

func (s liffService) GetAlumni(code int) (*dto.AlumniResponseBody, error) {
	user, _ := s.accRepo.Get(model.ACCOUNT_ID, code)
	if user != nil {
		return nil, errors.New(errs.LIFF_ALREADY_USED_ALUMNI_CODE)
	}

	result, err := s.alumniRepo.GetById(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			alumni := dto.AlumniResponseBody{
				InAlumni:    false,
				StudentCode: strconv.Itoa(code),
			}
			return &alumni, nil
		}
		return nil, errors.New("ไม่สามารถตรวจสอบรหัสนักศีกษาได้")
	}

	alumni := dto.AlumniResponseBody{
		InAlumni:    true,
		StudentCode: strconv.Itoa(code),
		Firstname:   &result.Firstname,
		Lastname:    &result.Lastname,
		Tel:         &result.Tel,
	}

	return &alumni, nil

}

func (s liffService) SignUp(body dto.RegisterRequestBody, lineID string) error {

	code, err := strconv.Atoi(body.StudentCode)
	if err != nil {
		log.Panic(err)
		return errors.New(errs.LIFF_ATOI_FAILED_CODE)
	}

	var flag model.Flag

	alumni, err := s.alumniRepo.GetById(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			flag = model.FLAG_NOTFOUND
		}
	} else {
		flag = model.FLAG_FOUND
		if body.Firstname != alumni.Firstname || body.Lastname != alumni.Lastname || body.Tel != alumni.Tel {
			flag = model.FLAG_EDIT
		}
	}

	account := model.Account{
		ID:        code,
		LineID:    lineID,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Tel:       body.Tel,
		Flag:      flag,
	}

	if err = s.accRepo.Create(&account); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_line_id_key\"") {
			return errors.New(errs.LIFF_DUPLICATE_LINE_CODE)
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_account_id_key\"") {
			return errors.New(errs.LIFF_DUPLICATE_STUDENT_ID_CODE)
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"accounts_tel_key\"") {
			return errors.New(errs.LIFF_DUPLICATE_TEL_CODE)
		}
		log.Panicf("%v", err)
		return errors.New(errs.LIFF_SIGN_UP_FAILED_CODE)
	}

	return nil
}
