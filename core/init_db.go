package core

import (
	"blogx_server/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func InitDB() *gorm.DB {
	dc := global.Config.DB
	dc1 := global.Config.DB1
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败")
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("获取底层数据库错误")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功！")
	if !dc1.Empty() {
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc.DSN())},
			Replicas: []gorm.Dialector{mysql.Open(dc1.DSN())},
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			logrus.Fatalf("读写分离配置错误：%v", err.Error())
			return nil
		}
	}
	return db
}
