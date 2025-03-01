package models

type UserPermissions struct {
	Key string `json:"key"`
	API string `json:"api"`
}

func PermissionsInfo() map[string]UserPermissions {
	return map[string]UserPermissions{
		"ruleSearch": {
			Key: "搜索告警规则",
			API: "/api/ops/rule/ruleSearch",
		},
		"calendarCreate": {
			Key: "发布日历表",
			API: "/api/ops/calendar/calendarCreate",
		},
		"calendarSearch": {
			Key: "搜索日历表",
			API: "/api/ops/calendar/calendarSearch",
		},
		"calendarUpdate": {
			Key: "更新日历表",
			API: "/api/ops/calendar/calendarUpdate",
		},
		"createDashboard": {
			Key: "创建仪表盘",
			API: "/api/ops/dashboard/createDashboard",
		},
		"createTenant": {
			Key: "创建租户",
			API: "/api/ops/tenant/createTenant",
		},
		"curEvent": {
			Key: "查看当前告警事件",
			API: "/api/ops/event/curEvent",
		},
		"dataSourceCreate": {
			Key: "创建数据源",
			API: "/api/ops/datasource/dataSourceCreate",
		},
		"dataSourceDelete": {
			Key: "删除数据源",
			API: "/api/ops/datasource/dataSourceDelete",
		},
		"dataSourceGet": {
			Key: "获取数据源",
			API: "/api/ops/datasource/dataSourceGet",
		},
		"dataSourceList": {
			Key: "查看数据源",
			API: "/api/ops/datasource/dataSourceList",
		},
		"dataSourceSearch": {
			Key: "搜索数据源",
			API: "/api/ops/datasource/dataSourceSearch",
		},
		"dataSourceUpdate": {
			Key: "更新数据源",
			API: "/api/ops/datasource/dataSourceUpdate",
		},
		"deleteDashboard": {
			Key: "删除仪表盘",
			API: "/api/ops/dashboard/deleteDashboard",
		},
		"deleteTenant": {
			Key: "删除租户",
			API: "/api/ops/tenant/deleteTenant",
		},
		"dutyManageCreate": {
			Key: "创建值班表",
			API: "/api/ops/dutyManage/dutyManageCreate",
		},
		"dutyManageDelete": {
			Key: "更新值班表",
			API: "/api/ops/dutyManage/dutyManageDelete",
		},
		"dutyManageList": {
			Key: "查看值班表",
			API: "/api/ops/dutyManage/dutyManageList",
		},
		"dutyManageSearch": {
			Key: "搜索值班表",
			API: "/api/ops/dutyManage/dutyManageSearch",
		},
		"dutyManageUpdate": {
			Key: "更新值班表",
			API: "/api/ops/dutyManage/dutyManageUpdate",
		},
		"getDashboard": {
			Key: "获取仪表盘",
			API: "/api/ops/dashboard/getDashboard",
		},
		"getTenantList": {
			Key: "查看租户",
			API: "/api/ops/tenant/getTenantList",
		},
		"hisEvent": {
			Key: "查看历史告警",
			API: "/api/ops/event/hisEvent",
		},
		"listDashboard": {
			Key: "查看仪表盘",
			API: "/api/ops/dashboard/listDashboard",
		},
		"noticeCreate": {
			Key: "创建通知对象",
			API: "/api/ops/notice/noticeCreate",
		},
		"noticeDelete": {
			Key: "删除通知对象",
			API: "/api/ops/notice/noticeDelete",
		},
		"noticeList": {
			Key: "查看通知对象",
			API: "/api/ops/notice/noticeList",
		},
		"noticeSearch": {
			Key: "搜索通知对象",
			API: "/api/ops/notice/noticeSearch",
		},
		"noticeTemplateCreate": {
			Key: "创建通知模版",
			API: "/api/ops/noticeTemplate/noticeTemplateCreate",
		},
		"noticeTemplateDelete": {
			Key: "删除通知模版",
			API: "/api/ops/noticeTemplate/noticeTemplateDelete",
		},
		"noticeTemplateList": {
			Key: "查看通知模版",
			API: "/api/ops/noticeTemplate/noticeTemplateList",
		},
		"noticeTemplateUpdate": {
			Key: "更新通知模版",
			API: "/api/ops/noticeTemplate/noticeTemplateUpdate",
		},
		"noticeUpdate": {
			Key: "更新通知对象",
			API: "/api/ops/notice/noticeUpdate",
		},
		"permsList": {
			Key: "查看用户权限",
			API: "/api/ops/permissions/permsList",
		},
		"register": {
			Key: "用户注册",
			API: "/api/system/register",
		},
		"roleCreate": {
			Key: "创建用户角色",
			API: "/api/ops/role/roleCreate",
		},
		"roleDelete": {
			Key: "删除用户角色",
			API: "/api/ops/role/roleDelete",
		},
		"roleList": {
			Key: "查看用户角色",
			API: "/api/ops/role/roleList",
		},
		"roleUpdate": {
			Key: "更新用户角色",
			API: "/api/ops/role/roleUpdate",
		},
		"ruleCreate": {
			Key: "创建告警规则",
			API: "/api/ops/rule/ruleCreate",
		},
		"ruleDelete": {
			Key: "删除告警规则",
			API: "/api/ops/rule/ruleDelete",
		},
		"ruleGroupCreate": {
			Key: "创建告警规则组",
			API: "/api/ops/ruleGroup/ruleGroupCreate",
		},
		"ruleGroupDelete": {
			Key: "删除告警规则组",
			API: "/api/ops/ruleGroup/ruleGroupDelete",
		},
		"ruleGroupList": {
			Key: "查看告警规则组",
			API: "/api/ops/ruleGroup/ruleGroupList",
		},
		"ruleGroupUpdate": {
			Key: "更新告警规则组",
			API: "/api/ops/ruleGroup/ruleGroupUpdate",
		},
		"ruleList": {
			Key: "查看告警规则",
			API: "/api/ops/rule/ruleList",
		},
		"ruleTmplCreate": {
			Key: "创建规则模版",
			API: "/api/ops/ruleTmpl/ruleTmplCreate",
		},
		"ruleTmplDelete": {
			Key: "删除规则模版",
			API: "/api/ops/ruleTmpl/ruleTmplDelete",
		},
		"ruleTmplGroupCreate": {
			Key: "创建规则模版组",
			API: "/api/ops/ruleTmplGroup/ruleTmplGroupCreate",
		},
		"ruleTmplGroupDelete": {
			Key: "删除规则模版组",
			API: "/api/ops/ruleTmplGroup/ruleTmplGroupDelete",
		},
		"ruleTmplGroupList": {
			Key: "查看规则模版组",
			API: "/api/ops/ruleTmplGroup/ruleTmplGroupList",
		},
		"ruleTmplList": {
			Key: "查看规则模版",
			API: "/api/ops/ruleTmpl/ruleTmplList",
		},
		"ruleUpdate": {
			Key: "更新告警规则",
			API: "/api/ops/rule/ruleUpdate",
		},
		"searchDashboard": {
			Key: "搜索仪表盘",
			API: "/api/ops/dashboard/searchDashboard",
		},
		"searchDutyUser": {
			Key: "搜索值班用户",
			API: "/api/ops/user/searchDutyUser",
		},
		"silenceCreate": {
			Key: "创建静默规则",
			API: "/api/ops/silence/silenceCreate",
		},
		"silenceDelete": {
			Key: "删除静默规则",
			API: "/api/ops/silence/silenceDelete",
		},
		"silenceList": {
			Key: "查看静默规则",
			API: "/api/ops/silence/silenceList",
		},
		"silenceUpdate": {
			Key: "更新静默规则",
			API: "/api/ops/silence/silenceUpdate",
		},
		"updateDashboard": {
			Key: "更新仪表盘",
			API: "/api/ops/dashboard/updateDashboard",
		},
		"updateTenant": {
			Key: "更新租户信息",
			API: "/api/ops/tenant/updateTenant",
		},
		"userChangePass": {
			Key: "修改用户密码",
			API: "/api/ops/user/userChangePass",
		},
		"userDelete": {
			Key: "删除用户",
			API: "/api/ops/user/userDelete",
		},
		"userList": {
			Key: "查看用户列表",
			API: "/api/ops/user/userList",
		},
		"userUpdate": {
			Key: "更新用户信息",
			API: "/api/ops/user/userUpdate",
		},
		"getJaegerService": {
			Key: "获取Jaeger服务列表",
			API: "/api/ops/c/getJaegerService",
		},
		"searchUser": {
			Key: "搜索用户",
			API: "/api/ops/user/searchUser",
		},
		"searchNoticeTmpl": {
			Key: "搜索通知模版",
			API: "/api/ops/noticeTemplate/searchNoticeTmpl",
		},
		"saveSystemSetting": {
			Key: "编辑系统配置",
			API: "/api/ops/setting/saveSystemSetting",
		},
		"getSystemSetting": {
			Key: "获取系统配置",
			API: "/api/ops/setting/getSystemSetting",
		},
		"promQuery": {
			Key: "Prometheus指标查询",
			API: "/api/ops/datasource/promQuery",
		},
		"getTenant": {
			Key: "获取租户详细信息",
			API: "/api/ops/tenant/getTenant",
		},
		"addUsersToTenant": {
			Key: "向租户添加成员",
			API: "/api/ops/tenant/addUsersToTenant",
		},
		"delUsersOfTenant": {
			Key: "删除租户成员",
			API: "/api/ops/tenant/delUsersOfTenant",
		},
		"getUsersForTenant": {
			Key: "获取租户成员列表",
			API: "/api/ops/tenant/getUsersForTenant",
		},
		"changeTenantUserRole": {
			Key: "修改租户成员角色",
			API: "/api/ops/tenant/changeTenantUserRole",
		},
		"createProbing": {
			Key: "创建拨测规则",
			API: "/api/ops/probing/createProbing",
		},
		"updateProbing": {
			Key: "更新拨测规则",
			API: "/api/ops/probing/updateProbing",
		},
		"deleteProbing": {
			Key: "删除拨测规则",
			API: "/api/ops/probing/deleteProbing",
		},
		"listProbing": {
			Key: "获取拨测规则列表",
			API: "/api/ops/probing/listProbing",
		},
		"searchProbing": {
			Key: "获取拨测规则信息",
			API: "/api/ops/probing/searchProbing",
		},
		"onceProbing": {
			Key: "一次性拨测任务",
			API: "/api/ops/probing/onceProbing",
		},
		"listFolder": {
			Key: "获取仪表盘目录列表",
			API: "/api/ops/dashboard/listFolder",
		},
		"getFolder": {
			Key: "获取仪表盘目录详情",
			API: "/api/ops/dashboard/getFolder",
		},
		"createFolder": {
			Key: "创建仪表盘目录",
			API: "/api/ops/dashboard/createFolder",
		},
		"updateFolder": {
			Key: "删除仪表盘目录",
			API: "/api/ops/dashboard/updateFolder",
		},
		"deleteFolder": {
			Key: "删除仪表盘目录",
			API: "/api/ops/dashboard/deleteFolder",
		},
		"listGrafanaDashboards": {
			Key: "获取仪表盘信息",
			API: "/api/ops/dashboard/listGrafanaDashboards",
		},
		"getDashboardFullUrl": {
			Key: "获取仪表盘完整URL",
			API: "/api/ops/dashboard/getDashboardFullUrl",
		},
		"createSubscribe": {
			Key: "创建告警订阅",
			API: "/api/ops/subscribe/createSubscribe",
		},
		"deleteSubscribe": {
			Key: "删除告警订阅",
			API: "/api/ops/subscribe/deleteSubscribe",
		},
		"listSubscribe": {
			Key: "获取告警订阅",
			API: "/api/ops/subscribe/listSubscribe",
		},
		"getSubscribe": {
			Key: "搜索告警订阅",
			API: "/api/ops/subscribe/getSubscribe",
		},
		"noticeRecordList": {
			Key: "获取通知记录列表",
			API: "/api/ops/notice/noticeRecordList",
		},
		"noticeRecordMetric": {
			Key: "获取通知记录指标",
			API: "/api/ops/notice/noticeRecordMetric",
		},
		"dataSourcePing": {
			Key: "数据源连接测试",
			API: "/api/ops/datasource/dataSourcePing",
		},
	}
}
