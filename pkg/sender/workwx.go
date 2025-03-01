package sender

import (
	"bytes"
	"fmt"
	"ops-monitor/pkg/tools"

	"github.com/pkg/errors"
)

type WeChatResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WeChatMessage struct {
	MsgType string            `json:"msgtype"`
	Text    WeChatTextContent `json:"text"`
}

type WeChatTextContent struct {
	Content string `json:"content"`
}

func SendToWeChat(hook, msg string) error {
	textContentByte := bytes.NewReader([]byte(msg))
	res, err := tools.Post(nil, hook, textContentByte, 10)
	if err != nil {
		return err
	}

	var response WeChatResponse
	if err := tools.ParseReaderBody(res.Body, &response); err != nil {
		return errors.New(fmt.Sprintf("Error unmarshalling WorkWx response: %s", err.Error()))
	}

	if response.ErrCode != 0 {
		return errors.New(response.ErrMsg)
	}

	return nil
}
