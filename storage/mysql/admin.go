package mysql

import (
	"context"
	"fmt"

	"GitHub.com/sattorovshoxrux3009/SavdoPall_back/storage/repo"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminStorage(db *gorm.DB) repo.AdminStorageI {
	return &adminRepo{db: db}
}

func (a *adminRepo) Create(ctx context.Context, req *repo.Admin) error {
	if err := a.db.WithContext(ctx).Create(req).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) GetByUName(ctx context.Context, username string) (*repo.Admin, error) {
	var admin repo.Admin
	if err := a.db.WithContext(ctx).Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (a *adminRepo) Update(ctx context.Context, id int, updates map[string]interface{}) error {
	var admin repo.Admin
	if err := a.db.WithContext(ctx).First(&admin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("admin with id %d not found", id)
		}
		return err
	}
	if err := a.db.WithContext(ctx).Model(&admin).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) Delete(ctx context.Context, id int) error {
	var admin repo.Admin
	if err := a.db.WithContext(ctx).First(&admin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("admin with id %d not found", id)
		}
		return err
	}
	if err := a.db.WithContext(ctx).Delete(&admin).Error; err != nil {
		return err
	}
	return nil
}
