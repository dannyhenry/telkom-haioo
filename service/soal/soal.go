package soal

import (
	"context"
	"fmt"
	"math"
	"telkom-haioo/domain/model/general"
	ms "telkom-haioo/domain/model/soal"

	"github.com/sirupsen/logrus"
)

type SoalService struct {
	conf general.AppService
	log  *logrus.Logger
}

func newSoalService(conf general.AppService, logger *logrus.Logger) SoalService {
	return SoalService{
		conf: conf,
		log:  logger,
	}
}

type Soal interface {
	Soal1FractionsMoney(ctx context.Context, money int) interface{}
	Soal2CompareTwoString(ctx context.Context, text1, text2 string) ms.ResponseCompareString
}

func (ss SoalService) Soal1FractionsMoney(ctx context.Context, money int) interface{} {
	moneyArr := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 50}
	result := make(map[string]int)
	var res []interface{}
	for _, m := range moneyArr {
		if money >= m {
			count := int(math.Floor(float64(money) / float64(m)))
			money = money % m
			result[fmt.Sprintf("Rp. %d", m)] = count
		}
	}

	for k, v := range result {
		res = append(res, map[string]interface{}{k: v})
	}

	var resultNew ms.ResponseMoney

	resultNew.Result = append(resultNew.Result, res)

	return resultNew
}

func (ss SoalService) Soal2CompareTwoString(ctx context.Context, text1, text2 string) ms.ResponseCompareString {
	var result ms.ResponseCompareString

	if math.Abs(float64(len(text1)-len(text2))) > 1 {
		result.Result = false
		return result
	}

	var i, j, edits int
	for i < len(text1) && j < len(text2) {
		if text1[i] != text2[j] {
			edits++
			if len(text1) > len(text2) {
				i++
			} else if len(text1) < len(text2) {
				j++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}

		if edits > 1 {
			result.Result = false
			return result
		}
	}

	result.Result = true
	return result
}
