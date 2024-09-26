package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourtRouter struct {}

// InitCourtRouter 初始化 法院 路由信息
func (s *CourtRouter) InitCourtRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	CRouter := Router.Group("C").Use(middleware.OperationRecord())
	CRouterWithoutRecord := Router.Group("C")
	CRouterWithoutAuth := PublicRouter.Group("C")
	{
		CRouter.POST("createCourt", CApi.CreateCourt)   // 新建法院
		CRouter.DELETE("deleteCourt", CApi.DeleteCourt) // 删除法院
		CRouter.DELETE("deleteCourtByIds", CApi.DeleteCourtByIds) // 批量删除法院
		CRouter.PUT("updateCourt", CApi.UpdateCourt)    // 更新法院
	}
	{
		CRouterWithoutRecord.GET("findCourt", CApi.FindCourt)        // 根据ID获取法院
		CRouterWithoutRecord.GET("getCourtList", CApi.GetCourtList)  // 获取法院列表
	}
	{
	    CRouterWithoutAuth.GET("getCourtPublic", CApi.GetCourtPublic)  // 法院开放接口
	}
}
