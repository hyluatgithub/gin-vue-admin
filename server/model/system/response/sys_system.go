package response

import "gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
