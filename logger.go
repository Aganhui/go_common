package main

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const (
	//jarvandeng: 统一日志上报字段名
	ContextTraceID     = "xverseTraceId"
	ContextRoomID      = "xverseRoomId"
	ContextUserID      = "xverseUserId"
	ContextSessionID   = "xverseSessionId"
	ContextServiceName = "xverseServiceName"
)

func Debug(args ...interface{}) {
	sugar.Debug(args...)
	return
}

// Info info
func Info(args ...interface{}) {
	sugar.Info(args...)
	return
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
	return
}

func Error(args ...interface{}) {
	sugar.Error(args...)
	return
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
	return
}

// Info info
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
	return
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
	return
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
	return
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	sugaredLogger := sugar0
	traceID, ok := ctx.Value(ContextTraceID).(string)
	if ok && traceID != "" {
		sugaredLogger = sugaredLogger.With(zap.String(ContextTraceID, traceID))
	}
	roomID, ok := ctx.Value(ContextRoomID).(string)
	if ok && roomID != "" {
		sugaredLogger = sugaredLogger.With(zap.String(ContextRoomID, roomID))
	}
	userID, ok := ctx.Value(ContextUserID).(string)
	if ok && userID != "" {
		sugaredLogger = sugaredLogger.With(zap.String(ContextUserID, userID))
	}
	sessionID, ok := ctx.Value(ContextSessionID).(string)
	if ok && sessionID != "" {
		sugaredLogger = sugaredLogger.With(zap.String(ContextSessionID, sessionID))
	}
	serviceName, ok := ctx.Value(ContextServiceName).(string)
	if ok && serviceName != "" {
		sugaredLogger = sugaredLogger.With(zap.String(ContextServiceName, serviceName))
	}
	return sugaredLogger
}

func WithIncomingContext(ctx context.Context) *zap.SugaredLogger {
	sugaredLogger := sugar
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return sugaredLogger
	}
	tidkey := strings.ToLower(ContextTraceID)
	if traceID, ok := md[tidkey]; ok && len(traceID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextTraceID, traceID[0]))
	}
	ridkey := strings.ToLower(ContextRoomID)
	if roomID, ok := md[ridkey]; ok && len(roomID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextRoomID, roomID[0]))
	}
	uidkey := strings.ToLower(ContextUserID)
	if userID, ok := md[uidkey]; ok && len(userID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextUserID, userID[0]))
	}
	sidkey := strings.ToLower(ContextSessionID)
	if sessionID, ok := md[sidkey]; ok && len(sessionID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextSessionID, sessionID[0]))
	}
	svckey := strings.ToLower(ContextServiceName)
	if svcName, ok := md[svckey]; ok && len(svcName) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextServiceName, svcName[0]))
	}
	return sugaredLogger
}

func WithOutgoingContext(ctx context.Context) *zap.SugaredLogger {
	sugaredLogger := sugar
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return sugaredLogger
	}
	tidkey := strings.ToLower(ContextTraceID)
	if traceID, ok := md[tidkey]; ok && len(traceID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextTraceID, traceID[0]))
	}
	ridkey := strings.ToLower(ContextRoomID)
	if roomID, ok := md[ridkey]; ok && len(roomID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextRoomID, roomID[0]))
	}
	uidkey := strings.ToLower(ContextUserID)
	if userID, ok := md[uidkey]; ok && len(userID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextUserID, userID[0]))
	}
	sidkey := strings.ToLower(ContextSessionID)
	if sessionID, ok := md[sidkey]; ok && len(sessionID) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextSessionID, sessionID[0]))
	}
	svckey := strings.ToLower(ContextServiceName)
	if svcName, ok := md[svckey]; ok && len(svcName) != 0 {
		sugaredLogger = sugaredLogger.With(zap.String(ContextServiceName, svcName[0]))
	}
	return sugaredLogger
}

func NewGrpcCtx(md metadata.MD) context.Context {

	return nil
}

func GetKeyFromIncomingCtx(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	key = strings.ToLower(key)
	if v, ok := md[key]; ok && len(v) != 0 {
		return v[0]
	}
	return ""

}

func GetKeyFromOutgoingCtx(ctx context.Context, key string) string {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return ""
	}
	key = strings.ToLower(key)
	if v, ok := md[key]; ok && len(v) != 0 {
		return v[0]
	}
	return ""

}
