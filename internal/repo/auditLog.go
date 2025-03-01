package repo

import (
	"ops-monitor/internal/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type (
	AuditLogRepo struct {
		entryRepo
	}

	InterAuditLogRepo interface {
		List(r models.AuditLogQuery) (models.AuditLogResponse, error)
		Search(r models.AuditLogQuery) (models.AuditLogResponse, error)
		Create(r models.AuditLog) error
	}
)

func newAuditLogInterface(db *gorm.DB, g InterGormDBCli) InterAuditLogRepo {
	return &AuditLogRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (a AuditLogRepo) Create(r models.AuditLog) error {
	err := a.db.Model(&models.AuditLog{}).Create(r).Error
	if err != nil {
		return err
	}

	return nil
}

func (a AuditLogRepo) List(r models.AuditLogQuery) (models.AuditLogResponse, error) {
	var db = a.db.Model(&models.AuditLog{})
	var data []models.AuditLog
	var count int64

	pageIndexInt := r.Page.Index
	pageSizeInt := r.Page.Size

	db.Where("tenant_id = ?", r.TenantId)

	db.Count(&count)

	db.Limit(int(pageSizeInt)).Offset(int((pageIndexInt - 1) * pageSizeInt)).Order("created_at desc")
	err := db.Find(&data).Error
	if err != nil {
		return models.AuditLogResponse{}, err
	}

	d := models.AuditLogResponse{
		List: data,
		Page: models.Page{
			Index: pageIndexInt,
			Size:  pageSizeInt,
			Total: count,
		},
	}
	return d, nil
}

func (a AuditLogRepo) Search(r models.AuditLogQuery) (models.AuditLogResponse, error) {
	var db = a.db.Model(&models.AuditLog{})
	var data []models.AuditLog
	var count int64

	pageIndexInt := r.Page.Index
	pageSizeInt := r.Page.Size

	db.Where("tenant_id = ?", r.TenantId)

	if r.Scope != "" {
		curTime := time.Now()
		i, _ := strconv.Atoi(r.Scope)
		eTime := curTime.Add(-time.Duration(i) * (time.Hour * 24))
		db.Where("created_at >= ?", eTime.Unix())
	}

	if r.Query != "" {
		db.Where("username LIKE ? OR ip_address LIKE ? OR audit_type LIKE ?",
			"%"+r.Query+"%", "%"+r.Query+"%", "%"+r.Query+"%")
	}

	db.Count(&count)

	db.Limit(int(pageSizeInt)).Offset(int((pageIndexInt - 1) * pageSizeInt)).Order("created_at desc")

	err := db.Find(&data).Error
	if err != nil {
		return models.AuditLogResponse{}, err
	}
	d := models.AuditLogResponse{
		List: data,
		Page: models.Page{
			Index: pageIndexInt,
			Size:  pageSizeInt,
			Total: count,
		},
	}

	return d, nil
}
