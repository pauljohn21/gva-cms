package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
    cmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/cms/request"
)

type ContactsService struct {}
// CreateContacts 创建联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService) CreateContacts(contacts *cms.Contacts) (err error) {
	err = global.GVA_DB.Create(contacts).Error
	return err
}

// DeleteContacts 删除联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService)DeleteContacts(ID string) (err error) {
	err = global.GVA_DB.Delete(&cms.Contacts{},"id = ?",ID).Error
	return err
}

// DeleteContactsByIds 批量删除联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService)DeleteContactsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cms.Contacts{},"id in ?",IDs).Error
	return err
}

// UpdateContacts 更新联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService)UpdateContacts(contacts cms.Contacts) (err error) {
	err = global.GVA_DB.Model(&cms.Contacts{}).Where("id = ?",contacts.ID).Updates(&contacts).Error
	return err
}

// GetContacts 根据ID获取联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService)GetContacts(ID string) (contacts cms.Contacts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contacts).Error
	return
}

// GetContactsInfoList 分页获取联系人记录
// Author [yourname](https://github.com/yourname)
func (contactsService *ContactsService)GetContactsInfoList(info cmsReq.ContactsSearch) (list []cms.Contacts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Contacts{})
    var contactss []cms.Contacts
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

	err = db.Find(&contactss).Error
	return  contactss, total, err
}
func (contactsService *ContactsService)GetContactsPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
