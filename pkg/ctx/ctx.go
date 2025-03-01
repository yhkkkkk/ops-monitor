package ctx

import (
	"context"
	"sync"

	"ops-monitor/internal/cache"
	"ops-monitor/internal/repo"
)

type Context struct {
	DB    repo.InterEntryRepo
	Redis cache.InterEntryCache
	Ctx   context.Context
	Mux   sync.RWMutex
}

var (
	instance *Context
	DB       repo.InterEntryRepo
	Redis    cache.InterEntryCache
	Ctx      context.Context
	mu       sync.RWMutex
)

func NewContext(ctx context.Context, db repo.InterEntryRepo, redis cache.InterEntryCache) *Context {
	mu.Lock()
	defer mu.Unlock()
	instance = &Context{
		DB:    db,
		Redis: redis,
		Ctx:   ctx,
	}
	return instance
}

// SetRequestContext 安全地设置请求上下文
func SetRequestContext(ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()
	if instance != nil {
		instance.Ctx = ctx
	}
}

func DO() *Context {
	return instance
}
