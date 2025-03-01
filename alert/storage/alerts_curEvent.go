package storage

import (
	"context"
	"ops-monitor/internal/models"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

// AlertsMQErrNotFound is returned if a Store cannot find the Alert.
var (
	AlertsMQErrNotFound = errors.New("alert not found")
)

// AlertsCurEventCache provides cache-coordinated access to alerts, keyed by
// their fingerprint.
type AlertsCurEventCache struct {
	cache *cache.Cache
}

// NewCurAlertsEventMap returns a new Alerts struct.
func NewCurAlertsEventMap() *AlertsCurEventCache {
	// 创建一个默认过期时间为24小时、清理间隔为1小时的缓存
	c := cache.New(24*time.Hour, 1*time.Hour)
	return &AlertsCurEventCache{
		cache: c,
	}
}

// Run starts the GC loop. The interval must be greater than zero; if not, the function will panic.
func (a *AlertsCurEventCache) Run(ctx context.Context, interval time.Duration) {
	// go-cache 已经内置了 GC 机制，这个方法可以保留但不需要实现具体逻辑
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// go-cache 会自动进行垃圾回收
		}
	}
}

// Get retrieves an alert by its fingerprint
func (a *AlertsCurEventCache) Get(fingerprint string) (models.AlertCurEvent, error) {
	if value, found := a.cache.Get(fingerprint); found {
		return value.(models.AlertCurEvent), nil
	}
	return models.AlertCurEvent{}, AlertsMQErrNotFound
}

// Set unconditionally sets the alert in cache
func (a *AlertsCurEventCache) Set(fingerprint string, alert models.AlertCurEvent) error {
	a.cache.Set(fingerprint, alert, cache.DefaultExpiration)
	return nil
}

// Delete removes the Alert with the matching fingerprint from the store
func (a *AlertsCurEventCache) Delete(fingerprint string) error {
	a.cache.Delete(fingerprint)
	return nil
}

// List returns a map of all Alerts currently held in cache
func (a *AlertsCurEventCache) List() map[string]models.AlertCurEvent {
	result := make(map[string]models.AlertCurEvent)
	items := a.cache.Items()

	for k, v := range items {
		result[k] = v.Object.(models.AlertCurEvent)
	}
	return result
}
