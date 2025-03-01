package probing

import (
	"fmt"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"

	"github.com/zeromicro/go-zero/core/logc"
	"golang.org/x/net/context"
)

// EvalStrategy 日志评估条件
type EvalStrategy struct {
	// 运算
	Operator string `json:"operator"`
	// 查询值
	QueryValue float64 `json:"queryValue"`
	// 预期值
	ExpectedValue float64 `json:"value"`
}

// EvalCondition 评估告警条件
func EvalCondition(ec EvalStrategy) bool {
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
		logc.Errorf(context.Background(), fmt.Sprintf("无效的评估条件", ec.Operator, ec.ExpectedValue))
	}
	return false
}

func SaveProbingEndpointEvent(event models.ProbingEvent) {
	firingKey := event.GetFiringAlertCacheKey()
	cache := ctx.DO().Redis.Event()
	resFiring, _ := cache.GetPECache(firingKey)
	event.FirstTriggerTime = cache.GetPEFirstTime(firingKey)
	event.LastEvalTime = cache.GetPELastEvalTime(firingKey)
	event.LastSendTime = resFiring.LastSendTime
	cache.SetPECache(event, 0)
}

func SetProbingValueMap(key string, m map[string]any) error {
	for k, v := range m {
		ctx.DO().Redis.Redis().HSet(key, k, v)
	}
	return nil
}

func GetProbingValueMap(key string) map[string]string {
	result := ctx.DO().Redis.Redis().HGetAll(key).Val()
	return result
}
