package mysql

import (
	"context"
	"errors"
	"fmt"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) repo.ProductStorageI {
	return &productRepo{db: db}
}
func (p *productRepo) Create(ctx context.Context, req *repo.Product) (*repo.Product, error) {
	if err := p.db.WithContext(ctx).Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}
func (p *productRepo) Get(ctx context.Context) (*[]repo.Product, error) {
	var products []repo.Product
	if err := p.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (p *productRepo) GetById(ctx context.Context, id int) (*repo.Product, error) {
	var product repo.Product
	if err := p.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (p *productRepo) Update(ctx context.Context, id int, updates map[string]interface{}) error {
	var product repo.Product
	if err := p.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found")
		}
		return err
	}

	if err := p.db.WithContext(ctx).Model(&product).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepo) Delete(ctx context.Context, id int) error {
	// Oldin mahsulot borligini tekshiramiz
	var product repo.Product
	if err := p.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found")
		}
		return err
	}

	// O‘chiramiz
	if err := p.db.WithContext(ctx).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
