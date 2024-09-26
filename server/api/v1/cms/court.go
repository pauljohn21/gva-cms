package cms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CourtApi struct {}



// CreateCourt 创建法院
// @Tags Court
// @Summary 创建法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "创建法院"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /C/createCourt [post]
func (CApi *CourtApi) CreateCourt(c *gin.Context) {
	var C cms.Court
	err := c.ShouldBindJSON(&C)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CService.CreateCourt(&C)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCourt 删除法院
// @Tags Court
// @Summary 删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "删除法院"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /C/deleteCourt [delete]
func (CApi *CourtApi) DeleteCourt(c *gin.Context) {
	ID := c.Query("ID")
	err := CService.DeleteCourt(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCourtByIds 批量删除法院
// @Tags Court
// @Summary 批量删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /C/deleteCourtByIds [delete]
func (CApi *CourtApi) DeleteCourtByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := CService.DeleteCourtByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCourt 更新法院
// @Tags Court
// @Summary 更新法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "更新法院"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /C/updateCourt [put]
func (CApi *CourtApi) UpdateCourt(c *gin.Context) {
	var C cms.Court
	err := c.ShouldBindJSON(&C)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CService.UpdateCourt(C)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCourt 用id查询法院
// @Tags Court
// @Summary 用id查询法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Court true "用id查询法院"
// @Success 200 {object} response.Response{data=cms.Court,msg=string} "查询成功"
// @Router /C/findCourt [get]
func (CApi *CourtApi) FindCourt(c *gin.Context) {
	ID := c.Query("ID")
	reC, err := CService.GetCourt(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reC, c)
}

// GetCourtList 分页获取法院列表
// @Tags Court
// @Summary 分页获取法院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CourtSearch true "分页获取法院列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /C/getCourtList [get]
func (CApi *CourtApi) GetCourtList(c *gin.Context) {
	var pageInfo cmsReq.CourtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := CService.GetCourtInfoList(pageInfo)
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

// GetCourtPublic 不需要鉴权的法院接口
// @Tags Court
// @Summary 不需要鉴权的法院接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CourtSearch true "分页获取法院列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /C/getCourtPublic [get]
func (CApi *CourtApi) GetCourtPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    CService.GetCourtPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的法院接口信息",
    }, "获取成功", c)
}
