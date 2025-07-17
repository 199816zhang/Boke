package flags

import (
	"blogx_server/global"
	"blogx_server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		logrus.Fatalf("auto migrate err: %v", err)
		return
	}
	logrus.Info("auto migrate success")
}
