package produk

import (
	"context"
	mp "telkom-haioo/domain/model/produk"
	"testing"

	"gopkg.in/guregu/null.v4"
)

func TestGetListProduk(t *testing.T) {
	ps := ProdukService{}
	ctx := context.Background()

	var (
		limit, offset null.Int
	)

	limit.Int64 = 1
	offset.Int64 = 1

	param := mp.ParamsProduk{Limit: limit, Offset: offset}

	t.Run("Positive Test Case: Limit and Offset are provided", func(t *testing.T) {
		dataProduk, msg, err := ps.GetListProduk(ctx, param)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if dataProduk == nil {
			t.Errorf("Error: %v", "data not found")
		}
		if msg["en"] != "successfully get list product" {
			t.Errorf("Error: %v", "unexpected error message")
		}
	})

	t.Run("Negative Test Case: Limit or Offset is not provided", func(t *testing.T) {
		limit.Int64 = 0
		offset.Int64 = 0

		param := mp.ParamsProduk{Limit: limit, Offset: offset}
		_, msg, err := ps.GetListProduk(ctx, param)
		if err == nil {
			t.Errorf("Error: %v", "error expected")
		}
		if msg["en"] != "param limit or offset is required" {
			t.Errorf("Error: %v", "unexpected error message")
		}
	})
}
