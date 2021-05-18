package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"
	"time"
)

var Migrate = new(migrate)

type migrate struct{}

var migrates = postgres.SysMigration{true,time.Now(),time.Now()}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: exa_file_upload_and_downloads 表初始化数据
func (m *migrate) Init() error {

	if err:= global.GVA_DB.Create(&migrates).Error;err != nil {
		return err
	}else {
		color.Info.Println("\n[postgres] --> migrate 表初始数据成功!")
		return nil
	}
}
