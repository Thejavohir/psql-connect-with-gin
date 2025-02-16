package models

type Remainder struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	BranchID   string  `json:"branch_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type RemainderPKey struct {
	ID string `json:"id"`
}

type CreateRemainder struct {
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	BranchID   string  `json:"branch_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

type UpdateRemainder struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	BranchID   string  `json:"branch_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

type RemainderrGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type RemainderrGetListResp struct {
	Count      int          `json:"count"`
	Remainders []*Remainder `json:"remianders"`
}
