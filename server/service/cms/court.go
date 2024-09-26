package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type CourtService struct {}
// CreateCourt 创建法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService) CreateCourt(C *cms.Court) (err error) {
	err = global.GVA_DB.Create(C).Error
	return err
}

// DeleteCourt 删除法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService)DeleteCourt(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.Court{},"id = ?",ID).Error
	return err
}

// DeleteCourtByIds 批量删除法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService)DeleteCourtByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.Court{},"id in ?",IDs).Error
	return err
}

// UpdateCourt 更新法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService)UpdateCourt(C cms.Court) (err error) {
	err = global.GVA_DB.Model(&cms.Court{}).Where("id = ?",C.ID).Updates(&C).Error
	return err
}

// GetCourt 根据ID获取法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService)GetCourt(ID string) (C cms.Court, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&C).Error
	return
}

// GetCourtInfoList 分页获取法院记录
// Author [yourname](https://github.com/yourname)
func (CService *CourtService)GetCourtInfoList(info cmsReq.CourtSearch) (list []cms.Court, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Court{})
    var Cs []cms.Court
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&Cs).Error
	return  Cs, total, err
}
func (CService *CourtService)GetCourtPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
