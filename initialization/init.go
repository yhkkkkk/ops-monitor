package initialization

import (
	"context"
	"fmt"
	"ops-monitor/alert"
	"ops-monitor/config"
	"ops-monitor/internal/cache"
	"ops-monitor/internal/global"
	"ops-monitor/internal/models"
	"ops-monitor/internal/repo"
	"ops-monitor/internal/services"
	"ops-monitor/pkg/ctx"

	"github.com/zeromicro/go-zero/core/logc"
	"golang.org/x/sync/errgroup"
)

func InitBasic() {
	// 初始化配置
	global.Config = config.InitConfig()

	dbRepo := repo.NewRepoEntry()
	rCache := cache.NewEntryCache()
	ctx := ctx.NewContext(context.Background(), dbRepo, rCache)

	services.NewServices(ctx)

	// 启用告警评估携程
	alert.Initialize(ctx)

	// 初始化监控分析数据(主要是当前程序的协程数)
	InitResource(ctx)

	// 初始化权限数据
	InitPermissionsSQL(ctx)

	// 初始化角色数据
	InitUserRolesSQL(ctx)

	// 导入数据源 Client 到存储池
	// importClientPools(ctx)
}

func importClientPools(ctx *ctx.Context) {
	list, err := ctx.DB.Datasource().List(models.DatasourceQuery{})
	if err != nil {
		logc.Error(ctx.Ctx, err.Error())
		return
	}

	g := new(errgroup.Group)
	for _, datasource := range list {
		datasource := datasource
		if !*datasource.Enabled {
			continue
		}
		g.Go(func() error {
			err := services.DatasourceService.WithAddClientToProviderPools(datasource)
			if err != nil {
				logc.Error(ctx.Ctx, fmt.Sprintf("添加到 Client 存储池失败, err: %s", err.Error()))
				return err
			}
			return nil
		})
	}
}
