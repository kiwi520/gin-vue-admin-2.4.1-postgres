package request

import (
	"gin-vue-admin/model/postgres"
)

type SysDictionaryDetailSearch struct {
	postgres.SysDictionaryDetail
	PageInfo
}
