package response

import (
	"gin-vue-admin/model"
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
	Menus []model.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenu `json:"menu"`
}
