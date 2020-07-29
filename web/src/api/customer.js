import service from '@/utils/request'

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const createExaCustomer = (data) => {
    return service({
        url: "/customer/customer",
        method: 'post',
        data
    })
}



// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
export const updateExaCustomer = (data) => {
    return service({
        url: "/customer/customer",
        method: 'put',
        data
    })
}


// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
export const deleteExaCustomer = (data) => {
    return service({
        url: "/customer/customer",
        method: 'delete',
        data
    })
}


// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
export const getExaCustomer = (params) => {
    return service({
        url: "/customer/customer",
        method: 'get',
        params
    })
}


// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
export const getExaCustomerList = (params) => {
    return service({
        url: "/customer/customerList",
        method: 'get',
        params
    })
}

export const getDeviceFile = (data) => {
    // eslint-disable-next-line no-debugger
    return service({
        url: "/device/getDeviceFile",
        method: 'post',
        data
    })
}

export const updateDevice = (data) => {
    // eslint-disable-next-line no-,no-debugger
    return service({
        url: "/device/updateDevice",
        method: 'post',
        data
    })
}

export const getNodes = () => {
    // eslint-disable-next-line no-
    return service({
        url: "/device/getNodes",
        method: 'post',
    })
}
export const getDeviceDetails = () => {
    // eslint-disable-next-line no-
    return service({
        url: "/device/getDeviceDetails",
        method: 'post',
    })
}
export const getCronLists = () => {
    // eslint-disable-next-line no-
    return service({
        url: "/device/getCronList",
        method: 'post',
    })
}