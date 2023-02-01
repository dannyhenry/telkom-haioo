package soal

import (
	"telkom-haioo/domain/model/general"

	"github.com/sirupsen/logrus"
)

type ServiceSoal struct {
	Soal SoalService
}

func NewSoalService(conf general.AppService, logger *logrus.Logger) ServiceSoal {
	return ServiceSoal{
		Soal: newSoalService(conf, logger),
	}
}
