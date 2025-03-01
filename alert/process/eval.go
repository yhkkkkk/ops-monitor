package process

import (
	"context"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"

	"github.com/zeromicro/go-zero/core/logc"
)

// EvalCondition 评估告警条件
func EvalCondition(ec models.EvalCondition) bool {
	switch ec.Operator {
	case ">":
		if ec.QueryValue > ec.ExpectedValue {
			return true
		}
	case ">=":
		if ec.QueryValue >= ec.ExpectedValue {
			return true
		}
	case "<":
		if ec.QueryValue < ec.ExpectedValue {
			return true
		}
	case "<=":
		if ec.QueryValue <= ec.ExpectedValue {
			return true
		}
	case "==":
		if ec.QueryValue == ec.ExpectedValue {
			return true
		}
	case "!=":
		if ec.QueryValue != ec.ExpectedValue {
			return true
		}
	default:
		logc.Error(context.Background(), "无效的评估条件", ec.Type, ec.Operator, ec.ExpectedValue)
	}
	return false
}

func SaveAlertEvent(ctx *ctx.Context, event models.AlertCurEvent) {
	ok := ctx.DB.Rule().GetRuleIsExist(event.RuleId)
	if ok {
		SaveEventCache(ctx, event)
	}
}
