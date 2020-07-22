package model

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	DeviceType  string  `json:"DeviceType" form:"DeviceType" gorm:"comment:'客户手机号'"`
	DeviceName          uint    `json:"sysUserId" form:"sysUserId" gorm:"comment:'管理ID'"`
	DeviceStatus string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:'管理角色ID'"`
}