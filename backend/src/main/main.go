package main

import (
	"fmt"
	_ "main/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
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

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
