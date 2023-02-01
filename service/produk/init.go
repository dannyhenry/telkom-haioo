package produk

import (
	"telkom-haioo/domain/model/general"
	"telkom-haioo/infra"
	"telkom-haioo/repository"

	"github.com/sirupsen/logrus"
)

type ServiceProduk struct {
	Produk ProdukService
}

func NewProdukService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) ServiceProduk {
	return ServiceProduk{
		Produk: newProdukService(repo.Produk, conf, dbList, logger),
	}
}
