package cron

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/gcron"
	"github.com/hhcool/gtls/log"
	"strconv"
	"time"
)

const fm = "2006-01-02 15:04:05"

// New
// @Description: 定时任务检查周期
// @param ct
func New(ct time.Duration) {
	go func() {
		time.Sleep(time.Second * 10)
		checkTask()
		t := time.NewTicker(ct)
		for {
			select {
			case <-t.C:
				checkTask()
			}
		}
	}()
}

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

type check struct {
	Name string
	Time string
}

// checkTask
// @Auth: oak  2021-08-11 13:33:55
// @Description:  检查正在执行的定时任务，输出到日志
func checkTask() {
	l := "检查定时任务：-----------"
	for _, entry := range gcron.Entries() {
		entry := entry
		var t check
		t.Time = entry.Time.Format(fm)
		t.Name = entry.Name
		b, _ := json.Marshal(t)
		l = fmt.Sprintf("%s\n%s", l, string(b))
	}
	log.Info(fmt.Sprintf("%s\n%s", l, "------------------------------------------------------"))
}
