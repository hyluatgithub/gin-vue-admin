package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"gin-vue-admin/server/config"
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/utils"
)

func Timer() {
	if global.ECOVACS_CONFIG.Timer.Start {
		for i := range global.ECOVACS_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.ECOVACS_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.ECOVACS_Timer.AddTaskByFunc("ClearDB", global.ECOVACS_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.ECOVACS_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.ECOVACS_CONFIG.Timer.Detail[i])
		}
	}
}
