package sysinfo

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"

	"go-snmp-agentx/util"
)

// CPUStat 获取 CPU 使用率
func CPUStat() interface{} {
	var stat = make(map[string]interface{})

	// 获取 CPU 使用率
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return ""
	}

	var sum float64
	for i, p := range percent {
		sum += p
		key := fmt.Sprintf("cpu%d", i)
		stat[key] = util.RoundFloat(p, 2)
	}

	// 计算平均值
	stat["cpu"] = util.RoundFloat(sum/float64(len(percent)), 2)

	// 将 map 转换为 json 字符串
	return util.Map2JSON(stat)
}
