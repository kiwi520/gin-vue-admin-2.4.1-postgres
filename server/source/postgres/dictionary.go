package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

var status = new(bool)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_dictionaries 表数据初始化
func (d *dictionary) Init() error {
	*status = true
	var dictionaries = []postgres.SysDictionary{
		{ Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{ Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{ Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{ Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{ Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{ Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]postgres.SysDictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Postgres] --> sys_dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Postgres] --> sys_dictionaries 表初始数据成功!")
		return nil
	})
}
