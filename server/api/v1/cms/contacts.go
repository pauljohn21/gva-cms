package cms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ContactsApi struct {}



// CreateContacts 创建联系人
// @Tags Contacts
// @Summary 创建联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Contacts true "创建联系人"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /contacts/createContacts [post]
func (contactsApi *ContactsApi) CreateContacts(c *gin.Context) {
	var contacts cms.Contacts
	err := c.ShouldBindJSON(&contacts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = contactsService.CreateContacts(&contacts)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteContacts 删除联系人
// @Tags Contacts
// @Summary 删除联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Contacts true "删除联系人"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /contacts/deleteContacts [delete]
func (contactsApi *ContactsApi) DeleteContacts(c *gin.Context) {
	ID := c.Query("ID")
	err := contactsService.DeleteContacts(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteContactsByIds 批量删除联系人
// @Tags Contacts
// @Summary 批量删除联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /contacts/deleteContactsByIds [delete]
func (contactsApi *ContactsApi) DeleteContactsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := contactsService.DeleteContactsByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateContacts 更新联系人
// @Tags Contacts
// @Summary 更新联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Contacts true "更新联系人"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /contacts/updateContacts [put]
func (contactsApi *ContactsApi) UpdateContacts(c *gin.Context) {
	var contacts cms.Contacts
	err := c.ShouldBindJSON(&contacts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = contactsService.UpdateContacts(contacts)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindContacts 用id查询联系人
// @Tags Contacts
// @Summary 用id查询联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Contacts true "用id查询联系人"
// @Success 200 {object} response.Response{data=cms.Contacts,msg=string} "查询成功"
// @Router /contacts/findContacts [get]
func (contactsApi *ContactsApi) FindContacts(c *gin.Context) {
	ID := c.Query("ID")
	recontacts, err := contactsService.GetContacts(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recontacts, c)
}

// GetContactsList 分页获取联系人列表
// @Tags Contacts
// @Summary 分页获取联系人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ContactsSearch true "分页获取联系人列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /contacts/getContactsList [get]
func (contactsApi *ContactsApi) GetContactsList(c *gin.Context) {
	var pageInfo cmsReq.ContactsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := contactsService.GetContactsInfoList(pageInfo)
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

// GetContactsPublic 不需要鉴权的联系人接口
// @Tags Contacts
// @Summary 不需要鉴权的联系人接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ContactsSearch true "分页获取联系人列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contacts/getContactsPublic [get]
func (contactsApi *ContactsApi) GetContactsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    contactsService.GetContactsPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的联系人接口信息",
    }, "获取成功", c)
}
