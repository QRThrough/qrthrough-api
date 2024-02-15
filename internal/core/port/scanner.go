package port

type ScannerRepo interface {
	CheckExistedScanner(mac string) bool
}
