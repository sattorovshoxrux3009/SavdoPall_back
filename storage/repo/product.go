package repo

import "context"

type ProductStorageI interface {
	Create(ctx context.Context, req *Product) (*Product, error)
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
	Left        int     `gorm:"not null"`
	Description string  `gorm:"size:1000"`
}

// {
//     id: 0,
//     img:  'https://images.uzum.uz/cp3o4bnfrr80f2gllh0g/original.jpg',
//     name: "Televizor Roison 32",
//     price:1454000,
//     height:55,
//     width: 70,
//     depth: 12,
//     quantity: 0,
//     left: 4,
//     description: "Televizor Roison Smart LED HD TV RE 32-060,43-430 BL, Аndroid 12, ovozli pulti bilan. Dasturiy ta'minot: Netflix, Youtube, Google Play"
// },
