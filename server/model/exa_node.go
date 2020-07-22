package model


import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	NodeName       string  `json:"NodeName" form:"NodeName" gorm:"comment:'客户名'"`

}