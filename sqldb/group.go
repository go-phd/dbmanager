package sqldb

import (
	"fmt"
	"time"

	"github.com/go-phd/ssf"
	"github.com/sirupsen/logrus"
)

// CshareGroup 共享群组描述符
type CshareGroup struct {
}

// CshareGroupTable 共享群组表结构
type CshareGroupTable struct {
	name       string
	reserveint int64
	reservestr string
}

// ShareGroup CshareGroup全局对象
var ShareGroup *CshareGroup

func initGroupTable() error {
	cg := CshareGroup{}

	sqlStmt := `
	create table if not exists ShareGroup (name text not null primary key, 
		reserveint integer,
		reservestr text,
		ts timestamp not null default(datetime('now','localtime')));
	`
	_, err := SQLDB.Exec(sqlStmt)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("create failed.")
		return err
	}
	/*
		cgt := CshareGroupTable{
			name:       "new1",
			reserveint: 666,
			reservestr: "66666",
		}*/

	//_ = cg.Insert(cgt)

	//time.Sleep(time.Duration(5) * time.Second)

	//cgt.reserveint = 000
	//cgt.reservestr = "000"
	//_ = cg.Update(cgt)

	//_ = cg.Delete(cgt)

	//cgts, _ := cg.QueueAll()
	cgts, _ := cg.Queue("new2")
	ssf.Logger.Warningln(cgts)

	ShareGroup = &cg

	ssf.Logger.Warningln("share group init success.")

	return nil
}

// QueueAll 查询全部
func (cg *CshareGroup) QueueAll() ([]CshareGroupTable, error) {
	sqlStmt := fmt.Sprintf(`select * from ShareGroup`)
	rows, err := SQLDB.Query(sqlStmt)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("select failed.")
	}
	defer rows.Close()

	shareGroupTable := make([]CshareGroupTable, 0)

	for rows.Next() {
		var name string
		var reserveint int64
		var reservestr string
		var ts time.Time

		err = rows.Scan(&name, &reserveint, &reservestr, &ts)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("rows scan failed.")
			break
		}

		//ssf.Logger.Warningln(ts)

		sg := CshareGroupTable{
			name:       name,
			reserveint: reserveint,
			reservestr: reservestr,
		}

		shareGroupTable = append(shareGroupTable, sg)
	}

	err = rows.Err()
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("rows failed.")
	}

	return shareGroupTable, err
}

// Queue 查询一行数据
func (cg *CshareGroup) Queue(n string) (*CshareGroupTable, error) {
	stmt, err := SQLDB.Prepare("select * from ShareGroup where name = ?")
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("select failed.")
		return nil, err
	}
	defer stmt.Close()

	var name string
	var reserveint int64
	var reservestr string
	var ts time.Time

	err = stmt.QueryRow(n).Scan(&name, &reserveint, &reservestr, &ts)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("rows scan failed.")
	}

	sg := CshareGroupTable{
		name:       name,
		reserveint: reserveint,
		reservestr: reservestr,
	}

	return &sg, err

	/*

		rows, err := SQLDB.Query(sqlStmt)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("select failed.")
		}
		defer rows.Close()

		var sg CshareGroupTable
		for rows.Next() {
			var name string
			var reserveint int64
			var reservestr string
			var ts time.Time

			err = rows.Scan(&name, &reserveint, &reservestr, &ts)
			if err != nil {
				ssf.Logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorf("rows scan failed.")
				break
			}

			//ssf.Logger.Warningln(ts)

			sg = CshareGroupTable{
				name:       name,
				reserveint: reserveint,
				reservestr: reservestr,
			}

			break
		}

		err = rows.Err()
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("rows failed.")
		}

		return &sg, err
	*/
}

// Insert 插入一行数据
func (cg *CshareGroup) Insert(data CshareGroupTable) error {
	sqlStmt := fmt.Sprintf(`insert into ShareGroup(name, reserveint, reservestr) values('%s', %d, '%s')`,
		data.name,
		data.reserveint,
		data.reservestr)
	_, err := SQLDB.Exec(sqlStmt)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("insert failed.")
	}

	return err
}

// Update 更新一行数据
func (cg *CshareGroup) Update(data CshareGroupTable) error {
	sqlStmt := fmt.Sprintf(`update ShareGroup set reserveint = %d, reservestr = '%s' where name = '%s'`,
		data.reserveint,
		data.reservestr,
		data.name)
	_, err := SQLDB.Exec(sqlStmt)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("update failed.")
	}

	return err
}

// Delete 删除一行数据
func (cg *CshareGroup) Delete(data CshareGroupTable) error {
	sqlStmt := fmt.Sprintf(`delete from ShareGroup where name = '%s'`,
		data.name)
	_, err := SQLDB.Exec(sqlStmt)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("delete failed.")
	}

	return err
}
