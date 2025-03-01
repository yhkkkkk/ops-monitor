package provider

import (
	"bytes"
	"fmt"
	"net/http"
	"ops-monitor/pkg/tools"
	"time"
)

const (
	GetHTTPMethod  = "GET"
	PostHTTPMethod = "POST"
)

type HTTPer struct{}

func NewEndpointHTTPer() EndpointFactoryProvider {
	return HTTPer{}
}

func (h HTTPer) Pilot(option EndpointOption) (EndpointValue, error) {
	var (
		ev  EndpointValue
		res *http.Response
		err error
	)

	// 开始时间
	start := time.Now()
	switch option.HTTP.Method {
	case GetHTTPMethod:
		res, err = tools.Get(option.HTTP.Header, option.Endpoint, option.Timeout)
		if err != nil {
			return ev, err
		}
		defer res.Body.Close()
	case PostHTTPMethod:
		res, err = tools.Post(option.HTTP.Header, option.Endpoint, bytes.NewReader([]byte(option.HTTP.Body)), option.Timeout)
		if err != nil {
			return ev, err
		}
		defer res.Body.Close()
	}
	end := time.Now()
	// 计算请求耗时
	latency := end.Sub(start).Milliseconds()
	fmt.Println(res.Request.RemoteAddr)

	return convertHTTPerToEndpointValue(HttperInformation{
		Address:    res.Request.URL.String(),
		StatusCode: float64(res.StatusCode),
		Latency:    float64(latency),
	}), nil
}

func convertHTTPerToEndpointValue(detail HttperInformation) EndpointValue {
	return EndpointValue{
		"address":    detail.Address,
		"StatusCode": detail.StatusCode,
		"Latency":    detail.Latency,
	}
}
