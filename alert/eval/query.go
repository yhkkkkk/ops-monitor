package eval

import (
	"fmt"
	"ops-monitor/alert/process"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
	"ops-monitor/pkg/provider"
	"ops-monitor/pkg/tools"
	"regexp"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
)

// Metrics 包含 Prometheus、VictoriaMetrics 数据源
func metrics(ctx *ctx.Context, datasourceId, datasourceType string, rule models.AlertRule) (curFiringKeys, curPendingKeys []string) {
	pools := ctx.Redis.ProviderPools()
	var (
		resQuery       []provider.Metrics
		externalLabels map[string]interface{}
	)

	switch datasourceType {
	case provider.PrometheusDsProvider:
		cli, err := pools.GetClient(datasourceId)
		if err != nil {
			logc.Errorf(ctx.Ctx, err.Error())
			return
		}

		resQuery, err = cli.(provider.PrometheusProvider).Query(rule.PrometheusConfig.PromQL)
		if err != nil {
			logc.Error(ctx.Ctx, err.Error())
			return
		}

		externalLabels = cli.(provider.PrometheusProvider).GetExternalLabels()
	case provider.VictoriaMetricsDsProvider:
		cli, err := pools.GetClient(datasourceId)
		if err != nil {
			logc.Errorf(ctx.Ctx, err.Error())
			return
		}

		resQuery, err = cli.(provider.VictoriaMetricsProvider).Query(rule.PrometheusConfig.PromQL)
		if err != nil {
			logc.Error(ctx.Ctx, err.Error())
			return
		}

		externalLabels = cli.(provider.VictoriaMetricsProvider).GetExternalLabels()
	default:
		logc.Errorf(ctx.Ctx, fmt.Sprintf("Unsupported metrics type, type: %s", datasourceType))
		return
	}

	if resQuery == nil {
		return
	}

	for _, v := range resQuery {
		for _, ruleExpr := range rule.PrometheusConfig.Rules {
			re := regexp.MustCompile(`([^\d]+)(\d+)`)
			matches := re.FindStringSubmatch(ruleExpr.Expr)
			t, _ := strconv.ParseFloat(matches[2], 64)

			event := func() models.AlertCurEvent {
				event := process.BuildEvent(rule)
				event.DatasourceId = datasourceId
				event.Fingerprint = v.GetFingerprint()
				event.Metric = v.GetMetric()
				event.Metric["severity"] = ruleExpr.Severity
				for ek, ev := range externalLabels {
					event.Metric[ek] = ev
				}
				event.Severity = ruleExpr.Severity
				event.Annotations = tools.ParserVariables(rule.PrometheusConfig.Annotations, event.Metric)

				firingKey := event.GetFiringAlertCacheKey()
				pendingKey := event.GetPendingAlertCacheKey()

				curFiringKeys = append(curFiringKeys, firingKey)
				curPendingKeys = append(curPendingKeys, pendingKey)

				return event
			}

			option := models.EvalCondition{
				Operator:      matches[1],
				QueryValue:    v.Value,
				ExpectedValue: t,
			}

			if process.EvalCondition(option) {
				process.SaveAlertEvent(ctx, event())
			}
		}
	}

	return
}

// Logs 包含 AliSLS、Loki、ElasticSearch 数据源
func logs(ctx *ctx.Context, datasourceId, datasourceType string, rule models.AlertRule) (curFiringKeys []string) {
	var (
		queryRes       []provider.Logs
		count          int
		evalOptions    models.EvalCondition
		externalLabels map[string]interface{}
	)

	pools := ctx.Redis.ProviderPools()
	switch datasourceType {
	case provider.AliCloudSLSDsProviderName:
		cli, err := pools.GetClient(datasourceId)
		if err != nil {
			logc.Errorf(ctx.Ctx, err.Error())
			return
		}

		curAt := time.Now()
		startsAt := tools.ParserDuration(curAt, rule.AliCloudSLSConfig.LogScope, "m")
		queryOptions := provider.LogQueryOptions{
			AliCloudSLS: provider.AliCloudSLS{
				Query:    rule.AliCloudSLSConfig.LogQL,
				Project:  rule.AliCloudSLSConfig.Project,
				LogStore: rule.AliCloudSLSConfig.Logstore,
			},
			StartAt: int32(startsAt.Unix()),
			EndAt:   int32(curAt.Unix()),
		}
		queryRes, count, err = cli.(provider.AliCloudSlsDsProvider).Query(queryOptions)
		if err != nil {
			logc.Error(ctx.Ctx, err.Error())
			return
		}

		externalLabels = cli.(provider.AliCloudSlsDsProvider).GetExternalLabels()

		evalOptions = models.EvalCondition{
			Operator:      rule.AliCloudSLSConfig.EvalCondition.Operator,
			QueryValue:    float64(count),
			ExpectedValue: rule.AliCloudSLSConfig.EvalCondition.ExpectedValue,
		}
	case provider.ElasticSearchDsProviderName:
		cli, err := pools.GetClient(datasourceId)
		if err != nil {
			logc.Errorf(ctx.Ctx, err.Error())
			return
		}

		curAt := time.Now()
		startsAt := tools.ParserDuration(curAt, int(rule.ElasticSearchConfig.Scope), "m")
		queryOptions := provider.LogQueryOptions{
			ElasticSearch: provider.Elasticsearch{
				Index:       rule.ElasticSearchConfig.Index,
				QueryFilter: rule.ElasticSearchConfig.Filter,
			},
			StartAt: tools.FormatTimeToUTC(startsAt.Unix()),
			EndAt:   tools.FormatTimeToUTC(curAt.Unix()),
		}
		queryRes, count, err = cli.(provider.ElasticSearchDsProvider).Query(queryOptions)
		if err != nil {
			logc.Error(ctx.Ctx, err.Error())
			return
		}

		externalLabels = cli.(provider.ElasticSearchDsProvider).GetExternalLabels()

		evalOptions = models.EvalCondition{
			Operator:      ">",
			QueryValue:    float64(count),
			ExpectedValue: 1,
		}
	}

	if count <= 0 {
		return
	}

	for _, v := range queryRes {
		event := func() models.AlertCurEvent {
			event := process.BuildEvent(rule)
			event.DatasourceId = datasourceId
			event.Fingerprint = v.GetFingerprint()
			event.Metric = v.GetMetric()
			for ek, ev := range externalLabels {
				event.Metric[ek] = ev
			}
			event.Annotations = fmt.Sprintf("统计日志条数: %d 条\n%s", count, tools.FormatJson(v.GetAnnotations()[0].(string)))

			key := event.GetPendingAlertCacheKey()
			curFiringKeys = append(curFiringKeys, key)

			return event
		}

		// 评估告警条件
		if process.EvalCondition(evalOptions) {
			process.SaveAlertEvent(ctx, event())
		}
	}

	return
}

// Traces 包含 Jaeger 数据源
func traces(ctx *ctx.Context, datasourceId, datasourceType string, rule models.AlertRule) (curFiringKeys []string) {
	var (
		queryRes       []provider.Traces
		externalLabels map[string]interface{}
	)

	pools := ctx.Redis.ProviderPools()
	switch datasourceType {
	case provider.JaegerDsProviderName:
		curAt := time.Now().UTC()
		startsAt := tools.ParserDuration(curAt, rule.JaegerConfig.Scope, "m")

		cli, err := pools.GetClient(datasourceId)
		if err != nil {
			logc.Errorf(ctx.Ctx, err.Error())
			return
		}

		queryOptions := provider.TraceQueryOptions{
			Tags:    rule.JaegerConfig.Tags,
			Service: rule.JaegerConfig.Service,
			StartAt: startsAt.UnixMicro(),
			EndAt:   curAt.UnixMicro(),
		}
		queryRes, err = cli.(provider.JaegerDsProvider).Query(queryOptions)
		if err != nil {
			logc.Error(ctx.Ctx, err.Error())
			return
		}

		externalLabels = cli.(provider.JaegerDsProvider).GetExternalLabels()
	}

	for _, v := range queryRes {
		event := process.BuildEvent(rule)
		event.DatasourceId = datasourceId
		event.Fingerprint = v.GetFingerprint()
		event.Metric = v.GetMetric()
		for ek, ev := range externalLabels {
			event.Metric[ek] = ev
		}
		event.Annotations = fmt.Sprintf("服务: %s 链路中存在异常状态码接口, TraceId: %s", rule.JaegerConfig.Service, v.TraceId)

		key := event.GetFiringAlertCacheKey()
		curFiringKeys = append(curFiringKeys, key)

		process.SaveAlertEvent(ctx, event)
	}

	return
}
