package request

import (
	"gin-vue-admin/server/model/common/request"
	"gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
