package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cg "telkom-haioo/domain/constants/general"
	"telkom-haioo/domain/model/general"
	ms "telkom-haioo/domain/model/soal"
	"telkom-haioo/domain/utils"
	sp "telkom-haioo/service/soal"

	"github.com/sirupsen/logrus"
)

type SoalHandler struct {
	Soal sp.SoalService
	conf general.AppService
	log  *logrus.Logger
}

func newSoalHandler(Soal sp.SoalService, conf general.AppService, logger *logrus.Logger) SoalHandler {
	return SoalHandler{
		Soal: Soal,
		conf: conf,
		log:  logger,
	}
}

func (ch SoalHandler) GetFractionsMoney(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param ms.GetFractionsMoney

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	data := ch.Soal.Soal1FractionsMoney(req.Context(), param.InputMoney)
	if err != nil {
		// respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status: cg.Success,
		// Message: message,
		Detail: data,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}

func (ch SoalHandler) GetCompareTwoString(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param ms.GetCompareTwoString

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	data := ch.Soal.Soal2CompareTwoString(req.Context(), param.Kalimat1, param.Kalimat2)
	if err != nil {
		// respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status: cg.Success,
		// Message: message,
		Detail: data,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}
