package cms

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ApplicantRouter
	ContactsRouter
	PolicyholderRouter
	CourtRouter
	MeLetterRouter
}

var (
	applicantApi    = api.ApiGroupApp.CmsApiGroup.ApplicantApi
	contactsApi     = api.ApiGroupApp.CmsApiGroup.ContactsApi
	policyholderApi = api.ApiGroupApp.CmsApiGroup.PolicyholderApi
	CApi            = api.ApiGroupApp.CmsApiGroup.CourtApi
	meLetterApi     = api.ApiGroupApp.CmsApiGroup.MeLetterApi
)
