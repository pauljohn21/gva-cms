package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type PolicyholderService struct {}
// CreatePolicyholder 创建投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService) CreatePolicyholder(policyholder *cms.Policyholder) (err error) {
	err = global.GVA_DB.Create(policyholder).Error
	return err
}

// DeletePolicyholder 删除投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService)DeletePolicyholder(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.Policyholder{},"id = ?",ID).Error
	return err
}

// DeletePolicyholderByIds 批量删除投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService)DeletePolicyholderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.Policyholder{},"id in ?",IDs).Error
	return err
}

// UpdatePolicyholder 更新投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService)UpdatePolicyholder(policyholder cms.Policyholder) (err error) {
	err = global.GVA_DB.Model(&cms.Policyholder{}).Where("id = ?",policyholder.ID).Updates(&policyholder).Error
	return err
}

// GetPolicyholder 根据ID获取投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService)GetPolicyholder(ID string) (policyholder cms.Policyholder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&policyholder).Error
	return
}

// GetPolicyholderInfoList 分页获取投保人记录
// Author [yourname](https://github.com/yourname)
func (policyholderService *PolicyholderService)GetPolicyholderInfoList(info cmsReq.PolicyholderSearch) (list []cms.Policyholder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Policyholder{})
    var policyholders []cms.Policyholder
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

	err = db.Find(&policyholders).Error
	return  policyholders, total, err
}
func (policyholderService *PolicyholderService)GetPolicyholderPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
