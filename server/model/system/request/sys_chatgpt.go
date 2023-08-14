package request

import (
	"gin-vue-admin/server/model/common/request"
	"gin-vue-admin/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
