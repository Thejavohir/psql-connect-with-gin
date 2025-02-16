package models

type UserPKey struct {
	ID string `json:"id"`
	Username string `json:"username"`
}

type CreateUser struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type User struct {
	ID           string    `json:"id"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}

type UpdateUser struct {
	ID         string   `json:"id"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type UserGetListReq struct {
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Search  string `json:"search"`
}

type UserGetListResp struct {
	Count   int        `json:"count"`
	Users []*User `json:"users"`
}
