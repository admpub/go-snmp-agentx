//go:build !windows && !plan9 && syslog
// +build !windows,!plan9,syslog

package logger

import (
	"log"
	"log/syslog"
	//myLog "github.com/admpub/log"
)

// 初始化 syslog
func initSyslog() {
	// tail -f /var/log/snmp_agentx.log
	sysLogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "snmp_agentx")
	if err != nil {
		log.Fatalf("Failed to initialize syslog: %v", err)
	}
	initLogger(sysLogger)
}

func init() {
	initSyslog()

	//initDefaultLog()
}
