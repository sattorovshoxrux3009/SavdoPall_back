package repo

import "context"

type ProductStorageI interface {
	Create(ctx context.Context, req *Product) (*Product, error)
	Get(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
}
type Product struct {
	Id          uint    `gorm:"primaryKey"`
	ImgUrl      string  `gorm:"size:255;not null"`
	Name        string  `gorm:"size:255;not null"`
	Price       float64 `gorm:"not null"`
	Height      float64 `gorm:"not null"`
	Width       float64 `gorm:"not null"`
	Depth       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	Description string  `gorm:"size:1000"`
}
