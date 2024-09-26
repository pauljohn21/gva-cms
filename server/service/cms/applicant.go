package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type ApplicantService struct {}
// CreateApplicant 创建applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService) CreateApplicant(applicant *cms.Applicant) (err error) {
	err = global.GVA_DB.Create(applicant).Error
	return err
}

// DeleteApplicant 删除applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService)DeleteApplicant(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.Applicant{},"id = ?",ID).Error
	return err
}

// DeleteApplicantByIds 批量删除applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService)DeleteApplicantByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.Applicant{},"id in ?",IDs).Error
	return err
}

// UpdateApplicant 更新applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService)UpdateApplicant(applicant cms.Applicant) (err error) {
	err = global.GVA_DB.Model(&cms.Applicant{}).Where("id = ?",applicant.ID).Updates(&applicant).Error
	return err
}

// GetApplicant 根据ID获取applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService)GetApplicant(ID string) (applicant cms.Applicant, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&applicant).Error
	return
}

// GetApplicantInfoList 分页获取applicant表记录
// Author [yourname](https://github.com/yourname)
func (applicantService *ApplicantService)GetApplicantInfoList(info cmsReq.ApplicantSearch) (list []cms.Applicant, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Applicant{})
    var applicants []cms.Applicant
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

	err = db.Find(&applicants).Error
	return  applicants, total, err
}
func (applicantService *ApplicantService)GetApplicantPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
