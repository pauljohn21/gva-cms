// 自动生成模板Policyholder
package cms
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 投保人 结构体  Policyholder
type Policyholder struct {
    global.GVA_MODEL
    Company  string `json:"company" form:"company" gorm:"column:company;comment:;size:191;"`  //company字段 
    Code  string `json:"code" form:"code" gorm:"column:code;comment:统一社会信用代码;size:191;"`  //统一社会信用代码 
    Bank  string `json:"bank" form:"bank" gorm:"column:bank;comment:开户行;size:191;"`  //开户行 
    Account  string `json:"account" form:"account" gorm:"column:account;comment:对公账户;size:191;"`  //对公账户 
    Addrs  string `json:"addrs" form:"addrs" gorm:"column:addrs;comment:地址;size:191;"`  //地址 
    Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:电话;size:191;"`  //电话 
    Status  string `json:"status" form:"status" gorm:"column:status;comment:状态;size:191;"`  //状态 
    CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`  //创建者 
    UpdatedBy  *int `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`  //更新者 
    DeletedBy  *int `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`  //删除者 
    Signature  string `json:"signature" form:"signature" gorm:"column:signature;comment:;size:191;"`  //signature字段 
}


// TableName 投保人 Policyholder自定义表名 policyholder
func (Policyholder) TableName() string {
    return "policyholder"
}

