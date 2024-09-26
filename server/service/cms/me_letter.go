package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type MeLetterService struct {}
// CreateMeLetter 创建我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService) CreateMeLetter(meLetter *cms.MeLetter) (err error) {
	err = global.GVA_DB.Create(meLetter).Error
	return err
}

// DeleteMeLetter 删除我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService)DeleteMeLetter(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.MeLetter{},"id = ?",ID).Error
	return err
}

// DeleteMeLetterByIds 批量删除我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService)DeleteMeLetterByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.MeLetter{},"id in ?",IDs).Error
	return err
}

// UpdateMeLetter 更新我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService)UpdateMeLetter(meLetter cms.MeLetter) (err error) {
	err = global.GVA_DB.Model(&cms.MeLetter{}).Where("id = ?",meLetter.ID).Updates(&meLetter).Error
	return err
}

// GetMeLetter 根据ID获取我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService)GetMeLetter(ID string) (meLetter cms.MeLetter, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&meLetter).Error
	return
}

// GetMeLetterInfoList 分页获取我的保函记录
// Author [yourname](https://github.com/yourname)
func (meLetterService *MeLetterService)GetMeLetterInfoList(info cmsReq.MeLetterSearch) (list []cms.MeLetter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.MeLetter{})
    var meLetters []cms.MeLetter
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

	err = db.Find(&meLetters).Error
	return  meLetters, total, err
}
func (meLetterService *MeLetterService)GetMeLetterPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
