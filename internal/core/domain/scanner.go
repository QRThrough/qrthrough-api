package domain

type ScannerService interface {
	VerifyQR(id int64) error
	CheckExistedScanner(mac string) bool
}
