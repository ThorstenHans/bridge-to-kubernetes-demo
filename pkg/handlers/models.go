package handlers

type CreateProductModel struct {
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	IsInStock bool    `json:"isInStock"`
}

type UpdateProductModel struct {
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	IsInStock bool    `json:"isInStock"`
}
