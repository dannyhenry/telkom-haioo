package produk

import (
	"context"
	"database/sql"
	"log"
	mp "telkom-haioo/domain/model/produk"
	"telkom-haioo/infra"
	"time"

	"github.com/sirupsen/logrus"
)

type ProdukConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newProduk(db *infra.DatabaseList, logger *logrus.Logger) ProdukConfig {
	return ProdukConfig{
		db:  db,
		log: logger,
	}
}

type Produk interface {
	GetListProduk(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseProduk, error)
	AddProduk(ctx context.Context, tx *sql.Tx, data mp.ParamsProduk) error
	IsExistProduk(ctx context.Context, data mp.ParamsProduk) (bool, error)
	GetListCart(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseCart, error)
	GetListCartByProdukID(ctx context.Context, produkID string) (int64, error)
	DeleteProduk(ctx context.Context, produkID string) error
	UpdateProduk(ctx context.Context, tx *sql.Tx, data mp.ParamsProduk) error
}

const (
	pcGetListProduk = `select product_id as kode_produk, product_name as nama_produk, price as harga FROM products OFFSET ((?-1)*?) rows FETCH NEXT ? ROWS ONLY `

	pcAddToCart = `INSERT INTO public.cart
	(product_id, quantity, created_at, updated_at)
	VALUES(?, ?, ?, ?);
	`

	// pcSelectExistProduk = `select exists (select c.product_id  from cart c where c.product_id in('?'))`

	pcUpdateCart = `UPDATE public.cart
	SET quantity=?, updated_at=?
	WHERE product_id=?;
	`

	pcListCart = `SELECT product_id as kode_produk, (select product_name from products p where p.product_id = c.product_id) as nama_produk,
		quantity FROM cart c OFFSET ((?-1)*?) rows FETCH NEXT ? ROWS ONLY
	`

	pcListCartWithNameProduk = `SELECT product_id as kode_produk, (select product_name from products p where p.product_id = c.product_id) as nama_produk,
	quantity FROM cart c where c.product_id in(select p2.product_id from products p2 where p2.product_name ilike '%' || ? || '%') OFFSET ((?-1)*?) rows FETCH NEXT ? ROWS ONLY
	`
)

func (pc ProdukConfig) GetListProduk(ctx context.Context, params mp.ParamsProduk) ([]mp.ResponseProduk, error) {
	var res []mp.ResponseProduk

	query, args, err := pc.db.Backend.Read.In(pcGetListProduk, params.Offset.Int64, params.Limit.Int64, params.Limit.Int64)
	if err != nil {
		return nil, err
	}

	query = pc.db.Backend.Read.Rebind(query)
	err = pc.db.Backend.Read.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return res, nil
}

func (pc ProdukConfig) AddProduk(ctx context.Context, tx *sql.Tx, data mp.ParamsProduk) error {
	param := make([]interface{}, 0)

	param = append(param, data.KodeProduk)
	param = append(param, data.Quantity)
	param = append(param, time.Now().UTC())
	param = append(param, time.Now().UTC())

	query, args, err := pc.db.Backend.Write.In(pcAddToCart, param...)
	if err != nil {
		return err
	}

	query = pc.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = pc.db.Backend.Write.QueryRow(ctx, query, args...)
	} else {
		res = tx.QueryRowContext(ctx, query, args...)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = res.Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil

}

func (pc ProdukConfig) IsExistProduk(ctx context.Context, data mp.ParamsProduk) (bool, error) {
	var result bool

	q := `select exists (select c.product_id  from cart c where c.product_id in('` + data.KodeProduk + `'))`

	query, args, err := pc.db.Backend.Read.In(q)
	if err != nil {
		return result, err
	}

	query = pc.db.Backend.Read.Rebind(query)
	err = pc.db.Backend.Read.GetContext(ctx, &result, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}

func (pc ProdukConfig) GetListCart(ctx context.Context, param mp.ParamsProduk) ([]mp.ResponseCart, error) {
	var res []mp.ResponseCart

	if param.Limit.Int64 != 0 && param.Offset.Int64 != 0 {

		query, args, err := pc.db.Backend.Read.In(pcListCart, param.Offset.Int64, param.Limit.Int64, param.Limit.Int64)
		if err != nil {
			return nil, err
		}

		query = pc.db.Backend.Read.Rebind(query)
		err = pc.db.Backend.Read.Select(&res, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

	} else if param.Limit.Int64 != 0 && param.Offset.Int64 != 0 && param.NamaProduk.String != "" {

		query, args, err := pc.db.Backend.Read.In(pcListCartWithNameProduk, param.NamaProduk.String, param.Offset.Int64, param.Limit.Int64, param.Limit.Int64)
		if err != nil {
			return nil, err
		}

		query = pc.db.Backend.Read.Rebind(query)
		err = pc.db.Backend.Read.Select(&res, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

	}

	return res, nil
}

func (pc ProdukConfig) GetListCartByProdukID(ctx context.Context, produkID string) (int64, error) {
	var res int64

	q := `SELECT quantity from cart c where c.product_id in('` + produkID + `')`

	query, args, err := pc.db.Backend.Read.In(q)
	if err != nil {
		return res, err
	}

	query = pc.db.Backend.Read.Rebind(query)
	err = pc.db.Backend.Read.GetContext(ctx, &res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	return res, nil
}

func (pc ProdukConfig) DeleteProduk(ctx context.Context, produkID string) error {

	pcDeleteCart := `DELETE FROM public.cart
	WHERE product_id = '` + produkID + `'
	`

	query, args, err := pc.db.Backend.Read.In(pcDeleteCart)
	if err != nil {
		return err
	}

	query = pc.db.Backend.Read.Rebind(query)

	query = pc.db.Backend.Read.Rebind(query)
	_, err = pc.db.Backend.Read.ExecContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil

}

func (pc ProdukConfig) UpdateProduk(ctx context.Context, tx *sql.Tx, data mp.ParamsProduk) error {
	param := make([]interface{}, 0)

	param = append(param, data.Quantity)
	param = append(param, time.Now().UTC())
	param = append(param, data.KodeProduk)

	query, args, err := pc.db.Backend.Write.In(pcUpdateCart, param...)
	if err != nil {
		return err
	}

	query = pc.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = pc.db.Backend.Write.QueryRow(ctx, query, args...)
	} else {
		res = tx.QueryRowContext(ctx, query, args...)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = res.Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
