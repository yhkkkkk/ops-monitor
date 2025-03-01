package v1

import (
	"ops-monitor/api"
)

var (
	Notice         = api.ApiGroupApp.NoticeController
	Silence        = api.ApiGroupApp.SilenceController
	Datasource     = api.ApiGroupApp.DatasourceController
	Duty           = api.ApiGroupApp.DutyController
	DutyCalendar   = api.ApiGroupApp.DutyCalendarController
	Rule           = api.ApiGroupApp.RuleController
	Auth           = api.ApiGroupApp.UserController
	AlertEvent     = api.ApiGroupApp.AlertEventController
	Role           = api.ApiGroupApp.UserRoleController
	Permissions    = api.ApiGroupApp.UserPermissionsController
	NoticeTemplate = api.ApiGroupApp.NoticeTemplateController
	RuleGroup      = api.ApiGroupApp.RuleGroupController
	RuleTmplGroup  = api.ApiGroupApp.RuleTmplGroupController
	RuleTmpl       = api.ApiGroupApp.RuleTmplController
	DashboardInfo  = api.ApiGroupApp.DashboardInfoController
	Tenant         = api.ApiGroupApp.TenantController
	Dashboard      = api.ApiGroupApp.DashboardController
	AuditLog       = api.ApiGroupApp.AuditLogController
	ClientApi      = api.ApiGroupApp.ClientController
	Setting        = api.ApiGroupApp.SettingsController
	Subscribe      = api.ApiGroupApp.SubscribeController
	Probing        = api.ApiGroupApp.ProbingController
)
