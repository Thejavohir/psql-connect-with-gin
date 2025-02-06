package storage

import "psql/models"

type StorageI interface{
	Close()
	Category() CategoryRepoI
	Product() ProductRepoI
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetById(*models.CategoryPKey) (*models.Category, error)
	GetList(*models.CategoryGetListReq) (*models.CategoryGetListResp, error)
	Update(*models.UpdateCategory) (*models.Category, error)
	Delete(*models.CategoryPKey) error 
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetById(*models.ProductPKey) (*models.Product, error)
	GetList(*models.ProductGetListReq) (*models.ProductGetListResp, error)
	Update(*models.UpdateProduct) (*models.Product, error)
	Delete(*models.ProductPKey) error 
}