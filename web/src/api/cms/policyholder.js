import service from '@/utils/request'
// @Tags Policyholder
// @Summary 创建投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policyholder true "创建投保人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /policyholder/createPolicyholder [post]
export const createPolicyholder = (data) => {
  return service({
    url: '/policyholder/createPolicyholder',
    method: 'post',
    data
  })
}

// @Tags Policyholder
// @Summary 删除投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policyholder true "删除投保人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /policyholder/deletePolicyholder [delete]
export const deletePolicyholder = (params) => {
  return service({
    url: '/policyholder/deletePolicyholder',
    method: 'delete',
    params
  })
}

// @Tags Policyholder
// @Summary 批量删除投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除投保人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /policyholder/deletePolicyholder [delete]
export const deletePolicyholderByIds = (params) => {
  return service({
    url: '/policyholder/deletePolicyholderByIds',
    method: 'delete',
    params
  })
}

// @Tags Policyholder
// @Summary 更新投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Policyholder true "更新投保人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /policyholder/updatePolicyholder [put]
export const updatePolicyholder = (data) => {
  return service({
    url: '/policyholder/updatePolicyholder',
    method: 'put',
    data
  })
}

// @Tags Policyholder
// @Summary 用id查询投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Policyholder true "用id查询投保人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /policyholder/findPolicyholder [get]
export const findPolicyholder = (params) => {
  return service({
    url: '/policyholder/findPolicyholder',
    method: 'get',
    params
  })
}

// @Tags Policyholder
// @Summary 分页获取投保人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取投保人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /policyholder/getPolicyholderList [get]
export const getPolicyholderList = (params) => {
  return service({
    url: '/policyholder/getPolicyholderList',
    method: 'get',
    params
  })
}

// @Tags Policyholder
// @Summary 不需要鉴权的投保人接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.PolicyholderSearch true "分页获取投保人列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /policyholder/getPolicyholderPublic [get]
export const getPolicyholderPublic = () => {
  return service({
    url: '/policyholder/getPolicyholderPublic',
    method: 'get',
  })
}
