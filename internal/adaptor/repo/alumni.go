package repo

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/port"
	"gorm.io/gorm"
)

type alumniRepo struct {
	db *gorm.DB
}

func NewAlumniRepo(db *gorm.DB) port.AlumniRepo {
	return &alumniRepo{
		db: db,
	}
}

func (r alumniRepo) Create(alumni *model.Alumni) error {
	return r.db.Create(alumni).Error
}

func (r alumniRepo) GetById(id int) (*model.Alumni, error) {
	var alumni model.Alumni
	if err := r.db.Take(&alumni, "student_code=?", id).Error; err != nil {
		return nil, err
	}

	return &alumni, nil
}

func (r alumniRepo) UpdateById(id int, alumni model.Alumni) error {
	return r.db.Where("student_code=?", id).Omit("ID").Updates(&alumni).Error
}
