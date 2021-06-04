package response

import (
	"gin-vue-admin/model/postgres"
)

type SysAPIResponse struct {
	Api postgres.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []postgres.SysApi `json:"apis"`
}
