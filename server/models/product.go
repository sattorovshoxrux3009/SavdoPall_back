package models

type CreateProduct struct {
	ImgUrl      string  `json:"img_url" validate:"required,url"`
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Height      float64 `json:"height" validate:"required,gt=0"`
	Width       float64 `json:"width" validate:"required,gt=0"`
	Depth       float64 `json:"depth" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gte=0"`
	Left        int     `json:"left" validate:"required,gte=0"`
	Description string  `json:"description" validate:"max=1000"`
}
