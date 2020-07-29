package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(Router *gin.RouterGroup) {
	DeviceRouter := Router.Group("device").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		DeviceRouter.POST("getDeviceFile", v1.GetDeviceFile)   // 获取设备
		DeviceRouter.POST("updateDevice",v1.UpdateDevice) //更新设备
		DeviceRouter.POST("getNodes",v1.GetNodes)
		DeviceRouter.POST("getDeviceDetails",v1.GetDeviceDetails)
		DeviceRouter.POST("getCronList",v1.GetCronLists)
	}
}