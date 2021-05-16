package postgres

import uuid "github.com/satori/go.uuid"

type SysUserAuthority struct {
	AuthorityId string `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	UserId uuid.UUID `json:"user_id" gorm:"comment:用户UUID"`
}