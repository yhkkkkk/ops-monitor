package client

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEmailClient_Send(t *testing.T) {
	eCli := NewEmailClient("smtp.qq.com", "811702605@qq.com", "xxx", 25)
	err := eCli.Send([]string{"811702605@qq.com"}, nil, "ops-monitor监控报警平台", []byte(`
{{ define "Event" -}}
{{- if not .IsRecovered -}}
<p>==========<strong>告警通知</strong>==========</p>
<strong>🤖 报警类型:</strong> ${rule_name}<br>
<strong>🫧 报警指纹:</strong> ${fingerprint}<br>
<strong>📌 报警等级:</strong> ${severity}<br>
<strong>🖥 报警主机:</strong> ${metric.node_name}<br>
<strong>🧚 容器名称:</strong> ${metric.pod}<br>
<strong>☘️ 业务环境:</strong> ${metric.namespace}<br>
<strong>🕘 开始时间:</strong> ${first_trigger_time_format}<br>
<strong>👤 值班人员:</strong> ${duty_user}<br>
<strong>📝 报警事件:</strong> ${annotations}<br>
{{- else -}}
<p>==========<strong>恢复通知</strong>==========</p>
<strong>🤖 报警类型:</strong> ${rule_name}<br>
<strong>🫧 报警指纹:</strong> ${fingerprint}<br>
<strong>📌 报警等级:</strong> ${severity}<br>
<strong>🖥 报警主机:</strong> ${metric.node_name}<br>
<strong>🧚 容器名称:</strong> ${metric.pod}<br>
<strong>☘️ 业务环境:</strong> ${metric.namespace}<br>
<strong>🕘 开始时间:</strong> ${first_trigger_time_format}<br>
<strong>🕘 恢复时间:</strong> ${recover_time_format}<br>
<strong>👤 值班人员:</strong> ${duty_user}<br>
<strong>📝 报警事件:</strong> ${annotations}<br>
{{- end -}}
{{ end }}
`))
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}
