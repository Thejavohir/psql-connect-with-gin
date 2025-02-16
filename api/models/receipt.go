package models

type Receipt struct {
	ID        string `json:"id"`
	ReceiptID string `json:"receipt_id"`
	BranchID  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReceiptPKey struct {
	ID string `json:"id"`
}

type CreateReceipt struct {
	ReceiptID string `json:"receipt_id"`
	BranchID  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
}

type UpdateReceipt struct {
	ID        string `json:"id"`
	ReceiptID string `json:"receipt_id"`
	BranchID  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
}

type ReceiptGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ReceiptGetListResp struct {
	Count   int        `json:"count"`
	Receipts []*Receipt `json:"receipts"`
}
