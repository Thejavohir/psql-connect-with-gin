package models

type ProductPKey struct {
	ID string `json:"id"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Price        float64       `json:"price"`
	CategoryData *Category `json:"category_data"`
	CategoryID   string    `json:"category_id"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}

type UpdateProduct struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type ProductGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ProductGetListResp struct {
	Count   int
	Product []*Product
}
