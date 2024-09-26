package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApplicantRouter struct {}

// InitApplicantRouter 初始化 applicant表 路由信息
func (s *ApplicantRouter) InitApplicantRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	applicantRouter := Router.Group("applicant").Use(middleware.OperationRecord())
	applicantRouterWithoutRecord := Router.Group("applicant")
	applicantRouterWithoutAuth := PublicRouter.Group("applicant")
	{
		applicantRouter.POST("createApplicant", applicantApi.CreateApplicant)   // 新建applicant表
		applicantRouter.DELETE("deleteApplicant", applicantApi.DeleteApplicant) // 删除applicant表
		applicantRouter.DELETE("deleteApplicantByIds", applicantApi.DeleteApplicantByIds) // 批量删除applicant表
		applicantRouter.PUT("updateApplicant", applicantApi.UpdateApplicant)    // 更新applicant表
	}
	{
		applicantRouterWithoutRecord.GET("findApplicant", applicantApi.FindApplicant)        // 根据ID获取applicant表
		applicantRouterWithoutRecord.GET("getApplicantList", applicantApi.GetApplicantList)  // 获取applicant表列表
	}
	{
	    applicantRouterWithoutAuth.GET("getApplicantPublic", applicantApi.GetApplicantPublic)  // applicant表开放接口
	}
}
