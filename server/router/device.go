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
		DeviceRouter.GET("getDevice", v1.GetDevice)   // 获取设备
		DeviceRouter.POST("updateDevice",v1.UpdateDevice) //更新设备
	}
}