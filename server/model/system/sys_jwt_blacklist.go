package system

import (
	"gin-vue-admin/server/global"
)

type JwtBlacklist struct {
	global.ECOVACS_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
