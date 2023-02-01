package produk

import (
	"telkom-haioo/infra"

	"github.com/sirupsen/logrus"
)

type DatabaseProduk struct {
	Produk Produk
}

func NewDatabaseProduk(db *infra.DatabaseList, logger *logrus.Logger) DatabaseProduk {
	return DatabaseProduk{
		Produk: newProduk(db, logger),
	}
}
