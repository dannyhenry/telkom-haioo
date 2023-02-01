package api

import (
	"telkom-haioo/domain/model/general"
	"telkom-haioo/handler/api/authorization"
	"telkom-haioo/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Token  authorization.TokenHandler
	Public authorization.PublicHandler
	Produk ProdukHandler
	Soal   SoalHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		Token:  authorization.NewTokenHandler(conf, logger),
		Public: authorization.NewPublicHandler(conf, logger),
		Produk: newProdukHandler(sv.Produk.Produk, conf, logger),
		Soal:   newSoalHandler(sv.Soal.Soal, conf, logger),
	}
}
