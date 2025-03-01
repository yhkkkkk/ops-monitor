package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var CodeInfo = map[int64]string{
	200: "OK",
	400: "请求失败",
	401: "Token鉴权失败",
	403: "权限不足",
}

// ResponseData 响应结构体
type ResponseData struct {
	Code int         `json:"code"` // 响应码
	Data interface{} `json:"data"` // 响应数据
	Msg  string      `json:"msg"`  // 响应消息
}

func Response(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, ResponseData{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

func TokenFail(ctx *gin.Context) {
	code := 401
	Response(ctx, code, code, nil, CodeInfo[int64(code)])
}

func PermissionFail(ctx *gin.Context) {
	code := 403
	Response(ctx, code, code, nil, CodeInfo[int64(code)])
}
