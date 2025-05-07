package storage

import (
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/mysql"
	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/gorm"
)

type StorageI interface {
	Product() repo.ProductStorageI
	Admin() repo.AdminStorageI
}
type storagePg struct {
	productRepo repo.ProductStorageI
	adminRepo   repo.AdminStorageI
}

func NewStorage(mysqlConn *gorm.DB) StorageI {
	return &storagePg{
		productRepo: mysql.NewProductStorage(mysqlConn),
		adminRepo:   mysql.NewAdminStorage(mysqlConn),
	}
}
func (s *storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
func (s *storagePg) Admin() repo.AdminStorageI {
	return s.adminRepo
}
