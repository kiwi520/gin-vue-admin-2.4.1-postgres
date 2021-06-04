import service from '@/utils/request'

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
export const deleteSysOperationRecord = (data) => {
    return service({
        url: "/sysOperationRecord/deleteSysOperationRecord",
        method: 'delete',
        data
    })
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
export const deleteSysOperationRecordByIds = (data) => {
    return service({
        url: "/sysOperationRecord/deleteSysOperationRecordByIds",
        method: 'delete',
        data
    })
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysOperationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
export const getSysOperationRecordList = (params) => {
    return service({
        url: "/sysOperationRecord/getSysOperationRecordList",
        method: 'get',
        params
    })
}

//@author: [granty1](https://github.com/kiwi520)
//@author: [piexlmax](https://github.com/kiwi520)
//@function: GetSysOperationRecordResponse
//@description: 获取操作记录响应信息
//@param: info request.SysOperationRecordSearch
//@return: err error, response string
export const SysOperationRecordResponse = (data) => {
    return service({
        url: "/sysOperationRecord/SysOperationRecordResponse",
        method: 'post',
        data: data
    })
}