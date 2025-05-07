package repo

import "context"

type AdminStorageI interface {
	Create(ctx context.Context, req *Admin) error
	GetByUName(ctx context.Context, username string) (*Admin, error)
	Update(ctx context.Context, id int, updates map[string]interface{}) error
	Delete(ctx context.Context, id int) error
}
type Admin struct {
	Id           uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"size:255;not null"`
	LastName     string `gorm:"size:255;not null"`
	Username     string `gorm:"size:255;not null;unique"`
	PasswordHash string `gorm:"size:255;not null"`
	Token        string `gorm:"size:255"`
}
