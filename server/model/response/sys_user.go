package response

import (
	"gin-vue-admin/model/postgres"
	"gin-vue-admin/model/postgres/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUserResponse struct {
	User postgres.SysUser `json:"user"`
}

type LoginResponse struct {
	User      User `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}



type User struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"userName"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickName"`
	HeaderImg   string       `json:"headerImg"`
	AuthorityId string       `json:"authorityId"`
	Authority request.SysAuthority       `json:"authority"`
}