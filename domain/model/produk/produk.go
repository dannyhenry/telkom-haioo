package produk

import "gopkg.in/guregu/null.v4"

type ParamsProduk struct {
	KodeProduk string `json:"kode_produk"`
	Quantity   int    `json:"quantity"`
	Offset     null.Int
	Limit      null.Int
	NamaProduk null.String
}

type ResponseCart struct {
	KodeProduk string `json:"kode_produk" db:"kode_produk"`
	NamaProduk string `json:"nama_produk" db:"nama_produk"`
	Quantity   int    `json:"quantity" db:"quantity"`
}

type ResponseProduk struct {
	KodeProduk string `json:"kode_produk" db:"kode_produk"`
	NamaProduk string `json:"nama_produk" db:"nama_produk"`
	Harga      int    `json:"harga" db:"harga"`
}
