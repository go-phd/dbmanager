package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/go-phd/ssf"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func init() {
	err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("RegisterDriver failed.")
		return
	}

	dbPath := fmt.Sprintf("%s/phdConfig.db", ssf.SSFConfig.DB.Path)
	err = orm.RegisterDataBase("default", "sqlite3", dbPath)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err":    err,
			"dbPath": dbPath,
		}).Errorln("RegisterDataBase failed.")
		return
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("RunSyncdb failed.")
		return
	}

	ssf.Logger.Infoln("main init success.")
}
