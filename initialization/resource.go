package initialization

import (
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
	"runtime"
	"time"
)

func InitResource(ctx *ctx.Context) {
	ticker := time.Tick(time.Second * 30)
	layout := "2006-01-02 15:04:05"
	go func() {
		for range ticker {
			curAt := time.Now()
			goNum := runtime.NumGoroutine()
			cleanupOldData(curAt)
			ctx.DB.DB().Model(&models.ServiceResource{}).Create(&models.ServiceResource{
				ID:    uint(curAt.Unix()),
				Time:  curAt.Format(layout),
				Value: goNum,
				Label: "goroutine",
			})
		}
	}()
}

func cleanupOldData(curAt time.Time) {
	c := ctx.DO()
	cutoffTime := curAt.Add(-6 * time.Hour)
	c.DB.DB().Where("time < ?", cutoffTime).Delete(&models.ServiceResource{})
}
