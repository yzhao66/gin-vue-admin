package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ExaCron struct {
	gorm.Model
	Name           string    `json:"name" form:"name" gorm:"comment:'名称'"`
	Note           string    `json:"note" form:"note" gorm:"comment:'备注'"`
	CronExpression string    `json:"cronexpression" form:"cronexpression" gorm:"comment:'cron表达式'"`
	FuncName       string    `json:"funcname" form:"funcname" gorm:"comment:'函数名'"`
	ExpiredTime    time.Time `json:"expiredtime" form:"expiredtime" gorm:"comment:'过期时间'"`
	StructName     string	 `json:"structname" form:"structname" gorm:"comment:'结构体名称'"`
}
