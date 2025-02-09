package storage

import (
	"context"
	"psql/api/models"
)

type StorageI interface{
	Close()
	Category() CategoryRepoI
	Product() ProductRepoI
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