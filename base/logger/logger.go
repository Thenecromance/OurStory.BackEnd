package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const (
	LogPath = "setting/stories.log"
)

var (
	logger *zap.SugaredLogger

	sqlWriter *loggerWriter
)

type loggerWriter struct {
	callback func(p []byte) (n int, err error)
}

func (l *loggerWriter) Write(p []byte) (n int, err error) {
	if l.callback == nil {
		return len(p), nil
	}
	return l.callback(p)
}

func GetWriter() io.Writer {
	return sqlWriter
}

func init() {
	Init()
}
func Init() {
	sqlWriter = &loggerWriter{}

	cfg := zap.NewProductionEncoderConfig()
	//zap.NewDevelopmentConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	file, _ := os.Create(LogPath)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(sqlWriter), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(file), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)
	logger = zap.New(core).Sugar()

}

func AddCallback(callback func(p []byte) (n int, err error)) {
	sqlWriter.callback = callback
}
func Get() *zap.SugaredLogger {
	return logger
}
