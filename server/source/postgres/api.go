package postgres

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/postgres"
	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []model.SysApi{
	{Path: "/base/login", Description: "用户登录",ApiGroup: "base", Method: "POST"},
	{Path:  "/user/register", Description:  "用户注册", ApiGroup: "user", Method:"POST"},
	{Path:  "/api/createApi",  Description: "创建api", ApiGroup: "api", Method:"POST"},
	{Path:  "/api/getApiList",  Description: "获取api列表", ApiGroup: "api", Method:"POST"},
	{Path:  "/api/getApiById",  Description: "获取api详细信息",ApiGroup:  "api", Method:"POST"},
	{Path:  "/api/deleteApi",  Description: "删除Api", ApiGroup: "api", Method:"POST"},
	{Path: "/api/updateApi",  Description: "更新Api", ApiGroup: "api", Method:"POST"},
	{Path: "/api/getAllApis",  Description: "获取所有api", ApiGroup: "api", Method:"POST"},
	{Path: "/authority/createAuthority",  Description: "创建角色", ApiGroup: "authority", Method:"POST"},
	{Path:  "/authority/deleteAuthority",  Description: "删除角色", ApiGroup: "authority", Method:"POST"},
	{Path: "/authority/getAuthorityList",  Description: "获取角色列表", ApiGroup: "authority", Method:"POST"},
	{Path: "/menu/getMenu",  Description: "获取菜单树", ApiGroup: "menu", Method:"POST"},
	{Path: "/menu/getMenuList", Description: "分页获取基础menu列表", ApiGroup: "menu", Method:"POST"},
	{Path:  "/menu/addBaseMenu", Description: "新增菜单", ApiGroup: "menu", Method:"POST"},
	{Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由", ApiGroup: "menu", Method:"POST"},
	{Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系", ApiGroup: "menu", Method:"POST"},
	{Path: "/menu/getMenuAuthority", Description: "获取指定角色menu",ApiGroup:  "menu", Method:"POST"},
	{Path: "/menu/deleteBaseMenu", Description: "删除菜单", ApiGroup: "menu", Method:"POST"},
	{Path:  "/menu/updateBaseMenu", Description: "更新菜单", ApiGroup: "menu", Method:"POST"},
	{Path:  "/menu/getBaseMenuById", Description: "根据id获取菜单",ApiGroup:  "menu", Method:"POST"},
	{Path: "/user/changePassword", Description: "修改密码", ApiGroup: "user", Method:"POST"},
	{Path:  "/user/getUserList", Description: "获取用户列表", ApiGroup: "user", Method:"POST"},
	{Path:  "/user/setUserAuthority", Description: "修改用户角色", ApiGroup: "user", Method:"POST"},
	{Path: "/fileUploadAndDownload/upload", Description: "文件上传示例", ApiGroup: "fileUploadAndDownload", Method:"POST"},
	{Path:  "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表",ApiGroup:  "fileUploadAndDownload", Method:"POST"},
	{Path:  "/casbin/updateCasbin", Description: "更改角色api权限", ApiGroup: "casbin", Method:"POST"},
	{Path:  "/casbin/getPolicyPathByAuthorityId",Description:  "获取权限列表",ApiGroup:  "casbin", Method:"POST"},
	{Path:  "/fileUploadAndDownload/deleteFile",Description:  "删除文件", ApiGroup: "fileUploadAndDownload", Method:"POST"},
	{Path:  "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出)", ApiGroup: "jwt", Method:"POST"},
	{Path:  "/authority/setDataAuthority", Description: "设置角色资源权限", ApiGroup: "authority", Method:"POST"},
	{Path: "/system/getSystemConfig", Description: "获取配置文件内容", ApiGroup: "system", Method:"POST"},
	{Path:  "/system/setSystemConfig", Description: "设置配置文件内容",ApiGroup:  "system", Method:"POST"},
	{Path:  "/customer/customer", Description: "创建客户", ApiGroup: "customer", Method:"POST"},
	{Path:  "/customer/customer", Description: "更新客户", ApiGroup: "customer", Method:"PUT"},
	{Path:  "/customer/customer", Description: "删除客户", ApiGroup: "customer", Method:"DELETE"},
	{Path:  "/customer/customer", Description: "获取单一客户", ApiGroup: "customer", Method:"GET"},
	{Path:  "/customer/customerList", Description: "获取客户列表", ApiGroup: "customer", Method:"GET"},
	{Path:  "/casbin/casbinTest/:pathParam", Description: "RESTFUL模式测试", ApiGroup: "casbin", Method:"GET"},
	{Path:  "/autoCode/createTemp", Description: "自动化代码",ApiGroup:  "autoCode", Method:"POST"},
	{Path:  "/authority/updateAuthority", Description: "更新角色信息", ApiGroup: "authority", Method:"PUT"},
	{Path:  "/authority/copyAuthority", Description: "拷贝角色", ApiGroup: "authority", Method:"POST"},
	{Path:  "/user/deleteUser", Description: "删除用户", ApiGroup: "user", Method:"DELETE"},
	{Path:  "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容", ApiGroup: "sysDictionaryDetail", Method:"POST"},
	{Path:  "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容", ApiGroup: "sysDictionaryDetail", Method:"DELETE"},
	{Path:  "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容", ApiGroup: "sysDictionaryDetail", Method:"PUT"},
	{Path:  "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容", ApiGroup: "sysDictionaryDetail", Method:"GET"},
	{Path:  "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表", ApiGroup: "sysDictionaryDetail", Method:"GET"},
	{Path:  "/sysDictionary/createSysDictionary", Description: "新增字典", ApiGroup: "sysDictionary", Method:"POST"},
	{Path:  "/sysDictionary/deleteSysDictionary", Description: "删除字典", ApiGroup: "sysDictionary", Method:"DELETE"},
	{Path:  "/sysDictionary/updateSysDictionary", Description: "更新字典", ApiGroup: "sysDictionary", Method:"PUT"},
	{Path:  "/sysDictionary/findSysDictionary", Description: "根据ID获取字典", ApiGroup: "sysDictionary", Method:"GET"},
	{Path:  "/sysDictionary/getSysDictionaryList", Description: "获取字典列表", ApiGroup: "sysDictionary", Method:"GET"},
	{Path:  "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录",ApiGroup:  "sysOperationRecord", Method:"POST"},
	{Path:  "/sysOperationRecord/deleteSysOperationRecord",Description:  "删除操作记录", ApiGroup: "sysOperationRecord", Method:"DELETE"},
	{Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录", ApiGroup: "sysOperationRecord", Method:"GET"},
	{Path:  "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表",ApiGroup:  "sysOperationRecord", Method:"GET"},
	{Path: "/autoCode/getTables", Description: "获取数据库表", ApiGroup: "autoCode", Method:"GET"},
	{Path:  "/autoCode/getDB", Description: "获取所有数据库", ApiGroup: "autoCode", Method:"GET"},
	{Path:  "/autoCode/getColumn", Description: "获取所选table的所有字段", ApiGroup: "autoCode", Method:"GET"},
	{Path:  "/sysOperationRecord/deleteSysOperationRecordByIds",Description:  "批量删除操作历史", ApiGroup: "sysOperationRecord", Method:"DELETE"},
	{Path:  "/simpleUploader/upload", Description: "插件版分片上传", ApiGroup: "simpleUploader", Method:"POST"},
	{Path:  "/simpleUploader/checkFileMd5",Description:  "文件完整度验证", ApiGroup: "simpleUploader", Method:"GET"},
	{Path:  "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件", ApiGroup: "simpleUploader", Method:"GET"},
	{Path:  "/user/setUserInfo", Description: "设置用户信息", ApiGroup: "user", Method:"PUT"},
	{Path:  "/system/getServerInfo", Description: "获取服务器信息", ApiGroup: "system", Method:"POST"},
	{Path:  "/email/emailTest", Description: "发送测试邮件", ApiGroup: "email", Method:"POST"},
	{Path:  "/autoCode/preview", Description: "预览自动化代码", ApiGroup: "autoCode", Method:"POST"},
	{Path:  "/excel/importExcel", Description: "导入excel", ApiGroup: "excel", Method:"POST"},
	{Path:  "/excel/loadExcel", Description: "下载excel", ApiGroup: "excel", Method:"GET"},
	{Path:  "/excel/exportExcel", Description: "导出excel",ApiGroup:  "excel", Method:"POST"},
	{Path:  "/excel/downloadTemplate", Description: "下载excel模板", ApiGroup: "excel", Method:"GET"},
	{Path:  "/api/deleteApisByIds", Description: "批量删除api", ApiGroup: "api", Method:"DELETE"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]postgres.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Postgres] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Postgres] --> sys_apis 表初始数据成功!")
		return nil
	})
}
