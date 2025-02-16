package storage

import (
	"context"
	"psql/api/models"
)

type CacheI interface {
	Close()
	Product() ProductRepoCacheI
}

type ProductRepoCacheI interface{
	CreateGetList(context.Context, *models.ProductGetListResp) error
	GetList(context.Context) (*models.ProductGetListResp, error) 
	Exists(context.Context) (bool, error)
}