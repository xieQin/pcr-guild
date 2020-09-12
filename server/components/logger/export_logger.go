package logger

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"net/http"
	"pcr-guild/enums"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/qbox/qvm-base/components/logger"
)

var (
	// Logger type
	Logger log.Logger
	pid    = uint32(time.Now().UnixNano() % 4294967291)

	// LoggerCtxKey logger context key
	LoggerCtxKey = "logger"
)

// InitLogger 初始化logger
func InitLogger() {
	Logger = log.NewLogfmtLogger(logger.StdLog.LogEntry.Logger.Out)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
}

// ReqLoggerMiddleware 生成每个请求的logger
func ReqLoggerMiddleware(ctx context.Context, request *http.Request) context.Context {
	log := GetLoggerFromReq(request)
	return context.WithValue(ctx, LoggerCtxKey, log)
}

// GetLoggerFromReq get logger from request
func GetLoggerFromReq(req *http.Request) *logger.Logger {
	reqID := req.Header.Get(enums.RequestIDHeaderKey)
	if reqID == "" {
		reqID = GenReqID()
	}
	log := logger.New(reqID)
	return log
}

// GenReqID 生成一个简单的 ReqID
func GenReqID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

// ReqLogger Get Log from ctx
func ReqLogger(ctx context.Context) *logger.Logger {
	l := ctx.Value(LoggerCtxKey)
	if log, ok := l.(*logger.Logger); ok {
		return log
	}

	logger.StdLog.Errorf("logger is missing in ctx, using std logger!")

	// default use std logger
	return logger.StdLog
}

// New return logger
func New(o ...interface{}) *logger.Logger {
	var reqID string

	if len(o) > 0 && o[0] != nil {
		a := o[0]

		l, ok := a.(*logger.Logger)
		if ok {
			return l
		}

		if reqID, ok = a.(string); !ok {
			if context, ok := a.(*gin.Context); ok {
				reqID = context.Writer.Header().Get(enums.RequestIDHeaderKey)
			}
		}
	}

	if len(reqID) == 0 {
		reqID = GenReqID()
	}

	log := logger.New(reqID)
	return log
}
