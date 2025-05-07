package mysql

import (
	"context"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) repo.ProductStorageI {
	return &productRepo{db: db}
}
func (a *productRepo) Create(ctx context.Context, req *repo.Product) (*repo.Product, error) {
	if err := a.db.WithContext(ctx).Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}
