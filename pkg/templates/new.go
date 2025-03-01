package templates

import (
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
)

type Template struct {
	CardContentMsg string
}

const (
	NoticeTypeFeiShu = "FeiShu"
	NoticeTypeEmail  = "Email"
	NoticeTypeWorkWx = "WorkWx"
)

func NewTemplate(ctx *ctx.Context, alert models.AlertCurEvent, notice models.AlertNotice) Template {
	noticeTmpl := ctx.DB.NoticeTmpl().Get(models.NoticeTemplateExampleQuery{Id: notice.NoticeTmplId})

	switch notice.NoticeType {
	case NoticeTypeFeiShu:
		return Template{CardContentMsg: feishuTemplate(alert, noticeTmpl)}
	case NoticeTypeEmail:
		return Template{CardContentMsg: emailTemplate(alert, noticeTmpl)}
	case NoticeTypeWorkWx:
		return Template{CardContentMsg: workwxTemplate(alert, noticeTmpl)}
	}

	return Template{}
}
