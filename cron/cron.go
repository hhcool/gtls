package cron

import (
	"fmt"
	"github.com/gogf/gf/os/gcron"
	"strconv"
	"time"
)

// RunCron
// @Auth: oak  2021-08-11 13:35:01
// @Description:  启动一个cron任务
// @param taskType
// @param code
// @param cron
// @param task
func RunCron(taskType string, tag string, cron string, task func()) {
	Remove(taskType, tag)
	_, _ = gcron.Add(cron, task, formatName(taskType, tag))
}

// RunInterval
// @Auth: oak  2021-08-11 13:34:23
// @Description:  启动一个普通定时任务，支持单位s/m/h/d等，默认分钟
// @param taskType
// @param code
// @param t
// @param task
func RunInterval(taskType string, tag string, t interface{}, task func()) {
	Remove(taskType, tag)
	_, _ = gcron.Add(formatEvery(t), task, formatName(taskType, tag))
}

func RunLadderInterval(taskType string, tag string, t time.Duration, task func()) {

}

// Remove
// @Auth: oak  2021-08-11 13:34:13
// @Description:  移除任务
// @param taskType
// @param code
func Remove(taskType string, tag string) {
	gcron.Remove(formatName(taskType, tag))
}

func formatName(taskType string, tag string) string {
	if tag == "" {
		tag = "DEFAULT"
	}
	return fmt.Sprintf("%s_%s", taskType, tag)
}
func formatEvery(time interface{}) string {
	switch time.(type) {
	case string:
		t, err := strconv.Atoi(time.(string))
		if err == nil {
			return fmt.Sprintf("@every %dm", t)
		}
	case int:
		return fmt.Sprintf("@every %dm", time)
	}
	return fmt.Sprintf("@every %s", time)
}
