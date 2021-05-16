package request

import (
	"gin-vue-admin/model/postgres"
)

type SysAuthority struct {
	postgres.SysAuthority
	Children     []SysAuthority          `json:"children"`
}
