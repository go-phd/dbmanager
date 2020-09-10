package controllers

import (
	"dbmanager/sqldb"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-phd/ssf"
	"github.com/sirupsen/logrus"
)

// Operations about object
type GroupController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Group	true		"The object content"
// @Success 200 {string} models.Group.Id
// @Failure 403 body is empty
// @router / [post]
func (o *GroupController) Post() {
	var sg sqldb.CshareGroupTable
	fmt.Println(o.Ctx.Input.RequestBody)
	fmt.Println(string(o.Ctx.Input.RequestBody))
	/*
		fmt.Println("============")
		obest := models.Group{
			GroupId:   "333",
			Score:      1,
			PlayerName: "111",
		}

		ret, _ := json.Marshal(obest)
		fmt.Println(ret)
		fmt.Println(string(ret))
		var obr models.Group
		json.Unmarshal(o.Ctx.Input.RequestBody, &obr)
		fmt.Println(ob)
		fmt.Println("============")
	*/
	json.Unmarshal(o.Ctx.Input.RequestBody, &sg)
	fmt.Println(sg)
	err := sqldb.ShareGroup.Insert(sg)
	if err != nil {
		ssf.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Insert failed.")
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "post success."
	}

	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Group
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *GroupController) Get() {
	name := o.Ctx.Input.Param(":name")
	ssf.Logger.Debugln(name)
	if name != "" {
		sg, err := sqldb.ShareGroup.Queue(name)
		ssf.Logger.Debugln(sg)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = sg
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Group
// @Failure 403 :objectId is empty
// @router / [get]
func (o *GroupController) GetAll() {
	sgs, err := sqldb.ShareGroup.QueueAll()
	ssf.Logger.Debugln(sgs)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = sgs
	}

	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Group	true		"The body"
// @Success 200 {object} models.Group
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *GroupController) Put() {
	//name := o.Ctx.Input.Param(":name")
	var sg sqldb.CshareGroupTable
	json.Unmarshal(o.Ctx.Input.RequestBody, &sg)

	err := sqldb.ShareGroup.Update(sg)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *GroupController) Delete() {
	name := o.Ctx.Input.Param(":name")
	err := sqldb.ShareGroup.Delete(name)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "delete success!"
	}
	o.ServeJSON()
}
