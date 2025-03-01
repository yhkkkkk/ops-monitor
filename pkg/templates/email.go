package templates

import "ops-monitor/internal/models"

func emailTemplate(alert models.AlertCurEvent, noticeTmpl models.NoticeTemplateExample) string {
	return ParserTemplate("Event", alert, noticeTmpl.Template)
}
