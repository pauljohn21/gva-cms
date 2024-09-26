package cms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PolicyholderApi struct {}



// CreatePolicyholder 创建投保人
// @Tags Policyholder
// @Summary 创建投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Policyholder true "创建投保人"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /policyholder/createPolicyholder [post]
func (policyholderApi *PolicyholderApi) CreatePolicyholder(c *gin.Context) {
	var policyholder cms.Policyholder
	err := c.ShouldBindJSON(&policyholder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = policyholderService.CreatePolicyholder(&policyholder)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePolicyholder 删除投保人
// @Tags Policyholder
// @Summary 删除投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Policyholder true "删除投保人"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /policyholder/deletePolicyholder [delete]
func (policyholderApi *PolicyholderApi) DeletePolicyholder(c *gin.Context) {
	ID := c.Query("ID")
	err := policyholderService.DeletePolicyholder(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePolicyholderByIds 批量删除投保人
// @Tags Policyholder
// @Summary 批量删除投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /policyholder/deletePolicyholderByIds [delete]
func (policyholderApi *PolicyholderApi) DeletePolicyholderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := policyholderService.DeletePolicyholderByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePolicyholder 更新投保人
// @Tags Policyholder
// @Summary 更新投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Policyholder true "更新投保人"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /policyholder/updatePolicyholder [put]
func (policyholderApi *PolicyholderApi) UpdatePolicyholder(c *gin.Context) {
	var policyholder cms.Policyholder
	err := c.ShouldBindJSON(&policyholder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = policyholderService.UpdatePolicyholder(policyholder)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPolicyholder 用id查询投保人
// @Tags Policyholder
// @Summary 用id查询投保人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Policyholder true "用id查询投保人"
// @Success 200 {object} response.Response{data=cms.Policyholder,msg=string} "查询成功"
// @Router /policyholder/findPolicyholder [get]
func (policyholderApi *PolicyholderApi) FindPolicyholder(c *gin.Context) {
	ID := c.Query("ID")
	repolicyholder, err := policyholderService.GetPolicyholder(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repolicyholder, c)
}

// GetPolicyholderList 分页获取投保人列表
// @Tags Policyholder
// @Summary 分页获取投保人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.PolicyholderSearch true "分页获取投保人列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /policyholder/getPolicyholderList [get]
func (policyholderApi *PolicyholderApi) GetPolicyholderList(c *gin.Context) {
	var pageInfo cmsReq.PolicyholderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := policyholderService.GetPolicyholderInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetPolicyholderPublic 不需要鉴权的投保人接口
// @Tags Policyholder
// @Summary 不需要鉴权的投保人接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.PolicyholderSearch true "分页获取投保人列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /policyholder/getPolicyholderPublic [get]
func (policyholderApi *PolicyholderApi) GetPolicyholderPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    policyholderService.GetPolicyholderPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的投保人接口信息",
    }, "获取成功", c)
}
