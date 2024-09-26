// 自动生成模板Court
package cms
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 法院 结构体  Court
type Court struct {
    global.GVA_MODEL
    Name  string `json:"name" form:"name" gorm:"column:name;comment:法院名;" binding:"required"`  //名称 
    Addr  string `json:"addr" form:"addr" gorm:"column:addr;comment:地址;" binding:"required"`  //地址 
}


// TableName 法院 Court自定义表名 court
func (Court) TableName() string {
    return "court"
}

