package service

import (
	"telkom-haioo/domain/model/general"
	"telkom-haioo/infra"
	"telkom-haioo/repository"

	sp "telkom-haioo/service/produk"
	ss "telkom-haioo/service/soal"

	"github.com/sirupsen/logrus"
)

type Service struct {
	Produk sp.ServiceProduk
	Soal   ss.ServiceSoal
}

func NewService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) Service {
	return Service{
		Produk: sp.NewProdukService(repo, conf, dbList, logger),
		Soal:   ss.NewSoalService(conf, logger),
	}
}
