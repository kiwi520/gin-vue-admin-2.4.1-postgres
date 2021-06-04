package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary postgres.SysDictionary
//@return: err error

func CreateSysDictionary(sysDictionary postgres.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&postgres.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary postgres.SysDictionary
//@return: err error

func DeleteSysDictionary(sysDictionary postgres.SysDictionary) (err error) {
	//err = global.GVA_DB.Delete(&sysDictionary).Delete(&sysDictionary.SysDictionaryDetails).Error
	//return err
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *postgres.SysDictionary
//@return: err error

func UpdateSysDictionary(sysDictionary *postgres.SysDictionary) (err error) {
	var dict postgres.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type == sysDictionary.Type {
		err = db.Updates(sysDictionaryMap).Error
	} else {
		if (!errors.Is(global.GVA_DB.First(&postgres.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的type，不允许创建")
		}
		err = db.Updates(sysDictionaryMap).Error

	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary postgres.SysDictionary

func GetSysDictionary(Type string, Id uint) (err error, sysDictionary postgres.SysDictionary) {
	err = global.GVA_DB.Where("type = ? OR id = ?", Type, Id).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func GetSysDictionaryInfoList(info request.SysDictionarySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&postgres.SysDictionary{})
	var sysDictionarys []postgres.SysDictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return err, sysDictionarys, total
}
