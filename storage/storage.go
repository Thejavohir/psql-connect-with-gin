package storage

import (
	"context"
	"psql/api/models"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Category() CategoryRepoI
	Product() ProductRepoI
	Branch() BranchRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetById(context.Context, *models.CategoryPKey) (*models.Category, error)
	GetList(context.Context, *models.CategoryGetListReq) (*models.CategoryGetListResp, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CategoryPKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetById(context.Context, *models.ProductPKey) (*models.Product, error)
	GetList(context.Context, *models.ProductGetListReq) (*models.ProductGetListResp, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPKey) error
}
type BranchRepoI interface {
	Create(context.Context, *models.CreateBranch) (string, error)
	GetById(context.Context, *models.BranchPKey) (*models.Branch, error)
	GetList(context.Context, *models.BranchGetListReq) (*models.BranchGetListResp, error)
	Update(context.Context, *models.UpdateBranch) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.BranchPKey) error
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetById(context.Context, *models.UserPKey) (*models.User, error)
	GetByUsername(context.Context, *models.UserPKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListReq) (*models.UserGetListResp, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.UserPKey) error
}
