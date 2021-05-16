package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var File = new(file)

type file struct{}

var files = []postgres.ExaFileUploadAndDownload{
	{Name: "10.png", Url: "http://qmplusimg.henrongyi.top/gvalogo.png", Tag: "png", Key: "158787308910.png"},
	{Name: "logo.png", Url: "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", Tag: "png", Key: "1587973709logo.png"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: exa_file_upload_and_downloads 表初始化数据
func (f *file) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]postgres.ExaFileUploadAndDownload{}).RowsAffected == 2 {
			color.Danger.Println("\n[postgres] --> exa_file_upload_and_downloads 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[postgres] --> exa_file_upload_and_downloads 表初始数据成功!")
		return nil
	})
}
