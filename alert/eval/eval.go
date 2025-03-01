package eval

import (
	"context"
	"fmt"
	"ops-monitor/alert/process"
	"ops-monitor/alert/storage"
	"ops-monitor/internal/global"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
	"ops-monitor/pkg/provider"
	"ops-monitor/pkg/tools"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"golang.org/x/sync/errgroup"
)

// AlertRuleEval 告警规则评估
type AlertRuleEval interface {
	Submit(rule models.AlertRule) error
	Stop(ruleId string)
	Eval(ctx context.Context, rule models.AlertRule)
	Recover(rule models.AlertRule, curKeys []string) error
	GC(rule models.AlertRule, curFiringKeys, curPendingKeys []string)
	RePushTask() error
}

// AlertRule 告警规则
type AlertRule struct {
	ctx                   *ctx.Context
	watchCtxMap           sync.Map
	alarmRecoverWaitStore storage.AlarmRecoverWaitStore
}

// EvalResult 评估结果结构
type EvalResult struct {
	FiringKeys  []string
	PendingKeys []string
	Error       error
}

func NewAlertRuleEval(ctx *ctx.Context, alarmRecoverWaitStore storage.AlarmRecoverWaitStore) AlertRuleEval {
	return &AlertRule{
		ctx:                   ctx,
		alarmRecoverWaitStore: alarmRecoverWaitStore,
	}
}

// Submit 提交新的告警规则
func (t *AlertRule) Submit(rule models.AlertRule) error {
	if !t.isRuleEnabled(rule.RuleId) {
		return fmt.Errorf("rule %s is disabled", rule.RuleId)
	}

	c, cancel := context.WithCancel(context.Background())

	// 检查是否存在旧的 cancel 函数
	if oldCancel, loaded := t.watchCtxMap.LoadOrStore(rule.RuleId, cancel); loaded {
		oldCancel.(context.CancelFunc)() // 如果存在，先取消旧的
		t.watchCtxMap.Store(rule.RuleId, cancel)
	}

	go t.Eval(c, rule)
	return nil
}

// Stop 停止指定规则的评估
func (t *AlertRule) Stop(ruleId string) {
	if cancel, ok := t.watchCtxMap.LoadAndDelete(ruleId); ok {
		cancel.(context.CancelFunc)()
	}
}

// Eval 评估告警规则
func (t *AlertRule) Eval(ctx context.Context, rule models.AlertRule) {
	timer := time.NewTicker(time.Second * time.Duration(rule.EvalInterval))
	defer func() {
		timer.Stop()
		if r := recover(); r != nil {
			logc.Error(t.ctx.Ctx, fmt.Sprintf("Recovered from rule eval panic: %v, rule: %s", r, rule.RuleId))
		}
	}()

	for {
		select {
		case <-timer.C:
			if !t.isRuleEnabled(rule.RuleId) {
				return
			}

			result := t.evaluateRule(rule)
			if result.Error != nil {
				logc.Error(t.ctx.Ctx, fmt.Sprintf("Rule evaluation failed: %v", result.Error))
				continue
			}

			if err := t.Recover(rule, result.FiringKeys); err != nil {
				logc.Error(t.ctx.Ctx, fmt.Sprintf("Recovery process failed: %v", err))
			}

			t.GC(rule, result.FiringKeys, result.PendingKeys)

		case <-ctx.Done():
			logc.Info(t.ctx.Ctx, fmt.Sprintf("Stopping watch routine for rule: %s", rule.RuleId))
			return
		}
		// timer.Reset(time.Second * time.Duration(rule.EvalInterval))
	}
}

// evaluateRule 评估单个规则
func (t *AlertRule) evaluateRule(rule models.AlertRule) EvalResult {
	var result EvalResult

	for _, dsId := range rule.DatasourceIdList {
		instance, err := t.ctx.DB.Datasource().GetInstance(dsId)
		if err != nil {
			logc.Error(t.ctx.Ctx, fmt.Sprintf("Failed to get datasource instance: %v", err))
			continue
		}

		if !provider.CheckDatasourceHealth(instance) {
			continue
		}

		var firingKeys, pendingKeys []string
		switch rule.DatasourceType {
		case "Prometheus", "VictoriaMetrics":
			firingKeys, pendingKeys = metrics(t.ctx, dsId, instance.Type, rule)
		case "AliCloudSLS", "ElasticSearch":
			firingKeys = logs(t.ctx, dsId, instance.Type, rule)
		case "Jaeger":
			firingKeys = traces(t.ctx, dsId, instance.Type, rule)
		}

		result.FiringKeys = append(result.FiringKeys, firingKeys...)
		result.PendingKeys = append(result.PendingKeys, pendingKeys...)
	}

	return result
}

// Recover 处理告警恢复
func (t *AlertRule) Recover(rule models.AlertRule, curKeys []string) error {
	firingKeys, err := ctx.Redis.Rule().GetAlertFiringCacheKeys(models.AlertRuleQuery{
		TenantId:         rule.TenantId,
		RuleId:           rule.RuleId,
		DatasourceIdList: rule.DatasourceIdList,
	})
	if err != nil {
		return fmt.Errorf("failed to get firing cache keys: %v", err)
	}

	recoverKeys := tools.GetSliceDifference(firingKeys, curKeys)
	if len(recoverKeys) == 0 {
		return nil
	}

	curTime := time.Now().Unix()
	for _, key := range recoverKeys {
		if err := t.processRecovery(key, curTime); err != nil {
			logc.Error(t.ctx.Ctx, fmt.Sprintf("Recovery processing failed for key %s: %v", key, err))
		}
	}

	return nil
}

// processRecovery 处理单个告警的恢复
func (t *AlertRule) processRecovery(key string, curTime int64) error {
	event := ctx.Redis.Event().GetCache(key)
	if event.IsRecovered {
		return nil
	}

	if wTime, exists := t.alarmRecoverWaitStore.Get(key); !exists {
		t.alarmRecoverWaitStore.Set(key, curTime)
		return nil
	} else {
		recoveryTime := time.Unix(wTime, 0).Add(time.Minute * time.Duration(global.Config.Server.AlarmConfig.RecoverWait))
		if time.Now().Before(recoveryTime) {
			return nil
		}
	}

	// 更新恢复状态
	event.IsRecovered = true
	event.RecoverTime = curTime
	event.LastSendTime = 0

	ctx.Redis.Event().SetCache("Firing", event, 0)

	t.alarmRecoverWaitStore.Remove(key)
	return nil
}

// GC 执行垃圾回收
//
//	func (t *AlertRule) GC(rule models.AlertRule, curFiringKeys, curPendingKeys []string) {
//		go process.GcPendingCache(t.ctx, rule, curPendingKeys)
//		go process.GcRecoverWaitCache(t.ctx, t.alarmRecoverWaitStore, rule, curFiringKeys)
//	}
func (t *AlertRule) GC(rule models.AlertRule, curFiringKeys, curPendingKeys []string) {
	g := new(errgroup.Group)

	g.Go(func() error {
		process.GcPendingCache(t.ctx, rule, curPendingKeys)
		return nil
	})

	g.Go(func() error {
		process.GcRecoverWaitCache(t.ctx, t.alarmRecoverWaitStore, rule, curFiringKeys)
		return nil
	})

	if err := g.Wait(); err != nil {
		logc.Error(t.ctx.Ctx, fmt.Sprintf("GC process failed: %v", err))
	}
}

// RePushTask 重新推送所有启用的规则
func (t *AlertRule) RePushTask() error {
	ruleList, err := t.getRuleList()
	if err != nil {
		logc.Error(t.ctx.Ctx, err.Error())
		return fmt.Errorf("failed to get rule list: %v", err)
	}

	g := new(errgroup.Group)
	for _, rule := range ruleList {
		rule := rule
		g.Go(func() error {
			return t.Submit(rule)
		})
	}

	if err := g.Wait(); err != nil {
		logc.Error(t.ctx.Ctx, err.Error())
		return fmt.Errorf("failed to re-push task: %v", err)
	}

	return nil
}

// isRuleEnabled 检查规则是否启用
func (t *AlertRule) isRuleEnabled(ruleId string) bool {
	// 直接检查数据库或缓存中的当前启用状态
	return *t.ctx.DB.Rule().GetRuleObject(ruleId).Enabled
}

func (t *AlertRule) getRuleList() ([]models.AlertRule, error) {
	var ruleList []models.AlertRule
	if err := t.ctx.DB.DB().Where("enabled = ?", "1").Find(&ruleList).Error; err != nil {
		return ruleList, fmt.Errorf("获取 Rule List 失败, err: %s", err.Error())
	}
	return ruleList, nil
}
