package tools

import (
	"context"
	"fmt"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logc"
)

func ConvertStringToInt(str string) int {
	num, err := cast.ToIntE(str)
	if err != nil {
		logc.Error(context.Background(), fmt.Sprintf("Convert String to int failed, err: %s", err.Error()))
		return 0
	}
	return num
}

func ConvertStringToInt64(str string) int64 {
	num64, err := cast.ToInt64E(str)
	if err != nil {
		logc.Error(context.Background(), fmt.Sprintf("Convert String to int64 failed, err: %s", err.Error()))
		return 0
	}
	return num64
}
