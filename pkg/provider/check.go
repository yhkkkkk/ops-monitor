package provider

import (
	"context"
	"fmt"
	"ops-monitor/internal/models"

	"github.com/zeromicro/go-zero/core/logc"
)

func CheckDatasourceHealth(datasource models.AlertDataSource) bool {
	var (
		err   error
		check bool
	)

	switch datasource.Type {
	case "Prometheus":
		prometheusClient, err := NewPrometheusClient(datasource)
		if err == nil {
			check, err = prometheusClient.Check()
		}
	case "VictoriaMetrics":
		vmClient, err := NewVictoriaMetricsClient(datasource)
		if err == nil {
			check, err = vmClient.Check()
		}
	case "ElasticSearch":
		searchClient, err := NewElasticSearchClient(context.Background(), datasource)
		if err == nil {
			check, err = searchClient.Check()
		}
	case "AliCloudSLS":
		slsClient, err := NewAliCloudSlsClient(datasource)
		if err == nil {
			check, err = slsClient.Check()
		}
	case "Jaeger":
		jaegerClient, err := NewJaegerClient(datasource)
		if err == nil {
			check, err = jaegerClient.Check()
		}
	case "CloudWatch":
		return true
	}

	// 检查数据源健康状况并返回结果
	if err != nil || !check {
		logc.Errorf(context.Background(), fmt.Sprintf("数据源不健康, Id: %s, Name: %s, Type: %s", datasource.Id, datasource.Name, datasource.Type))
		return false
	}

	return true
}
