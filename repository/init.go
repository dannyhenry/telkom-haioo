package repository

import (
	"telkom-haioo/infra"
	rp "telkom-haioo/repository/produk"

	"github.com/sirupsen/logrus"
)

type Repo struct {
	Produk rp.DatabaseProduk
}

func NewRepo(database *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		Produk: rp.NewDatabaseProduk(database, logger),
	}
}
