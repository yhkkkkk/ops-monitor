package main

import (
	"ops-monitor/initialization"
	"ops-monitor/internal/global"
)

const Version = "v0.0.1"

func main() {
	global.Version = Version
	initialization.InitBasic()
	initialization.InitRoute()
}
