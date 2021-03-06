package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	req "gin-vue-admin/model/postgres/request"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gorm.io/gorm"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]postgres.SysMenu

func getMenuTreeMap(authorityId string) (err error, treeMap map[int][]req.MenuList) {
	//var allMenus []postgres.SysMenu
	//treeMap = make(map[string][]postgres.SysMenu)
	//err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	//for _, v := range allMenus {
	//	treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	//}
	//return err, treeMap
	var allMenus []req.MenuList
	treeMap = make(map[int][]req.MenuList)
	err = global.GVA_DB.Table("sys_base_menus as b").Joins("left join sys_authority_menus as sm on sm.sys_base_menu_id=b.id").Where("sm.sys_authority_authority_id=?",authorityId).Order("sort").Find(&allMenus).Error
	//err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		pid,_:= strconv.Atoi(v.ParentId)
		treeMap[pid] = append(treeMap[pid], v)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []postgres.SysMenu

func GetMenuTree(authorityId string) (err error, menus []req.MenuList) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree[0]

	println("menuTree")
	println(menus)
	println("menuTree")
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *postgres.SysMenu, treeMap map[string][]postgres.SysMenu
//@return: err error

func getChildrenList(menu *req.MenuList, treeMap map[int][]req.MenuList) (err error) {
	menu.Children = treeMap[int(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func GetInfoList() (err error, list interface{}, total int64) {
	var menuList []response.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *postgres.SysBaseMenu, treeMap map[string][]postgres.SysBaseMenu
//@return: err error

func getBaseChildrenList(menu *response.SysBaseMenu, treeMap map[string][]response.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu postgres.SysBaseMenu
//@return: err error

func AddBaseMenu(menu postgres.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&postgres.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]postgres.SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]response.SysBaseMenu) {
	var allMenus []postgres.SysBaseMenu
	treeMap = make(map[string][]response.SysBaseMenu)
	//err = global.GVA_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	err = global.GVA_DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		var menu response.SysBaseMenu
		menu.ID = v.ID
		menu.CreatedAt = v.CreatedAt
		menu.MenuLevel = v.MenuLevel
		menu.Path = v.Path
		menu.ParentId = v.ParentId
		menu.Name = v.Name
		menu.Hidden = v.Hidden
		menu.Component = v.Component
		menu.Sort = v.Sort
		menu.Meta = v.Meta
		//MenuLevel     uint   `json:"-" gorm:"type:int;comment:菜单等级ID"`
		//ParentId      string `json:"parentId" gorm:"comment:父菜单ID"`
		//Path          string `json:"path" gorm:"comment:路由path"`
		//Name          string `json:"name" gorm:"comment:路由name"`
		//Hidden        bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
		//Component     string `json:"component" gorm:"comment:对应前端文件路径"`
		//Sort          int    `json:"sort" gorm:"type:int;comment:排序标记"`
		//Meta          `json:"meta" gorm:"comment:附加属性"`
		treeMap[v.ParentId] = append(treeMap[v.ParentId], menu)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []postgres.SysBaseMenu

func GetBaseMenuTree() (err error, menus []response.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []postgres.SysBaseMenu, authorityId string
//@return: err error

func AddMenuAuthority(menu req.AddMenuAuthorityInfo) (err error) {
	var auth postgres.SysAuthority
	auth.AuthorityId = menu.AuthorityId
	//auth.SysBaseMenus = menus
	authorityId,_:=strconv.Atoi(menu.AuthorityId)
	var muenlist []postgres.SysMenu

	for _,item := range menu.Menus {
		muenlist = append(muenlist,postgres.SysMenu{
			SysAuthorityAuthorityId: authorityId,
			SysBaseMenuId: item.MenuId,
		})
	}
	err = SetMenuAuthority(&muenlist,menu.AuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []postgres.SysMenu

func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []req.AuthorityMenuList) {
	err = global.GVA_DB.Table("sys_authority_menus").Where("sys_authority_authority_id = ? ", info.AuthorityId).Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GVA_DB.Raw(sql, info.AuthorityId).Scan(&menus).Error
	//println("menus")
	//println(len(menus))
	//println("menus")
	return err, menus
}
