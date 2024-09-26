package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PolicyholderRouter struct {}

// InitPolicyholderRouter 初始化 投保人 路由信息
func (s *PolicyholderRouter) InitPolicyholderRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	policyholderRouter := Router.Group("policyholder").Use(middleware.OperationRecord())
	policyholderRouterWithoutRecord := Router.Group("policyholder")
	policyholderRouterWithoutAuth := PublicRouter.Group("policyholder")
	{
		policyholderRouter.POST("createPolicyholder", policyholderApi.CreatePolicyholder)   // 新建投保人
		policyholderRouter.DELETE("deletePolicyholder", policyholderApi.DeletePolicyholder) // 删除投保人
		policyholderRouter.DELETE("deletePolicyholderByIds", policyholderApi.DeletePolicyholderByIds) // 批量删除投保人
		policyholderRouter.PUT("updatePolicyholder", policyholderApi.UpdatePolicyholder)    // 更新投保人
	}
	{
		policyholderRouterWithoutRecord.GET("findPolicyholder", policyholderApi.FindPolicyholder)        // 根据ID获取投保人
		policyholderRouterWithoutRecord.GET("getPolicyholderList", policyholderApi.GetPolicyholderList)  // 获取投保人列表
	}
	{
	    policyholderRouterWithoutAuth.GET("getPolicyholderPublic", policyholderApi.GetPolicyholderPublic)  // 投保人开放接口
	}
}
