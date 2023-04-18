package logger

import (
	"time"

	"go.uber.org/zap"
)

func Debuge(template string, args ...interface{}) {
	Sugar().With(zap.String("task_id", logTag.TaskID)).
		With(zap.String("machine_addr", logTag.MachineAddr)).
		With(zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05"))).
		Debugf(template, args...)
}

//Info info
func Infoe(template string, args ...interface{}) {
	Sugar().With(zap.String("task_id", logTag.TaskID)).
		With(zap.String("machine_addr", logTag.MachineAddr)).
		With(zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05"))).
		Infof(template, args...)
}

func Warne(template string, args ...interface{}) {
	Sugar().With(zap.String("task_id", logTag.TaskID)).
		With(zap.String("machine_addr", logTag.MachineAddr)).
		With(zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05"))).
		Warnf(template, args...)
}

func Errore(template string, args ...interface{}) {
	Sugar().With(zap.String("task_id", logTag.TaskID)).
		With(zap.String("machine_addr", logTag.MachineAddr)).
		With(zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05"))).
		Errorf(template, args...)
}
