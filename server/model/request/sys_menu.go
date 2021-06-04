package request

import (
	"gin-vue-admin/model/postgres"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []postgres.SysBaseMenu
	AuthorityId string
}

