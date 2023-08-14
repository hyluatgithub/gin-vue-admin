package global

import (
	"sync"

	"gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"gin-vue-admin/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ECOVACS_DB     *gorm.DB
	ECOVACS_DBList map[string]*gorm.DB
	ECOVACS_REDIS  *redis.Client
	ECOVACS_CONFIG config.Server
	ECOVACS_VP     *viper.Viper
	// ECOVACS_LOG    *oplogging.Logger
	ECOVACS_LOG                 *zap.Logger
	ECOVACS_Timer               = timer.NewTimerTask()
	ECOVACS_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return ECOVACS_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := ECOVACS_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
