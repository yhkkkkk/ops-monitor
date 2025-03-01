package storage

import (
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

// AlarmRecoverWaitStore 定义接口
type AlarmRecoverWaitStore interface {
	Set(key string, t int64)
	SetWithExpiration(key string, t int64, expiration time.Duration)
	Get(key string) (int64, bool)
	Remove(key string)
	Search(keyPrefix string) []string
	GetAll() map[string]int64
	Flush()
}

// alarmRecoverWaitStore 存储等待被恢复的告警的 Key
type alarmRecoverWaitStore struct {
	cache *cache.Cache
}

// NewAlarmRecoverStore 创建一个新的告警恢复存储实例
func NewAlarmRecoverStore() AlarmRecoverWaitStore {
	// 创建一个默认过期时间为24小时、清理间隔为1小时的缓存
	c := cache.New(24*time.Hour, 1*time.Hour)
	return &alarmRecoverWaitStore{
		cache: c,
	}
}

// Set 设置告警恢复等待记录
func (a *alarmRecoverWaitStore) Set(key string, t int64) {
	a.cache.Set(key, t, cache.DefaultExpiration)
}

// SetWithExpiration 设置告警恢复等待记录，并指定过期时间
func (a *alarmRecoverWaitStore) SetWithExpiration(key string, t int64, expiration time.Duration) {
	a.cache.Set(key, t, expiration)
}

// Get 获取告警恢复等待记录
func (a *alarmRecoverWaitStore) Get(key string) (int64, bool) {
	if value, found := a.cache.Get(key); found {
		return value.(int64), true
	}
	return 0, false
}

// Remove 删除告警恢复等待记录
func (a *alarmRecoverWaitStore) Remove(key string) {
	a.cache.Delete(key)
}

// Search 搜索指定前缀的告警恢复等待记录
func (a *alarmRecoverWaitStore) Search(keyPrefix string) []string {
	var keys []string
	items := a.cache.Items()

	for k := range items {
		if strings.HasPrefix(k, keyPrefix) {
			keys = append(keys, k)
		}
	}
	return keys
}

// GetAll 获取所有告警恢复等待记录
func (a *alarmRecoverWaitStore) GetAll() map[string]int64 {
	result := make(map[string]int64)
	items := a.cache.Items()

	for k, v := range items {
		result[k] = v.Object.(int64)
	}
	return result
}

// Flush 清空所有告警恢复等待记录
func (a *alarmRecoverWaitStore) Flush() {
	a.cache.Flush()
}
