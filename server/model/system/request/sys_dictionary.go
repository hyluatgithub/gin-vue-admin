package request

import (
	"gin-vue-admin/server/model/common/request"
	"gin-vue-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
