package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []postgres.SysBaseMenu{
	{ MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: postgres.Meta{Title: "仪表盘", Icon: "setting"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: postgres.Meta{Title: "关于我们", Icon: "info"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: postgres.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: postgres.Meta{Title: "角色管理", Icon: "s-custom"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: postgres.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: postgres.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: postgres.Meta{Title: "用户管理", Icon: "coordinate"}},
	{ MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: postgres.Meta{Title: "个人信息", Icon: "message-solid"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: postgres.Meta{Title: "示例文件", Icon: "s-management"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: postgres.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: postgres.Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: postgres.Meta{Title: "断点续传", Icon: "upload"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: postgres.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: postgres.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: postgres.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{ MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: postgres.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{ MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: postgres.Meta{Title: "系统配置", Icon: "s-operation"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: postgres.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{ MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: postgres.Meta{Title: "字典详情", Icon: "s-order"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: postgres.Meta{Title: "操作历史", Icon: "time"}},
	{ MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: postgres.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
	{ MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: postgres.Meta{Title: "官方网站", Icon: "s-home"}},
	{ MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: postgres.Meta{Title: "服务器状态", Icon: "cloudy"}},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]postgres.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[postgres] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[postgres] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
