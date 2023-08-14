package initialize

import (
	"gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.ECOVACS_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}
