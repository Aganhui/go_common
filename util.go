package logger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

func LoggerTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
