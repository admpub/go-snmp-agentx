package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	myLog "github.com/admpub/log"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

// 初始化 logger
func initLogger(w io.Writer) {
	debugLogger = log.New(w, "", log.Lmsgprefix)
	infoLogger = log.New(w, "", log.Lmsgprefix)
	warnLogger = log.New(w, "", log.Lmsgprefix)
	errorLogger = log.New(w, "", log.Lmsgprefix)
}

// Debug 级别日志
func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

// Info 级别日志
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// Warn 级别日志
func Warn(v ...interface{}) {
	warnLogger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

// Error 级别日志
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

func initDefaultLog() {
	rootDir := filepath.Dir(os.Args[0])
	if strings.HasPrefix(rootDir, os.TempDir()) {
		rootDir = "."
	}
	t := myLog.NewFileTarget()
	var err error
	t.FileName, err = filepath.Abs(filepath.Join(rootDir, "logs"))
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
		return
	}
	os.MkdirAll(t.FileName, os.ModePerm)
	t.FileName += string(filepath.Separator) + `agentx_{date:2006_01_02}.log`
	l := myLog.SetTarget(t).GetLogger(`AgentX`)
	debugLogger = log.New(l.Writer(myLog.LevelDebug), "", log.Lmsgprefix)
	infoLogger = log.New(l.Writer(myLog.LevelInfo), "", log.Lmsgprefix)
	warnLogger = log.New(l.Writer(myLog.LevelWarn), "", log.Lmsgprefix)
	errorLogger = log.New(l.Writer(myLog.LevelError), "", log.Lmsgprefix)
}
