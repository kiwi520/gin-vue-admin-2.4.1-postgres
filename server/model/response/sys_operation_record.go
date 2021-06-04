package response

import "gin-vue-admin/model/postgres"

type SysOperationRecordInfoList struct {
	postgres.SysOperationRecord
	NickName string `json:"nick_name"`
	Username string `json:"username"`
}
