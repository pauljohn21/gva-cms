// 自动生成模板Contacts
package cms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 联系人 结构体  Contacts
type Contacts struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;size:191;"`                   // name字段
	Phone     *int   `json:"phone" form:"phone" gorm:"column:phone;comment:电话;size:191;"`              // 电话
	Email     string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:191;"`              // 邮箱
	Addres    string `json:"addres" form:"addres" gorm:"column:addres;comment:地址;size:191;"`           // 地址
	CreatedBy *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"` // 创建者
	UpdatedBy *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"` // 更新者
	DeletedBy *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"` // 删除者
	Mobile    *int   `json:"mobile" form:"mobile" gorm:"column:mobile;comment:;size:19;"`              // mobile字段
	Addr      string `json:"addr" form:"addr" gorm:"column:addr;comment:;size:191;"`                   // addr字段
}

// TableName 联系人 Contacts自定义表名 contacts
func (Contacts) TableName() string {
	return "contacts"
}
