package response

import (
	"gin-vue-admin/model/postgres"
	"gin-vue-admin/model/postgres/request"
)

type SysMenusResponse struct {
	//Menus []model.SysMenu `json:"menus"`
	Menus []request.MenuList `json:"menus"`

}

type SysAuthorityMenusResponse struct {
	//Menus []model.SysMenu `json:"menus"`
	Menus []request.AuthorityMenuList `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu postgres.SysBaseMenu `json:"menu"`
}


type SysBaseMenu struct {
	postgres.SysBaseMenu
	Children     []SysBaseMenu          `json:"children"`
}