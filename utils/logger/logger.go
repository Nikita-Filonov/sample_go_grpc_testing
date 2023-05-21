package logger

import (
	"fmt"
	"go.uber.org/zap"
	"sample_go_grpc_testing/utils/reports"
)

type CtxLogger struct {
	Zap *zap.Logger
}

func (ctxlog *CtxLogger) NewPrefix(prefix string) *CtxLogger {
	return &CtxLogger{ctxlog.Zap.Named(prefix)}
}

func (ctxlog *CtxLogger) InfofJSON(logText string, a ...interface{}) {
	ctxlog.Zap.Info(fmt.Sprintf(logText+": %+v", a...))
	reports.AddAllureAttachmentJSON(logText, a)
}

func (ctxlog *CtxLogger) Errorf(format string, a ...interface{}) {
	ctxlog.Zap.Error(fmt.Sprintf(format, a...))
}
