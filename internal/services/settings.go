package services

import (
	"ops-monitor/internal/global"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
)

type (
	settingService struct {
		ctx *ctx.Context
	}

	InterSettingService interface {
		Save(req interface{}) (interface{}, interface{})
		Get() (interface{}, interface{})
	}
)

func newInterSettingService(ctx *ctx.Context) InterSettingService {
	return settingService{
		ctx: ctx,
	}
}

func (a settingService) Save(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Settings)
	if a.ctx.DB.Setting().Check() {
		err := a.ctx.DB.Setting().Update(*r)
		if err != nil {
			return nil, err
		}
	} else {
		err := a.ctx.DB.Setting().Create(*r)
		if err != nil {
			return nil, err
		}
	}

	global.Config.Server.AlarmConfig = r.AlarmConfig
	return nil, nil
}

func (a settingService) Get() (interface{}, interface{}) {
	get, err := a.ctx.DB.Setting().Get()
	if err != nil {
		return nil, err
	}
	get.AppVersion = global.Version

	return get, nil
}
