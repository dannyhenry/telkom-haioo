package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	cg "telkom-haioo/domain/constants/general"
	"telkom-haioo/domain/model/general"
	mp "telkom-haioo/domain/model/produk"
	"telkom-haioo/domain/utils"
	sp "telkom-haioo/service/produk"

	"github.com/sirupsen/logrus"
)

type ProdukHandler struct {
	produk sp.ProdukService
	conf   general.AppService
	log    *logrus.Logger
}

func newProdukHandler(produk sp.ProdukService, conf general.AppService, logger *logrus.Logger) ProdukHandler {
	return ProdukHandler{
		produk: produk,
		conf:   conf,
		log:    logger,
	}
}

func (ph ProdukHandler) GetListProduk(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param mp.ParamsProduk

	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(req.URL.Query().Get("offset"))

	param.Limit.Int64 = int64(limit)
	param.Offset.Int64 = int64(offset)

	data, message, err := ph.produk.GetListProduk(req.Context(), param)
	if err != nil {
		respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status:  cg.Success,
		Message: message,
		Detail:  data,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return

}

func (ph ProdukHandler) GetListCart(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param mp.ParamsProduk

	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(req.URL.Query().Get("offset"))
	nama_produk := req.URL.Query().Get("nama_produk")

	param.Limit.Int64 = int64(limit)
	param.Offset.Int64 = int64(offset)
	param.NamaProduk.String = nama_produk

	data, message, err := ph.produk.GetListCart(req.Context(), param)
	if err != nil {
		respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status:  cg.Success,
		Message: message,
		Detail:  data,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ph ProdukHandler) DeleteProduk(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var params mp.ParamsProduk

	produkID := req.URL.Query().Get("produk_id")

	params.KodeProduk = produkID

	message, err := ph.produk.DeleteProduk(req.Context(), params)
	if err != nil {
		respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return

}

func (ph ProdukHandler) AddToCart(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var params mp.ParamsProduk

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &params)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := ph.produk.AddToCart(req.Context(), params)
	if err != nil {
		respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}
