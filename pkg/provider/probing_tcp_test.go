package provider

import (
	"fmt"
	"ops-monitor/pkg/tools"
	"testing"
)

func TestNewEndpointTcper(t *testing.T) {
	buildOption := EndpointOption{
		Endpoint: "8.147.234.89:80",
		Timeout:  10,
	}
	pilot, err := NewEndpointTcper().Pilot(buildOption)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tools.JsonMarshal(pilot))
}
