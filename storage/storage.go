package storage

import (
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/mysql"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/gorm"
)

type StorageI interface {
	Product() repo.ProductStorageI
}
type storagePg struct {
	productRepo repo.ProductStorageI
}

func NewStorage(mysqlConn *gorm.DB) StorageI {
	return &storagePg{
		productRepo: mysql.NewProductStorage(mysqlConn),
	}
}
func (s *storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
