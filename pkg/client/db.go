package client

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ops-monitor/internal/global"
	"ops-monitor/internal/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	// 初始化本地 test.db 数据库文件
	//db, err := gorm.Open(sqlite.Open("data/sql.db"), &gorm.Config{})

	sql := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local&timeout=%s",
		sql.User,
		sql.Pass,
		sql.Host,
		sql.Port,
		sql.DBName,
		sql.Timeout)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Msgf("failed to connect database: %s", err.Error())
		return nil
	}

	// 检查 Product 结构是否变化，变化则进行迁移
	err = db.AutoMigrate(
		&models.DutySchedule{},
		&models.DutyManagement{},
		&models.AlertNotice{},
		&models.AlertDataSource{},
		&models.AlertRule{},
		&models.AlertCurEvent{},
		&models.AlertHisEvent{},
		&models.AlertSilences{},
		&models.Member{},
		&models.UserRole{},
		&models.UserPermissions{},
		&models.NoticeTemplateExample{},
		&models.RuleGroups{},
		&models.RuleTemplateGroup{},
		&models.RuleTemplate{},
		&models.ServiceResource{},
		&models.Tenant{},
		&models.Dashboard{},
		&models.AuditLog{},
		&models.Settings{},
		&models.TenantLinkedUsers{},
		&models.DashboardFolders{},
		&models.AlertSubscribe{},
		&models.NoticeRecord{},
		&models.ProbingRule{},
	)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	log.Info().Msg("database migration success")

	if global.Config.Server.Mode == "debug" {
		db.Debug()
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	log.Info().Msg("database connection success")

	return db
}

func GetDB() *gorm.DB {
	return db
}
