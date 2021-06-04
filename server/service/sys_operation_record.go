package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/postgres"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

//@author: [granty1](https://github.com/granty1)
//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord postgres.SysOperationRecord
//@return: err error

func CreateSysOperationRecord(sysOperationRecord postgres.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]postgres.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord postgres.SysOperationRecord
//@return: err error

func DeleteSysOperationRecord(sysOperationRecord postgres.SysOperationRecord) (err error) {
	err = global.GVA_DB.Delete(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord postgres.SysOperationRecord

func GetSysOperationRecord(id uint) (err error, sysOperationRecord postgres.SysOperationRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info request.SysOperationRecordSearch
//@return: err error, list interface{}, total int64

func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	//db := global.GVA_DB.Model(&postgres.SysOperationRecord{})
	db := global.GVA_DB.Table("sys_operation_records as b").Joins("left join sys_users as su on su.id=b.user_id").Select("su.nick_name,su.username,b.id,b.created_at,b.ip,b.method,b.path,b.status,b.user_id,b.body,(CASE  WHEN LENGTH(b.resp)> 0  THEN '1' else ''END) as resp")
	var sysOperationRecords []response.SysOperationRecordInfoList
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("b.method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("b.path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("b.status = ?", info.Status)
	}
	err = db.Count(&total).Error
	//err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	err = db.Order("b.id desc").Limit(limit).Offset(offset).Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}


//@author: [granty1](https://github.com/kiwi520)
//@author: [piexlmax](https://github.com/kiwi520)
//@function: GetSysOperationRecordResponse
//@description: 获取操作记录响应信息
//@param: info request.SysOperationRecordSearch
//@return: err error, response string

func GetSysOperationRecordResponse(id int) (err error, res string) {
	err = global.GVA_DB.Table("sys_operation_records").Select("resp").Where("id=?",id).Find(&res).Error
	return err, res
}
