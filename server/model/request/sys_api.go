package request

import (
	"gin-vue-admin/model/postgres"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	postgres.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
