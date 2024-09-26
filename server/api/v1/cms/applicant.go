package cms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ApplicantApi struct {}



// CreateApplicant 创建applicant表
// @Tags Applicant
// @Summary 创建applicant表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "创建applicant表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /applicant/createApplicant [post]
func (applicantApi *ApplicantApi) CreateApplicant(c *gin.Context) {
	var applicant cms.Applicant
	err := c.ShouldBindJSON(&applicant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = applicantService.CreateApplicant(&applicant)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteApplicant 删除applicant表
// @Tags Applicant
// @Summary 删除applicant表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "删除applicant表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /applicant/deleteApplicant [delete]
func (applicantApi *ApplicantApi) DeleteApplicant(c *gin.Context) {
	ID := c.Query("ID")
	err := applicantService.DeleteApplicant(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteApplicantByIds 批量删除applicant表
// @Tags Applicant
// @Summary 批量删除applicant表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /applicant/deleteApplicantByIds [delete]
func (applicantApi *ApplicantApi) DeleteApplicantByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := applicantService.DeleteApplicantByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateApplicant 更新applicant表
// @Tags Applicant
// @Summary 更新applicant表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "更新applicant表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /applicant/updateApplicant [put]
func (applicantApi *ApplicantApi) UpdateApplicant(c *gin.Context) {
	var applicant cms.Applicant
	err := c.ShouldBindJSON(&applicant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = applicantService.UpdateApplicant(applicant)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindApplicant 用id查询applicant表
// @Tags Applicant
// @Summary 用id查询applicant表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Applicant true "用id查询applicant表"
// @Success 200 {object} response.Response{data=cms.Applicant,msg=string} "查询成功"
// @Router /applicant/findApplicant [get]
func (applicantApi *ApplicantApi) FindApplicant(c *gin.Context) {
	ID := c.Query("ID")
	reapplicant, err := applicantService.GetApplicant(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reapplicant, c)
}

// GetApplicantList 分页获取applicant表列表
// @Tags Applicant
// @Summary 分页获取applicant表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ApplicantSearch true "分页获取applicant表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /applicant/getApplicantList [get]
func (applicantApi *ApplicantApi) GetApplicantList(c *gin.Context) {
	var pageInfo cmsReq.ApplicantSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := applicantService.GetApplicantInfoList(pageInfo)
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

// GetApplicantPublic 不需要鉴权的applicant表接口
// @Tags Applicant
// @Summary 不需要鉴权的applicant表接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ApplicantSearch true "分页获取applicant表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /applicant/getApplicantPublic [get]
func (applicantApi *ApplicantApi) GetApplicantPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    applicantService.GetApplicantPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的applicant表接口信息",
    }, "获取成功", c)
}
