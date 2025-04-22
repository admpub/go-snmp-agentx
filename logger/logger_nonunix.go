//go:build windows || plan9 || !syslog
// +build windows plan9 !syslog

package logger

// 初始化 log
func init() {
	initDefaultLog()
}
