package templates

import (
	"ops-monitor/internal/models"
	"ops-monitor/pkg/tools"
)

type WeChatMsg struct {
	MsgType  string      `json:"msgtype"`
	Markdown MarkdownMsg `json:"markdown"`
}

type MarkdownMsg struct {
	Content string `json:"content"`
}

// 企业微信机器人消息卡片模版
func workwxTemplate(alert models.AlertCurEvent, noticeTmpl models.NoticeTemplateExample) string {
	defaultTemplate := WeChatMsg{
		MsgType: "markdown",
		Markdown: MarkdownMsg{
			Content: "",
		},
	}

	if *noticeTmpl.EnableWeChatMarkdown {
		switch alert.IsRecovered {
		case false:
			defaultTemplate.Markdown.Content = ParserTemplate("Markdown", alert, noticeTmpl.TemplateFiring)
		case true:
			defaultTemplate.Markdown.Content = ParserTemplate("Markdown", alert, noticeTmpl.TemplateRecover)
		}
	} else {
		var content string

		titleColor := ParserTemplate("TitleColor", alert, noticeTmpl.Template)
		if titleColor == "red" {
			content += "<font color=\"warning\">"
		} else {
			content += "<font color=\"info\">"
		}
		content += "## " + ParserTemplate("Title", alert, noticeTmpl.Template)
		content += "</font>\n\n"
		content += ParserTemplate("Event", alert, noticeTmpl.Template)
		content += "\n\n"
		content += "---\n"
		content += ParserTemplate("Footer", alert, noticeTmpl.Template)
		defaultTemplate.Markdown.Content = content
	}

	return tools.JsonMarshal(defaultTemplate)
}
