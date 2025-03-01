package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/client"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logc"
)

type (
	eventCache struct {
		rc *redis.Client
		sync.RWMutex
	}

	InterEventCache interface {
		SetCache(cacheType string, event models.AlertCurEvent, expiration time.Duration)
		DelCache(key string)
		GetCache(key string) models.AlertCurEvent
		GetFirstTime(key string) int64
		GetLastEvalTime(key string) int64
		GetLastSendTime(key string) int64
		SetPECache(event models.ProbingEvent, expiration time.Duration)
		GetPECache(key string) (models.ProbingEvent, error)
		GetPEFirstTime(key string) int64
		GetPELastEvalTime(key string) int64
		GetPELastSendTime(key string) int64
	}
)

func newEventCacheInterface(r *redis.Client) InterEventCache {
	return &eventCache{
		rc: r,
	}
}

func (ec *eventCache) SetCache(cacheType string, event models.AlertCurEvent, expiration time.Duration) {
	ec.Lock()
	defer ec.Unlock()

	alertJson, _ := json.Marshal(event)
	switch cacheType {
	case "Firing":
		client.Redis.Set(event.GetFiringAlertCacheKey(), string(alertJson), expiration)
	case "Pending":
		client.Redis.Set(event.GetPendingAlertCacheKey(), string(alertJson), expiration)
	}

}

func (ec *eventCache) DelCache(key string) {
	ec.Lock()
	defer ec.Unlock()

	// 使用Scan命令获取所有匹配指定模式的键
	iter := client.Redis.Scan(0, key, 0).Iterator()
	keysToDelete := make([]string, 0)

	// 遍历匹配的键
	for iter.Next() {
		key := iter.Val()
		keysToDelete = append(keysToDelete, key)
	}

	if err := iter.Err(); err != nil {
		log.Fatal(err)
	}

	// 批量删除键
	if len(keysToDelete) > 0 {
		err := client.Redis.Del(keysToDelete...).Err()
		if err != nil {
			log.Fatal(err)
		}
		logc.Infof(context.Background(), fmt.Sprintf("移除告警消息 -> %s\n", keysToDelete))
	}
}

func (ec *eventCache) GetCache(key string) models.AlertCurEvent {

	var alert models.AlertCurEvent

	d, err := ec.rc.Get(key).Result()
	_ = json.Unmarshal([]byte(d), &alert)
	if err != nil {
		return alert
	}
	//logc.Info(alert)
	return alert

}

func (ec *eventCache) GetFirstTime(key string) int64 {

	ft := ec.GetCache(key).FirstTriggerTime
	if ft == 0 {
		return time.Now().Unix()
	}
	return ft

}

func (ec *eventCache) GetLastEvalTime(key string) int64 {

	curTime := time.Now().Unix()
	let := ec.GetCache(key).LastEvalTime
	if let == 0 || let < curTime {
		return curTime
	}

	return let

}

func (ec *eventCache) GetLastSendTime(key string) int64 {

	return ec.GetCache(key).LastSendTime

}

func (ec *eventCache) SetPECache(event models.ProbingEvent, expiration time.Duration) {
	ec.Lock()
	defer ec.Unlock()

	alertJson, _ := json.Marshal(event)
	ec.rc.Set(event.GetFiringAlertCacheKey(), string(alertJson), expiration)
}

func (ec *eventCache) GetPECache(key string) (models.ProbingEvent, error) {
	var alert models.ProbingEvent

	d, err := ec.rc.Get(key).Result()
	_ = json.Unmarshal([]byte(d), &alert)
	if err != nil {
		return alert, err
	}
	//global.Logger.Sugar().Info(alert)
	return alert, nil
}

func (ec *eventCache) GetPEFirstTime(key string) int64 {
	cache, _ := ec.GetPECache(key)
	ft := cache.FirstTriggerTime
	if ft == 0 {
		return time.Now().Unix()
	}
	return ft
}

func (ec *eventCache) GetPELastEvalTime(key string) int64 {
	curTime := time.Now().Unix()
	cache, _ := ec.GetPECache(key)
	let := cache.LastEvalTime
	if let == 0 || let < curTime {
		return curTime
	}

	return let
}

func (ec *eventCache) GetPELastSendTime(key string) int64 {
	cache, _ := ec.GetPECache(key)
	return cache.LastSendTime
}
