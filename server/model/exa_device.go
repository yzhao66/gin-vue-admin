package model

import (
	"github.com/jinzhu/gorm"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Device struct {
	gorm.Model
	DeviceType  string  `json:"DeviceType" form:"DeviceType" gorm:"comment:'客户手机号'"`
	DeviceName          string    `json:"DeviceName" form:"DeviceName" gorm:"comment:'管理ID'"`
	NodeName string  `json:"NodeName" form:"NodeName" gorm:"comment:'管理角色ID'"`
	/*DeviceSpec string  `json:"DeviceSpec" form:"DeviceSpec" gorm:"comment:'管理角色ID'"`
	Status v1alpha1.DeviceStatus     `json:"DeviceStatus" form:"DeviceStatus" gorm:"comment:'管理角色ID'"`*/
	CreateTime v1.Time  `json:"CreateTime" form:"CreateTime" gorm:"comment:'管理角色ID'"`
	NameSpace string `json:"NameSpace" form:"NameSpace" gorm:"comment:'管理角色ID'"`
}