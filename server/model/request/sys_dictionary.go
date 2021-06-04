package request

import (
	"gin-vue-admin/model/postgres"
)

type SysDictionarySearch struct {
	postgres.SysDictionary
	PageInfo
}
