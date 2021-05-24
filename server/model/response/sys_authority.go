package response

import (
	"gin-vue-admin/model/postgres"
)

type SysAuthorityResponse struct {
	Authority postgres.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      postgres.SysAuthority `json:"authority"`
	OldAuthorityId string             `json:"oldAuthorityId"`
}
