package alert

import (
	"ops-monitor/alert/consumer"
	"ops-monitor/alert/eval"
	"ops-monitor/alert/probing"
	"ops-monitor/alert/storage"
	"ops-monitor/internal/global"
	"ops-monitor/pkg/ctx"

	"github.com/zeromicro/go-zero/core/logc"
)

var (
	AlertRule      eval.AlertRuleEval
	ProductProbing probing.ProductProbing
	ConsumeProbing probing.ConsumeProbing
	// ProbingService *probing.ProbingService
)

func Initialize(ctx *ctx.Context) {
	// 初始化告警规则消费任务
	consumer.NewInterEvalConsumeWork(ctx).Run()
	// 初始化监控告警的基础配置
	initAlarmConfig(ctx)
	alarmRecoverWaitStore := storage.NewAlarmRecoverStore()

	// 初始化告警规则评估任务
	AlertRule = eval.NewAlertRuleEval(ctx, alarmRecoverWaitStore)
	_ = AlertRule.RePushTask()

	// 初始化拨测任务
	ConsumeProbing = probing.NewProbingConsumerTask(ctx)
	ProductProbing = probing.NewProbingTask(ctx)
	ProductProbing.RePushRule(&ConsumeProbing)
}

func initAlarmConfig(ctx *ctx.Context) {
	get, err := ctx.DB.Setting().Get()
	if err != nil {
		logc.Errorf(ctx.Ctx, err.Error())
		return
	}

	global.Config.Server.AlarmConfig = get.AlarmConfig

	if global.Config.Server.AlarmConfig.RecoverWait == 0 {
		global.Config.Server.AlarmConfig.RecoverWait = 1
	}

	if global.Config.Server.AlarmConfig.GroupInterval == 0 {
		global.Config.Server.AlarmConfig.GroupInterval = 120
	}

	if global.Config.Server.AlarmConfig.GroupWait == 0 {
		global.Config.Server.AlarmConfig.GroupWait = 10
	}
}

//func Initialize2(ctx *ctx.Context) {
//	// 初始化告警规则消费任务
//	consumer.NewInterEvalConsumeWork(ctx).Run()
//
//	// 初始化监控告警的基础配置
//	initAlarmConfig(ctx)
//
//	// 初始化告警规则评估任务
//	AlertRule = eval.NewAlertRuleEval(ctx)
//	AlertRule.RePushTask()
//
//	// 初始化探测服务
//	redisOpt := asynq.RedisClientOpt{
//		Addr:     ctx.Config.Redis.Host,
//		Password: ctx.Config.Redis.Password,
//		DB:       ctx.Config.Redis.DB,
//	}
//
//	var err error
//	ProbingService, err = probing.NewProbingService(ctx, redisOpt)
//	if err != nil {
//		logc.Errorf(ctx.Ctx, "Failed to create probing service: %v", err)
//		return
//	}
//
//	// 启动消费者
//	go func() {
//		if err := ProbingService.StartConsumer(); err != nil {
//			logc.Errorf(ctx.Ctx, "Failed to start probing consumer: %v", err)
//		}
//	}()
//
//	// 重新加载探测规则
//	// rePushProbingRules(ctx)
//}
//
//func rePushProbingRules(ctx *ctx.Context) {
//	var ruleList []models.ProbingRule
//	if err := ctx.DB.DB().Where("enabled = ?", true).Find(&ruleList).Error; err != nil {
//		logc.Errorf(ctx.Ctx, "Failed to get probing rules: %v", err)
//		return
//	}
//
//	for _, rule := range ruleList {
//		if err := ProbingService.StartProducer(rule); err != nil {
//			logc.Errorf(ctx.Ctx, "Failed to start producer for rule %s: %v", rule.RuleId, err)
//		}
//	}
//}
//
//func Shutdown() {
//	if ProbingService != nil {
//		ProbingService.Shutdown()
//	}
//}
