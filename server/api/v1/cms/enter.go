package cms

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ApplicantApi
	ContactsApi
	PolicyholderApi
	CourtApi
	MeLetterApi
}

var (
	applicantService    = service.ServiceGroupApp.CmsServiceGroup.ApplicantService
	contactsService     = service.ServiceGroupApp.CmsServiceGroup.ContactsService
	policyholderService = service.ServiceGroupApp.CmsServiceGroup.PolicyholderService
	CService            = service.ServiceGroupApp.CmsServiceGroup.CourtService
	meLetterService     = service.ServiceGroupApp.CmsServiceGroup.MeLetterService
)
