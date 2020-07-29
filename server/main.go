package main

import (
	"fmt"
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"github.com/gin-gonic/gin"
	"reflect"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	gin.Default()
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	// case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	initialize.InitCron();
	defer global.GVA_DB.Close()
	//测试代码
	v := reflect.ValueOf(v1.Add)
	paramList :=[]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	retList := v.Call(paramList)
	fmt.Println(retList[0].Int())
	//
	core.RunWindowsServer()

}
