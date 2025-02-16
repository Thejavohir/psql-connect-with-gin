package models

type ReceiptProduct struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	ReceiptID  string  `json:"receipt_id"`
	CategoryID string  `json:"category_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type ReceiptProductPKey struct {
	ID string `json:"id"`
}

type CreateReceiptProduct struct {
	Name       string  `json:"name"`
	ReceiptID  string  `json:"receipt_id"`
	CategoryID string  `json:"category_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

type UpdateReceiptProduct struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	ReceiptID  string  `json:"receipt_id"`
	CategoryID string  `json:"category_id"`
	Barcode    string  `json:"barcode"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

type ReceiptPrGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ReceiptPrGetListResp struct {
	Count    int        `json:"count"`
	ReceiptPs []*ReceiptProduct `json:"receipt_ps"`
}
