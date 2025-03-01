package v1

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	v1 := engine.Group("api")
	{
		system := v1.Group("system")
		{
			DashboardInfo.API(v1)
			system.POST("register", Auth.Register)
			system.POST("login", Auth.Login)
			system.GET("checkUser", Auth.CheckUser)
			system.GET("checkNoticeStatus", Notice.Check)
			system.GET("userInfo", Auth.Get)
		}

		ops := v1.Group("ops")
		{
			Auth.API(ops)
			Permissions.API(ops)
			AlertEvent.API(ops)
			Role.API(ops)
			Dashboard.API(ops)
			Datasource.API(ops)
			RuleGroup.API(ops)
			Rule.API(ops)
			Silence.API(ops)
			Notice.API(ops)
			NoticeTemplate.API(ops)
			Tenant.API(ops)
			RuleTmplGroup.API(ops)
			RuleTmpl.API(ops)
			Duty.API(ops)
			DutyCalendar.API(ops)
			AuditLog.API(ops)
			ClientApi.API(ops)
			Setting.API(ops)
			Subscribe.API(ops)
			Probing.API(ops)
		}
	}
}
