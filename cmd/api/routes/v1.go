package routes

import (
	"net/http"

	"telkom-haioo/domain/model/general"
	"telkom-haioo/handler/api"

	"github.com/gorilla/mux"
)

func getV1(router, routerJWT *mux.Router, conf *general.AppService, handler api.Handler) {
	router.HandleFunc("/v1/fractionsmoney", handler.Soal.GetFractionsMoney).Methods(http.MethodPost)
	router.HandleFunc("/v1/comparestring", handler.Soal.GetCompareTwoString).Methods(http.MethodPost)
	router.HandleFunc("/v1/products", handler.Produk.GetListProduk).Methods(http.MethodGet)
	router.HandleFunc("/v1/carts", handler.Produk.GetListCart).Methods(http.MethodGet)
	router.HandleFunc("/v1/cart", handler.Produk.AddToCart).Methods(http.MethodPost)
	router.HandleFunc("/v1/cart", handler.Produk.DeleteProduk).Methods(http.MethodDelete)
}
