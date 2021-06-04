package request

import (
	"gin-vue-admin/model/postgres"
)

type SysOperationRecordSearch struct {
	postgres.SysOperationRecord
	PageInfo
}

type SysOperationRecordResponse struct {
	Id int `json:"id"`
}

type SysOperationRecordResponseRes struct {
	Resp string `json:"resp"`
}
