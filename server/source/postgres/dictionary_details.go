package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionary_details 表数据初始化
func (d *dictionaryDetail) Init() error {
	var details = []postgres.SysDictionaryDetail{
		{Label: "smallint", Value: 1,  Status: status,Sort: 1, SysDictionaryID: 2},
		{Label:  "mediumint",  Value:2,  Status: status, Sort: 2,SysDictionaryID: 2},
		{Label:  "int",  Value:3,  Status: status,Sort:  3, SysDictionaryID:2},
		{Label: "bigint",  Value:4,  Status: status, Sort: 4, SysDictionaryID:2},
		{Label:  "date", Value: 0,  Status: status, Sort: 0, SysDictionaryID:3},
		{Label:  "time",  Value:1,  Status: status,Sort:  1, SysDictionaryID:3},
		{Label:  "year",  Value:2,  Status: status, Sort: 2, SysDictionaryID:3},
		{Label:  "datetime",  Value:3,  Status: status, Sort: 3, SysDictionaryID:3},
		{Label:  "timestamp",  Value:5,  Status: status, Sort: 5, SysDictionaryID:3},
		{Label:  "float",  Value:0,  Status: status, Sort: 0, SysDictionaryID:4},
		{Label: "double",  Value:1,  Status: status, Sort: 1, SysDictionaryID:4},
		{Label:  "decimal", Value: 2,  Status: status, Sort: 2, SysDictionaryID:4},
		{Label:  "char", Value: 0,  Status: status, Sort: 0, SysDictionaryID:5},
		{Label:  "varchar", Value: 1,  Status: status, Sort: 1, SysDictionaryID:5},
		{Label:  "tinyblob", Value: 2,  Status: status, Sort: 2, SysDictionaryID:5},
		{Label:  "tinytext", Value: 3,  Status: status, Sort: 3, SysDictionaryID:5},
		{Label:  "text", Value: 4,  Status: status, Sort: 4, SysDictionaryID:5},
		{Label: "blob",  Value:5,  Status: status, Sort: 5, SysDictionaryID:5},
		{Label:  "mediumblob", Value: 6,  Status: status, Sort: 6, SysDictionaryID:5},
		{Label:  "mediumtext", Value: 7,  Status: status, Sort: 7, SysDictionaryID:5},
		{Label: "longblob", Value: 8,  Status: status, Sort: 8, SysDictionaryID:5},
		{Label: "longtext", Value: 9,  Status: status, Sort: 9, SysDictionaryID:5},
		{Label: "tinyint",  Value:0,  Status: status, Sort: 0, SysDictionaryID:6},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]postgres.SysDictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[postgres:] --> sys_dictionary_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[postgres] --> sys_dictionary_details 表初始数据成功!")
		return nil
	})
}
