package models

import "ops-monitor/config"

type Settings struct {
	IsInit      int                `json:"isInit"`
	AlarmConfig config.AlarmConfig `json:"alarmConfig" gorm:"alarmConfig;serializer:json"`
	EmailConfig emailConfig        `json:"emailConfig" gorm:"emailConfig;serializer:json"`
	AppVersion  string             `json:"appVersion" gorm:"-"`
}

type emailConfig struct {
	ServerAddress string `json:"serverAddress"`
	Port          int    `json:"port"`
	Email         string `json:"email"`
	Token         string `json:"token"`
}
