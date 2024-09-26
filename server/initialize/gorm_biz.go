package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(cms.Applicant{}, cms.Contacts{}, cms.Policyholder{}, cms.Court{}, cms.MeLetter{})
	if err != nil {
		return err
	}
	return nil
}
