package main

import (
	_ "kandao_backend/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "kandao_backend/models/appinfo" //这种导入方式会执行package里的init()函数
)

func CreateTable(){
	name := "appinfo"
	err := orm.RunSyncdb(name,false,true)
	if err != nil{
		beego.Error(err)
	}
}

func main() {

	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:asdasd@tcp(127.0.0.1:3306)/backend2?charset=utf8")
	//o := orm.NewOrm()
	//o.Using("default")
	CreateTable()
	beego.Run()
}

