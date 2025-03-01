package sender

import (
	"bytes"
	"fmt"
	"ops-monitor/pkg/tools"

	"github.com/pkg/errors"
)

type FeiShuResponseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SendToFeiShu(hook, msg string) error {
	cardContentByte := bytes.NewReader([]byte(msg))
	res, err := tools.Post(nil, hook, cardContentByte, 10)
	if err != nil {
		return err
	}

	var response FeiShuResponseMsg
	if err := tools.ParseReaderBody(res.Body, &response); err != nil {
		return errors.New(fmt.Sprintf("Error unmarshalling Feishu response: %s", err.Error()))
	}
	if response.Code != 0 {
		return errors.New(response.Msg)
	}

	return nil
}
