package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func GetCronList() (err error, list interface{}, total int){


	db := global.GVA_DB.Model(&model.ExaCron{})
	/*db.LogMode(true)*/
	var exaCronList []model.ExaCron
	err = db.Find(&exaCronList).Error
	err = db.Find(&exaCronList).Count(&total).Error
	if err != nil {
		return err, exaCronList, total
	}
	return err, exaCronList, total
}