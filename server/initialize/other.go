package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"gin-vue-admin/server/global"
	"gin-vue-admin/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.ECOVACS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.ECOVACS_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
