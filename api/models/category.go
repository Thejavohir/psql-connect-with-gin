package models

type CategoryPKey struct {
	ID string `json:"id"`
}

type CreateCategory struct {
	Title    string `json:"title"`
	ParentID string `json:"parent_id"`
}

type Category struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCategory struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
}

type CategoryGetListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Search string `json:"search"`
}

type CategoryGetListResp struct {
	Count int
	Category []*Category
}