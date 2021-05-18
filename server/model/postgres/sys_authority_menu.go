package postgres

type SysMenu struct {
	SysAuthorityAuthorityId int                 `json:"authorityId" gorm:"comment:角色ID"`
	SysBaseMenuId      int                 `json:"menuId" gorm:"comment:菜单ID"`
}

func (s SysMenu) TableName() string {
	return "sys_authority_menus"
}
