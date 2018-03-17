package main

import (
	"fmt"
	_ "main/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlURI := beego.AppConfig.String("mysqlurls")
	mysqlDB := beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@"+mysqlURI+"/"+mysqlDB+"?charset=utf8")
}

func main() {
	orm.Debug = true
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
