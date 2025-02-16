package redis

import (
	"context"
	"encoding/json"
	"psql/api/models"

	"github.com/redis/go-redis/v9"
)

const PRODUCT = "products"

type productRepo struct {
	rdb *redis.Client
}

func NewProductRepo(rdb *redis.Client) *productRepo {
	return &productRepo{
		rdb: rdb,
	}
}

func (c *productRepo) Exists(ctx context.Context) (bool, error) {

	rowsAffected, err := c.rdb.Exists(ctx, "products").Result()
	if err != nil {
		return false, nil
	}

	if rowsAffected <= 0 {
		return false, nil
	}
	return true, nil

}

func (c *productRepo) CreateGetList(ctx context.Context, req *models.ProductGetListResp) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.rdb.Set(ctx, PRODUCT, body, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *productRepo) GetList(ctx context.Context) (*models.ProductGetListResp, error) {
	var response *models.ProductGetListResp

	data, err := c.rdb.Get(ctx, PRODUCT).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
