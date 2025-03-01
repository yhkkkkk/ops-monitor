package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type NoticeController struct{}

/*
通知对象 API
/api/ops/sender
*/
func (nc NoticeController) API(gin *gin.RouterGroup) {
	noticeA := gin.Group("notice")
	noticeA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		noticeA.POST("noticeCreate", nc.Create)
		noticeA.POST("noticeUpdate", nc.Update)
		noticeA.POST("noticeDelete", nc.Delete)
	}

	noticeB := gin.Group("notice")
	noticeB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		noticeB.GET("noticeList", nc.List)
		noticeB.GET("noticeSearch", nc.Search)
		noticeB.GET("noticeRecordList", nc.ListRecord)
		noticeB.GET("noticeRecordMetric", nc.GetRecordMetric)
	}
}

func (nc NoticeController) List(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.List(r)
	})
}

func (nc NoticeController) Create(ctx *gin.Context) {
	r := new(models.AlertNotice)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Create(r)
	})
}

func (nc NoticeController) Update(ctx *gin.Context) {
	r := new(models.AlertNotice)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Update(r)
	})
}

func (nc NoticeController) Delete(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Delete(r)
	})
}

func (nc NoticeController) Get(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Get(r)
	})

}

func (nc NoticeController) Check(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Check(r)
	})
}

func (nc NoticeController) Search(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Search(r)
	})
}

func (nc NoticeController) ListRecord(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.ListRecord(r)
	})
}

func (nc NoticeController) GetRecordMetric(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.GetRecordMetric(r)
	})
}
