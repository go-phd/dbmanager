package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/go-phd/ssf"
	"github.com/sirupsen/logrus"
)

// User 用户表
type User struct {
	Id       int64 //主键
	Username string
	Password string
}

// AddUser 添加用户
func AddUser(u User) (string, error) {

	// Username 不能相同
	_, err := GetUser(u.Username)
	if err == nil {
		errStr := "Username is already exist."
		ssf.Logger.WithFields(logrus.Fields{
			"Username": u.Username,
		}).Errorln(errStr)
		return "", errors.New(errStr)
	}

	// 主键不可设置，只能累加
	u.Id = 0
	ret, err := orm.NewOrm().Insert(&u)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("insert failed.")
		return "", errors.New("Failed to add a new user.")
	}

	// 发送 insert 广播
	data, err := ssf.Serializer.Marshal(u)
	if err == nil {
		ds := DbSyncDes{InsertOp, UserTable, ssf.Ssf.GetUUID(), data}
		err = ssf.MQ.Notify(DbSyncNotify, ds)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("notify failed.")
		}
	} else {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Marshal failed.")
	}

	return strconv.FormatInt(ret, 10), nil
}

// GetUser 获取一条用户，支持通过用户名查找
func GetUser(name string) (*User, error) {
	var user User

	err := orm.NewOrm().QueryTable("user").Filter("Username", name).One(&user)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Read failed.")
	}

	ssf.Logger.Debugln(user)

	return &user, err
}

// GetAllUsers 返回全部用户信息
func GetAllUsers() ([]*User, error) {
	var users []*User

	num, err := orm.NewOrm().QueryTable("user").All(&users)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Read failed.")
		return nil, err
	}

	ssf.Logger.Debugln(num, err, users)
	return users, nil
}

// UpdateUser 更新一条用户信息
func UpdateUser(name string, u User) error {
	ssf.Logger.Debugln(name)
	ssf.Logger.Debugln(u)

	num, err := orm.NewOrm().QueryTable("user").Filter("Username", name).Update(orm.Params{
		"Username": u.Username,
		"Password": u.Password,
	})
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Update failed.")
		return err
	}

	if num == 0 {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("User not exist.")
		return errors.New("User not exist.")
	}

	// 发送 update 广播
	data, err := ssf.Serializer.Marshal(u)
	if err == nil {
		ds := DbSyncDes{UpdateOp, UserTable, ssf.Ssf.GetUUID(), data}
		err = ssf.MQ.Notify(DbSyncNotify, ds)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("notify failed.")
		}
	} else {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Marshal failed.")
	}

	ssf.Logger.Debugln(num, err)

	return nil
}

// DeleteUser 删除用户
func DeleteUser(name string) error {
	ssf.Logger.Debugln(name)

	num, err := orm.NewOrm().QueryTable("user").Filter("Username", name).Delete()
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Delete failed.")
		return err
	}

	if num == 0 {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("User not exist.")
		return errors.New("User not exist.")
	}

	// 发送 delete 广播
	u := User{
		Username: name,
	}
	data, err := ssf.Serializer.Marshal(u)
	if err == nil {
		ds := DbSyncDes{DeleteOp, UserTable, ssf.Ssf.GetUUID(), data}
		err = ssf.MQ.Notify(DbSyncNotify, ds)
		if err != nil {
			ssf.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("notify failed.")
		}
	} else {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Marshal failed.")
	}

	ssf.Logger.Debugln(num, err)

	return nil
}
