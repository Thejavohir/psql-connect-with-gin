package models

type ProductPKey struct {
	ID string `json:"id"`
}

type CreateProduct struct {
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	CategoryID string   `json:"category_id"`
	BranchIDs  []string `json:"branch_ids"`
	Barcode    string   `json:"barcode"`
}

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	CategoryID   string    `json:"category_id"`
	Barcode      string    `json:"barcode"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	CategoryData *Category `json:"category_data"`
	Branches     []*Branch `json:"branches" swaggerignore:"true"`
}

type UpdateProduct struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	CategoryID string   `json:"category_id"`
	Barcode    string   `json:"barcode"`
	BranchIDs  []string `json:"branch_ids"`
}

type ProductGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
	Barcode string `json:"barcode"`
}

type ProductGetListResp struct {
	Count   int        `json:"count"`
	Product []*Product `json:"product"`
}
