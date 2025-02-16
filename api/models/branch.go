package models

type BranchPKey struct {
	ID string `json:"id"`
}

type CreateBranch struct {
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber string   `json:"phone_number"`
	// ProductIDs  []string `json:"product_ids"`
}

type Branch struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Products    []*Product `json:"products"`
}

type UpdateBranch struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber string   `json:"phone_number"`
	Products    []string `json:"products"`
}

type BranchGetListReq struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BranchGetListResp struct {
	Count  int       `json:"count"`
	Branch []*Branch `json:"branch"`
}