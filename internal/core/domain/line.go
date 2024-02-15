package domain

type LineService interface {
	CreateQR(id int64, uid string) error
}
