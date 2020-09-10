package controllers

import (
	"dbmanager/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"uid": uid}
	}

	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, err := models.GetAllUsers()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = users
	}

	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	name := u.GetString(":name")
	if name != "" {
		user, err := models.GetUser(name)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	var user models.User
	name := u.GetString(":name")
	if name != "" {
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		err := models.UpdateUser(name, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success."
		}
	}

	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	name := u.GetString(":name")
	if name != "" {
		err := models.DeleteUser(name)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "delete success."
		}
	}

	u.ServeJSON()
}
