package produk

import (
	"context"
	"errors"
	"telkom-haioo/domain/model/general"
	mp "telkom-haioo/domain/model/produk"
	"telkom-haioo/infra"
	rp "telkom-haioo/repository/produk"

	"github.com/sirupsen/logrus"
)

type ProdukService struct {
	dbTrx  rp.DatabaseProduk
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newProdukService(dbTrx rp.DatabaseProduk, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) ProdukService {
	return ProdukService{
		dbTrx:  dbTrx,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type Produk interface {
	AddToCart(ctx context.Context, param mp.ParamsProduk) (map[string]string, error)
	GetListCart(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseCart, map[string]string, error)
	DeleteProduk(ctx context.Context, param mp.ParamsProduk) (map[string]string, error)
	GetListProduk(ctx context.Context) ([]mp.ResponseProduk, map[string]string, error)
}

func (ps ProdukService) AddToCart(ctx context.Context, param mp.ParamsProduk) (map[string]string, error) {

	if param.KodeProduk == "" {
		return map[string]string{
			"en": "produk is required",
			"id": "produk is required",
		}, errors.New("produk is required")
	} else if param.Quantity == 0 {
		return map[string]string{
			"en": "quantity is required",
			"id": "quantity is required",
		}, errors.New("quantity is required")
	}

	isExist, err := ps.dbTrx.Produk.IsExistProduk(ctx, param)
	if err != nil {
		return map[string]string{
			"en": "failed to check is exist product",
			"id": "ada masalah saat mengecek produk di dalam keranjang",
		}, err
	}

	if isExist {
		getQuantity, err := ps.dbTrx.Produk.GetListCartByProdukID(ctx, param.KodeProduk)
		if err != nil {
			return map[string]string{
				"en": "failed to check is exist product",
				"id": "ada masalah saat mengecek produk di dalam keranjang",
			}, err
		}

		quantityNew := getQuantity + int64(param.Quantity)

		err = ps.dbTrx.Produk.UpdateProduk(ctx, nil, mp.ParamsProduk{
			Quantity:   int(quantityNew),
			KodeProduk: param.KodeProduk,
		})
		if err != nil {
			return map[string]string{
				"en": "failed add to cart",
				"id": "ada masalah saat menambahkan produk ke keranjang",
			}, err
		}
	} else {
		err = ps.dbTrx.Produk.AddProduk(ctx, nil, mp.ParamsProduk{
			KodeProduk: param.KodeProduk,
			Quantity:   param.Quantity,
		})
		if err != nil {
			return map[string]string{
				"en": "failed add to cart",
				"id": "ada masalah saat menambahkan produk ke keranjang",
			}, err
		}
	}

	return map[string]string{
		"en": "successfully add to cart",
		"id": "berhasil menambahkan produk",
	}, nil
}

func (ps ProdukService) GetListCart(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseCart, map[string]string, error) {

	if param.Limit.Int64 == 0 || param.Offset.Int64 == 0 {
		return nil, map[string]string{
			"en": "limit or offset required",
			"id": "limit of offset required",
		}, errors.New("error params")
	}

	dataListCart, err := ps.dbTrx.Produk.GetListCart(ctx, param)
	if err != nil {
		return nil, map[string]string{
			"en":  "failed to get list product in cart",
			"id":  "ada masalah saat menampilkan daftar produk di kerangjang",
			"err": err.Error(),
		}, err
	}

	return dataListCart, map[string]string{
		"en": "successfully get list product in cart",
		"id": "berhasil menampilkan daftar produk di keranjang",
	}, nil
}

func (ps ProdukService) DeleteProduk(ctx context.Context, param mp.ParamsProduk) (map[string]string, error) {
	if param.KodeProduk == "" {
		return map[string]string{
			"en": "param produk id required",
			"id": "param produk id required",
		}, errors.New("param required")
	}

	err := ps.dbTrx.Produk.DeleteProduk(ctx, param.KodeProduk)
	if err != nil {
		return map[string]string{
			"en":  "failed to delete product",
			"id":  "ada masalah saat menghapus produk",
			"err": err.Error(),
		}, err
	}

	return map[string]string{
		"en": "successfully deleted product",
		"id": "berhasil menghapus produk",
	}, nil
}

func (ps ProdukService) GetListProduk(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseProduk, map[string]string, error) {
	if param.Limit.Int64 == 0 || param.Offset.Int64 == 0 {
		return nil, map[string]string{
			"en": "param limit or offset is required",
			"id": "param limit or offset is required",
		}, errors.New("error")
	}

	dataProduk, err := ps.dbTrx.Produk.GetListProduk(ctx, param)
	if err != nil {
		return nil, map[string]string{
			"en":  "failed to get list product",
			"id":  "ada masalah saat menampilkan produk",
			"err": err.Error(),
		}, err
	}

	if dataProduk == nil {
		return nil, map[string]string{
			"en": "data not found",
			"id": "data produk tidak ditemukan",
		}, errors.New("data not found")
	}

	return dataProduk, map[string]string{
		"en": "successfully get list product",
		"id": "berhasil menampilkan daftar produk",
	}, nil
}
