package main

import (
	_ "dbmanager/routers"

	"github.com/astaxie/beego"
)

/*
type Student struct {
	Id    int // 主键
	Name  string
	Age   int
	Sex   string
	Score float32
	Addr  string
}*/

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 初始化dbmanager
	//Init()
	/*
		o := orm.NewOrm()
		o.Using("default")
		stu := new(Student)

		stu.Name = "bei"
		stu.Age = 25
		stu.Sex = "m"
		stu.Score = 88
		stu.Addr = "hunan.leiyang"

		fmt.Println(o.Insert(stu))
	*/
	beego.Run()
}
