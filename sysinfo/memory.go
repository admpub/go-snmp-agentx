package sysinfo

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/mem"

	"go-snmp-agentx/util"
)

// MemoryStat 获取内存使用率
func MemoryStat() interface{} {
	var result = make(map[string]interface{})

	// 获取内存信息
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	result["total"] = formatSize(vmStat.Total)
	result["used"] = formatSize(vmStat.Used)
	result["available"] = formatSize(vmStat.Available)
	result["usedPercent"] = fmt.Sprintf("%.2f%%", vmStat.UsedPercent)

	return util.Map2JSON(result)
}

func formatSize(size uint64) string {
	if size > 1024*1024*1024 {
		return fmt.Sprintf("%.2fG", float64(size/1024/1024/1024))
	}

	return fmt.Sprintf("%.2fM", float64(size/1024/1024))
}
