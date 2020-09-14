package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/go-phd/dbmanager/models"
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

	// 注册数据库同步 notify， 在有些操作时广播，每个微服务实例都会收到
	err = ssf.MQ.RegisterNotifyMethod(models.DbSyncNotify, dbSyncnotifyCB, false)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("RegisterNotifyMethod failed.")
		return
	}

	// 注册 debug cast
	err = ssf.MQ.RegisterCastMethod("debug", debugCB)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("RegisterCastMethod failed.")
	}

	ssf.Logger.Infoln("main init success.")
}

// dbSyncnotifyCB 数据库同步广播回调
func dbSyncnotifyCB(arg interface{}) interface{} {
	ds := &models.DbSyncDes{}

	data := arg.([]byte)
	err := ssf.Serializer.Unmarshal(data, ds)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Unmarshal failed.")
		return nil
	}

	if ds.UUID == ssf.Ssf.GetUUID() {
		ssf.Logger.WithFields(logrus.Fields{
			"uuid": ds.UUID,
		}).Warningln("Ignore local operations.")
		return nil
	}

	switch ds.Table {
	case models.UserTable:
		var u models.User
		err := ssf.Serializer.Unmarshal(ds.Data, &u)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Unmarshal failed.")
			return nil
		}

		switch ds.Type {
		case models.InsertOp:
			_, err := models.AddUser(u)
			if err != nil {
				ssf.Logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("AddUser failed.")
			}
		case models.UpdateOp:
			err := models.UpdateUser(u.Username, u)
			if err != nil {
				ssf.Logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("UpdateUser failed.")
			}
		case models.DeleteOp:
			err := models.DeleteUser(u.Username)
			if err != nil {
				ssf.Logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("DeleteUser failed.")
			}
		default:
			ssf.Logger.WithFields(logrus.Fields{
				"type": ds.Type,
			}).Errorln("unknown type.")
		}

	default:
		ssf.Logger.WithFields(logrus.Fields{
			"table": ds.Table,
		}).Errorln("unknown table.")
	}

	return nil
}

// dbSyncnotifyCB 数据库同步广播回调
func debugCB(arg interface{}) interface{} {

	return nil
}
