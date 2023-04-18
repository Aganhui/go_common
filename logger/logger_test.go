package logger

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func TestLogger(t *testing.T) {
	Init()
	Infof("a:%d, b:%d", 1, 2)

	Sugar().With(zap.String("haha", "lala")).Infof("a:%d, b:%d", 1, 2)

	Errorf("a:%d, b:%d", 1, 2)

	Errore("eslog test %v, %v", 1, 2)

	ctx := context.WithValue(context.Background(), ContextTraceID, "test-traceid")
	WithContext(ctx).Infof("withContext")

	ctx2 := context.WithValue(context.Background(), ContextTraceID, "test-traceid2")
	ctx2 = context.WithValue(ctx2, ContextRoomID, "test-roomid2")
	ctx2 = context.WithValue(ctx2, ContextUserID, "test-userid2")
	ctx2 = context.WithValue(ctx2, ContextSessionID, "test-sessionid2")
	ctx2 = context.WithValue(ctx2, ContextServiceName, "sn")
	log := WithContext(ctx2)
	log.Infof("two parse withContext")
	Info()
	omd := metadata.Pairs(ContextSessionID, "true_session_id")
	octx := metadata.NewOutgoingContext(context.Background(), omd)
	l := WithOutgoingContext(octx)
	l.Info("omd test")

	imd := metadata.Pairs(ContextSessionID, "incoming_true_session_id")
	ictx := metadata.NewIncomingContext(context.Background(), imd)
	val := GetKeyFromIncomingCtx(ictx, ContextSessionID)
	l = WithIncomingContext(ictx)
	l.Info("session id is ", val)

}
