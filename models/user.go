package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/go-phd/ssf"
	"github.com/sirupsen/logrus"
)

func init() {
	orm.RegisterModel(new(User))
	ssf.Logger.Infoln("user table init success.")
}

// User 用户表
type User struct {
	ID       int64 //主键
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
	u.ID = 0
	ret, err := orm.NewOrm().Insert(&u)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("insert failed.")
		return "", errors.New("Failed to add a new user.")
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

	ssf.Logger.Debugln(num, err)
	return users, nil
}

// UpdateUser 更新一条用户信息
func UpdateUser(name string, user *User) error {
	ssf.Logger.Debugln(user)
	ssf.Logger.Debugln(name)

	num, err := orm.NewOrm().QueryTable("user").Filter("Username", name).Update(orm.Params{
		"Username": user.Username,
		"Password": user.Password,
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
		}).Errorln("Update failed.")
		return errors.New("User Not Exist")
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
		}).Errorln("Update failed.")
		return errors.New("User Not Exist")
	}

	ssf.Logger.Debugln(num, err)

	return nil
}
