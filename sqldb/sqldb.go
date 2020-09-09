package sqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-phd/ssf"
	"github.com/sirupsen/logrus"
)

// SQLDB 数据库驱动全局对象
var SQLDB *sql.DB

func init() {
	var err error

	dbPath := fmt.Sprintf("%s/phdConfig.db", ssf.SSFConfig.DB.Path)
	ssf.Logger.Warningln(dbPath)
	if ssf.SSFConfig.Sqlite.Enable {
		SQLDB, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Open db failed.")
		}
	} else {
		// 其他数据库类型
	}

	initGroupTable()
}
