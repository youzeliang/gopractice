package helper

import (
	"github.com/youzeliang/gopractice/base"
	"github.com/youzeliang/gopractice/conf"
	"time"
	"github.com/jinzhu/gorm"

)


var MysqlClient *gorm.DB

func InitMysql() {
	var loglevel bool
	if conf.Conf.LogLevel == "debug" {
		loglevel = true
	} else {
		loglevel = false
	}

	var err error
	MysqlClient, err = base.InitMysqlClient(base.MysqlConf{
		User:            conf.Conf.Mysql.User,
		Password:        conf.Conf.Mysql.PassWord,
		Addr:            conf.Conf.Mysql.Addr,
		DataBase:        conf.Conf.Mysql.DataBase,
		MaxIdleConns:    conf.Conf.Mysql.MaxIdleConns,
		MaxOpenConns:    conf.Conf.Mysql.MaxOpenConns,
		ConnMaxLifeTime: 3600 * time.Second,
		LogMode:         loglevel,
	})
	if err != nil {
		base.PanicfLogger(nil, "mysql connect error: %v", err)
	}
}