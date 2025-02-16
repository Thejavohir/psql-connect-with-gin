package postgres

import (
	"context"
	"psql/api/models"
	"testing"
)

func TestProduct(t *testing.T) {

	tests := []struct{
		Name string
		Input *models.CreateProduct
		Output string
		WantErr bool
	}{
		{
			Name: "Product Case 1",
			Input: &models.CreateProduct{
				Name: "Alpomish",
				Price: 80000,
			},
			WantErr: true,
		},
		{
			Name: "Product Case 2",
			Input: &models.CreateProduct{
				Name: "iPhone 14 Pro Max",
				Price: 7000000,
				CategoryID: "73d61a2e-0fe7-4901-8384-a7fbe459fb78",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := productTestRepo.Create(context.Background(), test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got %+v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %+v", test.Name, err)
				return
			}
		})
	}
}