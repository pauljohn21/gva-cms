// 自动生成模板Applicant
package cms
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// applicant表 结构体  Applicant
type Applicant struct {
    global.GVA_MODEL
    Company  string `json:"company" form:"company" gorm:"column:company;comment:公司名;size:191;"`  //公司名 
    Code  string `json:"code" form:"code" gorm:"column:code;comment:统一社会信用代码;size:191;"`  //统一社会信用代码 
    CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`  //创建者 
    UpdatedBy  *int `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`  //更新者 
    DeletedBy  *int `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`  //删除者 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`  //name字段 
    CardId  string `json:"cardId" form:"cardId" gorm:"column:card_id;comment:;size:191;"`  //cardId字段 
    Amount  string `json:"amount" form:"amount" gorm:"column:amount;comment:;size:191;"`  //amount字段 
}


// TableName applicant表 Applicant自定义表名 applicant
func (Applicant) TableName() string {
    return "applicant"
}

